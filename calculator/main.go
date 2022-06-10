package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func parse(value string) int {
	data, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (Calc) operate(input string, operator string) int {
	values := strings.Split(input, operator)
	var result int
	for index, value := range values {
		data := parse(value)
		if index == 0 {
			result = data
		} else {
			switch operator {
			case "+":
				result += data
			case "-":
				result -= data
			case "*":
				result *= data
			case "/":
				result /= data
			default:
				log.Fatalln("Invalid operator", operator)
			}
		}
	}

	return result
}

func (Calc) scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var calc Calc
	operation := calc.scanner()
	result := calc.operate(operation, "+")
	fmt.Println(operation, "=", result)
}
