package main

import (
	"fmt"
	"io"
	"os"

	"github.com/emojive/emoji"
	cli "github.com/jawher/mow.cli"
)

type Config struct {
	FontPath   string
	OutputPath string
	Text       []string
}

func run(cfg Config) {
	fmt.Println(cfg)
	var writer io.Writer
	var err error

	if cfg.OutputPath != "" {
		if writer, err = os.Create(cfg.OutputPath); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		writer = os.Stdout
	}

	for _, word := range cfg.Text {
		for _, letter := range word {
			err, emojiStr := emoji.GenerateEmojiLetter(letter, "💚", "❤️", 5)
			if err != nil {
				fmt.Fprintln(writer, "Error:", err)
			} else {
				fmt.Fprintln(writer, emojiStr)
			}
		}

	}
}

func main() {
	var cfg Config

	app := cli.App("emojive", "Generate emoji letters")

	app.Spec = "[-f=<font>] [-w=<file>] TEXT..."

	app.StringOptPtr(&cfg.FontPath, "f font", "test", "Font path")
	app.StringOptPtr(&cfg.OutputPath, "w file", "", "Output file path")
	app.StringsArgPtr(&cfg.Text, "TEXT", nil, "Text to transform")

	app.Action = func() {
		run(cfg)
	}

	app.Run(os.Args)
}
