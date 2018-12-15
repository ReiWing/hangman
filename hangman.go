package main

import (
	"fmt"

	"github.com/reiwing/hangman/functions"
)

func main() {
	abc := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	newWord := keyboard.EnterWord()
	fmt.Println(abc, newWord)
}
