// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

// EnterWord func allow user inter new word
func EnterWord(request string) (string, error) {
	fmt.Println(request)
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		return "0", err
	}
	word = s.TrimSpace(s.ToUpper(word))
	return word, nil
}
