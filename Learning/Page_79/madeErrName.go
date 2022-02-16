package main

import (
	"fmt"
)

func main() {
	var count int = 12
	var surfix string = "minutes of bonus footage"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol")
	fmt.Println(count, surfix, "on", format, languages)
}
