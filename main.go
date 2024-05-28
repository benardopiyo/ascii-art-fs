package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art/asciimodifier"
)

/*
The main function checks if the commandline arguments are two
and works with the argument at index 1. Checks for special characters
except for '\n' where we print a newline. Finally the function splits
the input with a literal '\n'.
*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: Provide a string and a flag")
	}

	strWord := os.Args[1]

	strWord = strings.ReplaceAll(strWord, "\\r", "\r")
	strWord = strings.ReplaceAll(strWord, "\\a", "\a")
	strWord = strings.ReplaceAll(strWord, "\\b", "\b")
	strWord = strings.ReplaceAll(strWord, "\\v", "\v")
	strWord = strings.ReplaceAll(strWord, "\\t", "\t")
	strWord = strings.ReplaceAll(strWord, "\\f", "\f")

	for _, char := range strWord {
		if char >= 0 && char <= 31 || char > 126 {
			fmt.Println("Error: non-printable character ")
			return
		}
	}

	if strWord == "\\n" {
		fmt.Println()
		return

	} else if strWord == "" {
		return
	}

	words := strings.Split(strWord, "\\n")
	output1 := asciimodifier.StandardReader(words)
	fmt.Print(output1)
}
