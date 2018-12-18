package main

import (
	"fmt"
	s "strings"

	k "github.com/reiwing/hangman/packages/keyboard"
)

func main() {
	usedSymbols := ""
	guessesCustom, _ := k.GetInt("Enter guesses: ")
	newWord, _ := k.GetWord("Enter new word: ")
	length := len(newWord)
	fmt.Println(length)

	for guesses := 0; guesses < guessesCustom; {

		newSymbol, _ := k.GetWord("Enter new symbol: ")
		if s.Contains(newWord, newSymbol) == true {
			fmt.Println("tryam")
		} else {
			fmt.Println("bol'")
			usedSymbols = usedSymbols + newSymbol
			fmt.Println("Used symbols: ", usedSymbols)
			guesses++
			fmt.Println("You have ", guessesCustom-guesses, " guesses left")
		}

	}

}
