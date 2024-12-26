package logic

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const PROFILE_FILE = "/awscpm/profiles"
const DEFAULT_FILE = "/awscpm/default"
const CREDENTIALS_FILE = "/.aws/credentials"

func GetProfiles() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(dir + PROFILE_FILE)
	if err != nil {
		return "", errors.New("no profiles added")
	}

	return string(data), nil
}

func AddProfile(arr [3]string) error {
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	data := fmt.Sprintf("profile;%s;%s;%s", arr[0], arr[1], arr[2])
	file, err := os.OpenFile(dir+PROFILE_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	_, err = fmt.Fprintln(file, data)
	if err != nil {
		file.Close()
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func SetDefault(profile string) error {
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// find the profile from profile_file and get the key id and secret key
	pdata, err := GetProfiles()
	if err != nil {
		return err
	}
	profiles := strings.Split(pdata, "\n")
	aws_access_key_id := ""
	aws_secret_access_key := ""
	for _, element := range profiles {
		if !strings.Contains(element, profile) {
			continue
		}
		columns := strings.Split(element, ";")
		aws_access_key_id = columns[2]
		aws_secret_access_key = columns[3]
	}

	// create a new .aws/credentials file to home + credentials_file
	credentials := fmt.Sprintf("[default]\naws_access_key_id=%s\naws_secret_access_key=%s", aws_access_key_id, aws_secret_access_key)
	err = os.WriteFile(dir+CREDENTIALS_FILE, []byte(credentials), 0644)

	// update the default file
	data := fmt.Sprintf("default;%s", profile)
	defFile, err := os.OpenFile(dir+DEFAULT_FILE, os.O_WRONLY|os.O_CREATE, 0644)
	_, err = fmt.Fprintln(defFile, data)
	if err != nil {
		defFile.Close()
		return err
	}
	err = defFile.Close()
	if err != nil {
		return err
	}

	return nil
}
