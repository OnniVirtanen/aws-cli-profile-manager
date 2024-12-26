package main

import (
	"log"
	"os"
)

func main() {
	argCount := len(os.Args)

	const helpText = "See 'awscpm --help'."

	if argCount > 3 || argCount < 1 {
		log.Fatalf("Unvalid amount of arguments. %s\n", helpText)
	}

	if os.Args[1] == "--help" {
		log.Printf("Available commands are\n'awscpm --help'\n'awscpm df <profile>'\n'awscpm ls'")
	} else if os.Args[1] == "df" {
		log.Printf("Setting default profile to %s.", os.Args[2])
	} else if os.Args[1] == "ls" && argCount == 2 {
		log.Println("Listing all available AWS profiles.")
	} else {
		log.Fatalf("Not valid argument. %s\n", helpText)
	}

	/*
		fmt.Printf("Number of arguments: %d\n", argCount)
		fmt.Println("Arguments:", os.Args)
		fmt.Println("Argument [1]", os.Args[1])
	*/
}
