package main

import (
	"fmt"
	"log"
	"os"

	"github.com/onni.virtanen/aws.cli.profile.manager/logic"
)

func main() {
	argCount := len(os.Args)

	const helpText = "See 'awscpm --help'."

	if os.Args[1] == "--help" {
		log.Printf("Available commands are\n'awscpm --help'\n'awscpm df <profile>'\n'awscpm ls'")
	} else if os.Args[1] == "df" {
		err := logic.SetDefault(os.Args[2])
		if err != nil {
			log.Fatalf("Could not set default profile: %s", err)
		}
		log.Printf("Default profile set successfully to <%s>.", os.Args[2])
	} else if os.Args[1] == "ls" && argCount == 2 {
		log.Println("Listing all available AWS profiles.")
		profiles, err := logic.GetProfiles()
		if err != nil {
			log.Fatalf("Could not list all available AWS profiles: %s", err)
		}
		fmt.Println(profiles)
	} else if os.Args[1] == "ap" && argCount == 5 {
		log.Println("got here")
		logic.AddProfile([3]string{os.Args[2], os.Args[3], os.Args[4]})
	} else {
		log.Fatalf("Not valid argument. %s\n", helpText)
	}
}
