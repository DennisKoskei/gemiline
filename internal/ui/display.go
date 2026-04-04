package ui

import (
	"fmt"
	"strings"
)

func PrintResponse(text string) {
	fmt.Println("Gemiline:")
	fmt.Println("---")

	for line := range strings.SplitSeq(text, "\n") {
		fmt.Println("|", line)
	}

	fmt.Println("---")
}
