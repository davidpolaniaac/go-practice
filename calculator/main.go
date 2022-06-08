package main

import (
	"bufio"
	"fmt"
	"log"
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
		data, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalln(err)
		}
		result += data
	}
	fmt.Println("Total:", result)
}
