package main

import (
	"strconv"
	"bufio"
	"os"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n","", -1)
	return text
}

func readNumberInput() int {
	text := readLine()
	number, _ := strconv.Atoi(text)
	return number
}
