package main

import (
	"fmt"

	"github.com/hsiuhsiu/cgo-exp/mycgolib"
)

func main() {
	fmt.Println(mycgolib.Sum(17, 42))
}
