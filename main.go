package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/onni.virtanen/aws.cli.profile.manager/logic"
)

const VERSION = "1.0"
const UNVALID_ARGUMENT = "Not valid argument. See 'apm --help'."

func main() {
	argCount := len(os.Args)

	if argCount == 1 {
		log.Fatalln(UNVALID_ARGUMENT)
	}

	if os.Args[1] == "--help" && argCount == 2 {
		c1 := "  apm --help                                             list available commands"
		c2 := "  apm default <profile>                                  set default profile"
		c3 := "  apm ls                                                 list available profiles"
		c4 := "  apm show default                                       show current default profile"
		c5 := "  apm add <profile> <access_key_id> <secret_access_key>  add profile"
		c6 := "  apm rmv <profile>                                      remove profile"
		c7 := "  apm --v                                                show current version"
		log.Printf("These are all available apm commands:\n%s\n%s\n%s\n%s\n%s\n%s\n%s", c1, c2, c3, c4, c5, c6, c7)
	} else if os.Args[1] == "default" && argCount == 3 {
		err := logic.SetDefault(os.Args[2])
		if err != nil {
			log.Fatalf("Could not set default profile: %s", err)
		}
		log.Printf("Default profile set successfully to %s.", os.Args[2])
	} else if os.Args[1] == "ls" && argCount == 2 {
		log.Println("Listing all available AWS profiles:")
		profiles, err := logic.GetProfiles()
		if err != nil {
			log.Fatalf("Could not list all available AWS profiles: %s", err)
		}
		splitProfiles := strings.Split(profiles, "\n")
		for _, element := range splitProfiles {
			columns := strings.Split(element, ";")
			if len(columns) > 2 {
				fmt.Println(columns[1])
			}
		}
	} else if os.Args[1] == "add" && (argCount == 5 || argCount == 6) {
		token := ""
		if argCount == 6 {
			token = os.Args[5]
		}
		err := logic.AddProfile([4]string{os.Args[2], os.Args[3], os.Args[4], token})
		if err != nil {
			log.Fatalf("Could not add profile: %s", err)
		}
		fmt.Printf("Added AWS profile: %s\n", os.Args[2])
	} else if os.Args[1] == "show" && os.Args[2] == "default" && argCount == 3 {
		data, err := logic.GetDefault()
		if err != nil {
			log.Fatalf("Could not show default profile: %s", err)
		}
		after, found := strings.CutPrefix(data, "default;")
		if found {
			fmt.Printf("AWS CLI default profile is %s\n", after)
		}
	} else if os.Args[1] == "rmv" && argCount == 3 {
		err := logic.RemoveProfile(os.Args[2])
		if err != nil {
			log.Fatalf("Could not remve AWS profile: %s", err)
		}
		log.Printf("Removed profile %s successfully.", os.Args[2])
	} else if os.Args[1] == "--v" && argCount == 2 {
		log.Printf("apm version %s\n", VERSION)
	} else {
		log.Fatalln(UNVALID_ARGUMENT)
	}
}
