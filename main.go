package main

import "fmt"

func main() {

	for _, v := range features {
		fmt.Println(v)
	}

	// go run .
	// go run -tags 'pro' .
	// go run -tags 'pro enter' .
}
