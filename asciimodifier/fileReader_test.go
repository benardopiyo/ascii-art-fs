package asciimodifier

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestStandardReader(t *testing.T) {
	t1, err := os.ReadFile("test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	t2, err := os.ReadFile("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	t3, err := os.ReadFile("test3.txt")
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		input    string
		expected string
	}{
		{`hello`, string(t1)},
		{`Hello\n\nThere`, string(t2)},
		{`Hello\nThere`, string(t3)},
	}

	for _, test := range tests {
		words := strings.Split(test.input, "\\n")
		output := standardReader(words, "standard")
		if output != test.expected {
			log.Fatal("error occured")
		}
	}
}
