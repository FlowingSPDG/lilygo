package lilygo

import (
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
	"golang.org/x/xerrors"
)

// UTF-8 から ShiftJIS
func utf8ToShiftJIS(str string) io.Reader {
	iostr := strings.NewReader(str)
	return transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
}

func shiftJISToUTF8(r io.Reader) (string, error) {
	// ShiftJIS to UTF-8
	rio := transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	ret, err := io.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func shiftJISToMacRoman(r io.Reader) (string, error) {
	// ShiftJIS to MacRoman
	rio := transform.NewReader(r, charmap.Macintosh.NewDecoder())
	ret, err := io.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func macRomanToShiftJIS(s string) io.Reader {
	// MacRoman to ShiftJIS
	return transform.NewReader(strings.NewReader(s), charmap.Macintosh.NewEncoder())
}

// ConvertToLily 受け取ったUTF-8のstringをリリイ文字に変換し、その結果を返します。
// リリイ文字はShiftJISに変換後、MacRomanに変換した結果の文字列です。
// 受け取った文字列に半角の英数字が含まれておりconvertAlphanumeric==true の場合、対象の文字列を全角に変換します。
// そのため、再変換した場合、すべての文字が全角になります。
// また、絵文字など、ShiftJISが対応していない文字を受け取った場合空文字を返します。
// 戻り値の２つ目はerror型ですが、現状errorは(恐らく)発生しません。
func ConvertToLily(s string, convertAlphanumeric bool) (string, error) {
	// 全角に変換
	if convertAlphanumeric {
		s = width.Widen.String(s)
	}

	r := utf8ToShiftJIS(s)
	result, err := shiftJISToMacRoman(r)
	if err != nil {
		xerrors.Errorf("failed to convert to Lily lang: %w", err)
	}

	return result, nil
}

// ConvertFromLily 受け取ったリリイ文字のstringをUTF-8に変換し、その結果を返します。
// リリイ文字がShiftJISを採用しているため、元の文字列に半角の英数字が含まれている場合、その文字列は全角に変換されます。
// また、ShiftJISが対応していない文字を受け取った場合、空文字を返します。
// 戻り値の２つ目はerror型ですが、現状errorは(恐らく)発生しません。
func ConvertFromLily(s string) (string, error) {
	r := macRomanToShiftJIS(s)
	result, err := shiftJISToUTF8(r)
	if err != nil {
		xerrors.Errorf("failed to convert to UTF-8: %w", err)
	}
	return result, nil
}
