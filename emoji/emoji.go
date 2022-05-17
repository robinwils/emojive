package emoji

import (
	"fmt"
)

type Emoji struct {
	Number int
	Name   string
	Code   rune
}

var lettersMatrices map[rune][5][5]int = map[rune][5][5]int{
	'A': {
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
	},
	'B': {
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	},
	'C': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	},
	'D': {
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 0},
	},
	'E': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	},
	'F': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'G': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	},
	'H': {
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	},
	'I': {
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'J': {
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 1, 0, 0},
	},
	'K': {
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 0, 1, 0},
	},
	'L': {
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	},
	'M': {
		{1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	},
	'N': {
		{1, 0, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1},
	},
	'O': {
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	},
	'P': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'Q': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 1, 1},
		{1, 1, 1, 1, 1},
	},
	'R': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 1},
	},
	'S': {
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	},
	'T': {
		{1, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'U': {
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	},
	'V': {
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
	},
	'W': {
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 1},
	},
	'X': {
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
	},
	'Y': {
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
	},
	'Z': {
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	},
}

func GenerateEmojiLetter(letter rune, fg, bg string, size int) (error, string) {
	emojiString := ""
	letterMatrice, ok := lettersMatrices[letter]
	if !ok {
		return fmt.Errorf("letter '%c' unknown", letter), ""
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if letterMatrice[i][j] == 0 {
				emojiString += string(bg)
			} else {
				emojiString += string(fg)
			}
		}
		emojiString += "\n"
	}
	return nil, emojiString
}
