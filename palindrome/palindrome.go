package main 

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	process := strings.ReplaceAll(strings.ToLower(text), " ", "");

	length := len(process);
	for i := 0; i < length / 2; i++ {
		if process[i]  != process[length - 1 - i] {
			return false
		}
	}

	return true
}

func main() {
	var text string
	fmt.Print("Palindrom kontrol metni gir: ");
	fmt.Scan(&text);

	if isPalindrome(text) {
		fmt.Printf("%s metni bir palindromdur", text);
	} else {
		fmt.Printf("%s metni palindrom değildir", text);
	}
}