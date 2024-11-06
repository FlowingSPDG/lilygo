package lilygo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/lilygo"
)

func TestConvertToLily(t *testing.T) {
	type tt struct {
		name                string
		input               string
		convertAlphanumeric bool
		expected            string
		err                 error
	}

	tests := []tt{
		{
			name:                "Test1",
			input:               "あいうえお",
			convertAlphanumeric: false,
			expected:            "Ç†Ç¢Ç§Ç¶Ç®",
			err:                 nil,
		},
		{
			name:                "Test2",
			input:               "かきくけこ",
			convertAlphanumeric: false,
			expected:            "Ç©Ç´Ç≠ÇØÇ±",
			err:                 nil,
		},
		{
			name:                "Test3",
			input:               "さしすせそ",
			convertAlphanumeric: false,
			expected:            "Ç≥ÇµÇ∑ÇπÇª",
			err:                 nil,
		},
		{
			name:                "半角の英数字が含まれている場合_1",
			input:               "ether",
			convertAlphanumeric: false,
			expected:            "ether",
			err:                 nil,
		},
		{
			name:                "半角の英数字が含まれている場合_2",
			input:               "ether",
			convertAlphanumeric: true,
			expected:            "ÇÖÇîÇàÇÖÇí",
			err:                 nil,
		},
		{
			name:                "ShiftJIS未対応の文字が含まれている場合_1",
			input:               "𠮟塡剝頰",
			convertAlphanumeric: false,
			expected:            "",
			err:                 nil,
		},
		{
			name:                "ShiftJIS未対応の文字が含まれている場合_2",
			input:               "👍ZXCV123",
			convertAlphanumeric: false,
			expected:            "",
			err:                 nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			asserts := assert.New(t)

			actual, err := lilygo.ConvertToLily(tc.input, tc.convertAlphanumeric)
			asserts.Equal(tc.expected, actual)
			asserts.Equal(tc.err, err)
		})
	}
}

func TestConvertToLilyWithOriginal(t *testing.T) {
	type tt struct {
		name                string
		input               string
		convertAlphanumeric bool
		expected            string
		err                 error
	}

	tests := []tt{
		{
			name:                "Test1",
			input:               "あいうえお",
			convertAlphanumeric: false,
			expected:            "Ç†Ç¢Ç§Ç¶Ç®\n《あいうえお》",
			err:                 nil,
		},
		{
			name:                "Test2",
			input:               "かきくけこ",
			convertAlphanumeric: false,
			expected:            "Ç©Ç´Ç≠ÇØÇ±\n《かきくけこ》",
			err:                 nil,
		},
		{
			name:                "Test3",
			input:               "さしすせそ",
			convertAlphanumeric: false,
			expected:            "Ç≥ÇµÇ∑ÇπÇª\n《さしすせそ》",
			err:                 nil,
		},
		{
			name:                "半角の英数字が含まれている場合_1",
			input:               "ether",
			convertAlphanumeric: false,
			expected:            "ether\n《ether》",
			err:                 nil,
		},
		{
			name:                "半角の英数字が含まれている場合_2",
			input:               "ether",
			convertAlphanumeric: true,
			expected:            "ÇÖÇîÇàÇÖÇí\n《ether》",
			err:                 nil,
		},
		{
			name:                "ShiftJIS未対応の文字が含まれている場合_1",
			input:               "𠮟塡剝頰",
			convertAlphanumeric: false,
			expected:            "\n《𠮟塡剝頰》",
			err:                 nil,
		},
		{
			name:                "ShiftJIS未対応の文字が含まれている場合_2",
			input:               "👍ZXCV123",
			convertAlphanumeric: false,
			expected:            "\n《👍ZXCV123》",
			err:                 nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			asserts := assert.New(t)

			actual, err := lilygo.ConvertToLilyWithOriginal(tc.input, tc.convertAlphanumeric)
			asserts.Equal(tc.expected, actual)
			asserts.Equal(tc.err, err)
		})
	}
}

func TestConvertFromLily(t *testing.T) {
	type tt struct {
		name     string
		input    string
		expected string
		err      error
	}

	tests := []tt{
		{
			name:     "Test1",
			input:    "Ç†Ç¢Ç§Ç¶Ç®",
			expected: "あいうえお",
			err:      nil,
		},
		{
			name:     "Test2",
			input:    "Ç©Ç´Ç≠ÇØÇ±",
			expected: "かきくけこ",
			err:      nil,
		},
		{
			name:     "Test3",
			input:    "Ç≥ÇµÇ∑ÇπÇª",
			expected: "さしすせそ",
			err:      nil,
		},
		{
			name:     "ShiftJIS未対応の文字が含まれている場合_1",
			input:    "👍",
			expected: "",
			err:      nil,
		},
		{
			name:     "ShiftJIS未対応の文字が含まれている場合_2",
			input:    "𠮟塡剝頰",
			expected: "",
			err:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			asserts := assert.New(t)

			actual, err := lilygo.ConvertFromLily(tc.input)
			asserts.Equal(tc.expected, actual)
			asserts.Equal(tc.err, err)
		})
	}
}
