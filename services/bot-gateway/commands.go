package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseAddCommand(text string) (string, int, error) {
	words := strings.Fields(text)

	if len(words) < 3 {
		return "", 0, fmt.Errorf("invalid command format")
	}

	if words[0] != "/add" {
		return "", 0, fmt.Errorf("invalid command")
	}

	productName := strings.Join(words[1:len(words)-1], " ")

	quantity, err := strconv.Atoi(words[len(words)-1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid quantity")
	}

	if quantity <= 0 {
		return "", 0, fmt.Errorf("quantity must be positive")
	}

	return productName, quantity, nil
}
