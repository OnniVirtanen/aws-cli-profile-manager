package main

import (
	"fmt"
	"log"
	"os"

	"github.com/onni.virtanen/aws.cli.profile.manager/logic"
)

const VERSION = "1.0"
const UNVALID_ARGUMENT = "Not valid argument. See 'apm --help'."

func main() {
	argCount := len(os.Args)

	if argCount == 1 {
		log.Fatalln(UNVALID_ARGUMENT)
	}

	if os.Args[1] == "--help" {
		c1 := "'apm --help'                                           list available commands"
		c2 := "'apm df <profile>'                                     set default profile"
		c3 := "'apm ls'                                               list available profiles"
		c4 := "'apm show default'                                     show current default profile"
		c5 := "'apm ap <profile> <access_key_id> <secret_access_key>' add profile"
		c6 := "'apm rp <profile>'                                     remove profile"
		c7 := "'apm --v'                                              show current version"
		log.Printf("Available commands are\n%s\n%s\n%s\n%s\n%s\n%s\n%s", c1, c2, c3, c4, c5, c6, c7)
	} else if os.Args[1] == "df" && argCount == 3 {
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
		logic.AddProfile([3]string{os.Args[2], os.Args[3], os.Args[4]})
	} else if os.Args[1] == "show" && os.Args[2] == "default" && argCount == 3 {
		data, err := logic.GetDefault()
		if err != nil {
			log.Fatalf("Could not show default profile: %s", err)
		}
		fmt.Println(data)
	} else if os.Args[1] == "rp" && argCount == 3 {
		err := logic.RemoveProfile(os.Args[2])
		if err != nil {
			log.Fatalf("Could not remve AWS profile: %s", err)
		}
		log.Printf("Removed profile <%s> successfully.", os.Args[2])
	} else if os.Args[1] == "--v" && argCount == 2 {
		log.Printf("Current version of apm is %s\n", VERSION)
	} else {
		log.Fatalln(UNVALID_ARGUMENT)
	}
}
