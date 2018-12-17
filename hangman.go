package main

import (
	"fmt"
	s "strings"

	"github.com/reiwing/hangman/functions"
)

func main() {
	newWord, _ := keyboard.EnterWord("Enter new word: ")
	newSymbol, _ := keyboard.EnterWord("Enter new symbol: ")

	if s.Contains(newWord, newSymbol) == true {
		fmt.Println("tryam")
	} else {
		fmt.Println("bol'")
	}
}
