package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	num := os.Args[1]
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	if i%2 != 0 {
		fmt.Printf("%d is odd\n", i)
	} else {
		fmt.Printf("%d is even\n", i)
	}
}
