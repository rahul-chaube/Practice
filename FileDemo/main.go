package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Testing"), 0644)
	if err != nil {
		fmt.Println("Error occured while creating file ", err)
	}
	fmt.Println("File is created success fdully ")
}
