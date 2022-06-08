package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	values := strings.Split(operation, "+")
	var result int = 0
	for _, value := range values {
		data, _ := strconv.Atoi(value)
		result += data
	}
	fmt.Println("Total:", result)
}
