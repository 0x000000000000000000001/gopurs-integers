package main

import (
	"fmt"
	"strconv"
)

func main() {
	val, err := strconv.ParseInt("2147483648", 10, 32)
	fmt.Printf("val=%v, err=%v\n", val, err)
}
