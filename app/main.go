package main

import (
	"betypes/beerrors"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	var e beerrors.BeUnknownError = fmt.Errorf("unknown error")
	fmt.Println(e)
}
