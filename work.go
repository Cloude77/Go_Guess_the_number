package main

import (

	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("work.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
}
