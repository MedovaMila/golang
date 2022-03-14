package main

import (
	"bufio" // пакет для функции Reader
	"fmt"
	"log" // для вывода ошибки
	"os"  // для чтения с клавиатуры
	"strconv"
	"strings"
)

func main() {

	fmt.Print("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
