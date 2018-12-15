// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EnterWord func allow user inter new word
func EnterWord() []string {
	fmt.Println("Enter new word: ")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	// word, err := reader.ReadString('\n')
	// if err != nil {
	// 	return err
	// }
	word = strings.ToUpper(word)
	array := strings.Split(word, "")
	return array
}

// Remove func delets elements from array
// func Remove(a []string, b []string) []string {
// 	fmt.Println(a)
// 	for _, ip := range a {
// 		if
// 		// if net.ParseIP(ip).To4() != nil {
// 		// 	b = append(b, ip)
// 		}
// 	}
// 	return b
// }
