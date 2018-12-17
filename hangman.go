package main

import (
	"fmt"
	s "strings"

	"github.com/reiwing/hangman/functions"
)

func main() {
	newWord, _ := keyboard.EnterWord()
	fmt.Println(newWord)

	newSymbol, _ := keyboard.EnterSymbol()
	fmt.Println(newSymbol)

	if s.Contains(newWord, newSymbol) == true {
		fmt.Println("tryam")
	} else {
		fmt.Println("bol'")
	}
}
