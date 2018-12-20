// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

// GetWord func allow user to inter new word
func GetWord(request string) (string, error) {
	fmt.Println(request)
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		return "0", err
	}
	word = s.TrimSpace(s.ToUpper(word))
	return word, nil
}

//GetInt func return user number
func GetInt(request string) (int, error) {
	fmt.Println(request)
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	number = s.TrimSpace(s.ToUpper(number))
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	return numberInt, nil
}

// ReplaceAll func replace all symbols in a string
func ReplaceAll(word string, symbol string) string {
	re := regexp.MustCompile(".")
	reString := re.ReplaceAllString(word, symbol)
	return reString
}

//ClearTerminal just clear terminal with ANSI escape codes
func ClearTerminal() {
	print("\033[H\033[2J")
}
