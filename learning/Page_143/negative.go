package main

import "fmt"

func negative(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	truth := true
	fmt.Println(truth)
	negative(&truth)
	fmt.Println(truth)
	lies := false
	negative(&lies)
	fmt.Println(lies)
}
