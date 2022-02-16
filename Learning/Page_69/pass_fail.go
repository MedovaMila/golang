package main

import (
	"bufio" // пакет для функции Reader
	"fmt"
	"log" // для вывода ошибки
	"os"  // для чтения с клавиатуры
)

func main() {

	fmt.Print("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)

}
