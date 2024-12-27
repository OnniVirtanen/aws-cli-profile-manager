package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)

	if argCount != 2 {
		fmt.Println("Not a valid argument.")
	}

	profile := os.Args[1]
	os.Setenv("AWS_DEFAULT_PROFILE", profile)
	fmt.Printf("AWS_DEFAULT_PROFILE set to %s\n", profile)
}
