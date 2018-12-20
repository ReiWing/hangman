package main

import (
	"fmt"
	s "strings"

	k "github.com/reiwing/hangman/packages/keyboard"
)

func main() {
	guessesCustom, _ := k.GetInt("Enter available guesses: ")
	newWord, _ := k.GetWord("Enter new word: ")
	k.ClearTerminal()
	fmt.Println(k.ReplaceAll(newWord, "_ "))

	usedSymbols := ""
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
