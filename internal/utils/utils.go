package utils

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadInt(reader *bufio.Reader) int {
	input := ReadLine(reader)
	n, _ := strconv.Atoi(input)
	return n
}
