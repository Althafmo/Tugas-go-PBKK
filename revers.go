package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan minimal 3 kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	if len(words) < 3 {
		fmt.Println("Masukkan minimal 3 kata.")
		return
	}

	reversedWords := make([]string, len(words))
	for i, word := range words {
		runes := []rune(word)
		for j := len(runes) - 1; j >= 0; j-- {
			reversedWords[i] += string(runes[j])
		}
	}

	fmt.Println(strings.Join(reversedWords, " "))
}
