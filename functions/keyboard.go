// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EnterWord func allow user inter new word
func EnterWord() (string, error) {
	fmt.Println("Enter new word: ")
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		return "0", err
	}
	word = strings.ToUpper(word)
	return word, nil
}
