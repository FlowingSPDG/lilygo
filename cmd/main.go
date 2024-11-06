package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/FlowingSPDG/lilygo"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for {
			fmt.Printf("> ")
			scanner.Scan()

			s := scanner.Text()

			result, err := lilygo.ConvertToLilyWithOriginal(s, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		}
	}()

	<-ctx.Done()
	fmt.Printf("\nBye.\n")
}
