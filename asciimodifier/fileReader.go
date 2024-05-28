package asciimodifier

import (
	// "fmt"

	"fmt"
	"log"
	"os"
	"strings"
)

/*
This function receives the commandline arguments
and reads the selected banner file, then matches
the commandline arguments to respective ascii representative
from the banner file.
*/
func StandardReader (words []string) string {
	var flag string

	if len(os.Args) == 3 {
		flag = os.Args[2]
		if flag != "shadow" && flag != "thinkertoy" && flag != "standard" {
			return ""
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		os.Exit(1)
	}

	return standardReader(words, flag)
}

func standardReader(words []string, flag string) string {
	

	var patternFile string

	switch flag {
	case "shadow":
		patternFile = "shadow.txt"
	case "thinkertoy":
		patternFile = "thinkertoy.txt"
	case "standard":
		patternFile = "standard.txt"

	}

	var str string
	fileContent, err := os.ReadFile(patternFile)
	if err != nil {
		log.Fatal(err)
	}

	var strArray []string
	if patternFile == "thinkertoy.txt" {
		strArray = strings.Split(string(fileContent), "\r\n")
	} else {
		strArray = strings.Split(string(fileContent), "\n")
	}

	if len(strArray) != 856 {
		fmt.Println("Error: The file is modified")
		os.Exit(1)
	}

	for _, word := range words {
		if word == "" || word == "\n" {
			str = str + "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				startPoint := int(((char - 32) * 9) + 1)
				str = str + strArray[startPoint+i]
			}
			str = str + "\n"
		}
	}
	return str
}
