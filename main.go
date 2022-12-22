package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	BacaInputan()
}

func BacaInputan() {
	c := gogpt.NewClient(os.Getenv("CHAT_GPT"))
Loop:
	for {
		scan := bufio.NewScanner(os.Stdin)
		repeat := bufio.NewScanner(os.Stdin)
		color.Red("Silakan Dek Mau Tanya Apa Asal kan jangan jodoh! ")
		fmt.Print("=>")

		for scan.Scan() {
			ask := scan.Text()
			CallAI(ask+"\n", c)
			color.Red("Mau Lanjut ya/tidak")
			fmt.Print("=>")
			for repeat.Scan() {
				switch repeat.Text() {
				case "ya":
					goto Loop
				case "tidak":
					break Loop
				default:
					color.Yellow("Command Tidak Di Temukan")
					break Loop
				}
			}

		}
	}
}

func CallAI(prompt string, c *gogpt.Client) {

	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 4000,
		Prompt:    prompt,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}

	for _, val := range resp.Choices {
		color.Yellow(val.Text)
	}

}
