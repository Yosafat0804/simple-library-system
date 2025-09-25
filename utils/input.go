package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func InputInt(prompt string) int {
	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err == nil {
			return num
		}
		fmt.Println("Input tidak valid, masukkan angka!")
	}
}
