// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EnterWord func allow user inter new word
func EnterWord() ([]string, error) {
	fmt.Println("Enter new word: ")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	// word, err := reader.ReadString('\n')
	// if err != nil {
	// 	return err
	// }
	s := strings.Split(word, "")
	fmt.Println(s)
	return s, nil
}
