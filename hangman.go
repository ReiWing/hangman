package main

import (
	"fmt"

	"github.com/reiwing/hangman/functions"
)

func main() {
	newWord, _ := keyboard.EnterWord()
	fmt.Println(newWord)
}
