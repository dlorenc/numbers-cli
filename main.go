package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	odd "github.com/dlorenc/is-odd"
)

func main() {
	num := os.Args[1]
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	if odd.IsOdd(i) {
		fmt.Printf("%d is odd\n", i)
	} else {
		fmt.Printf("%d is even\n", i)
	}
}
