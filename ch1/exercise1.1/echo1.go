package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//file name will also be printed as file name is first arg in os.Args
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
