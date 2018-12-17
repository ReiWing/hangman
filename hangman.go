package main

import (
	"fmt"
	s "strings"

	k "github.com/reiwing/hangman/packages/keyboard"
)

func main() {
	usedSymbols := ""
	guessesCustom, _ := k.GetInt("Enter guesses: ")
	newWord, _ := k.ReturnWord("Enter new word: ")

	for guesses := 0; guesses < guessesCustom; {

		newSymbol, _ := k.ReturnWord("Enter new symbol: ")
		if s.Contains(newWord, newSymbol) == true {
			fmt.Println("tryam")
			usedSymbols = usedSymbols + newSymbol
			fmt.Println("Used symbols:", usedSymbols)
		} else {
			fmt.Println("bol'")
			usedSymbols = usedSymbols + newSymbol
			fmt.Println("Used symbols: ", usedSymbols)
			guesses++
			fmt.Println("You have ", guessesCustom-guesses, " guesses left")
		}

	}

}
