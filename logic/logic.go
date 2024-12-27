package logic

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const PROFILE_FILE = "/.aws/profiles"
const DEFAULT_FILE = "/.aws/default"
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
	if err != nil {
		return err
	}
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

	pdata, err := GetProfiles()
	if err != nil {
		return err
	}
	profiles := strings.Split(pdata, "\n")
	aws_access_key_id := ""
	aws_secret_access_key := ""
	profileFound := false
	for _, element := range profiles {
		if !strings.Contains(element, profile) {
			continue
		}
		profileFound = true
		columns := strings.Split(element, ";")
		aws_access_key_id = columns[2]
		aws_secret_access_key = columns[3]
	}
	if !profileFound {
		return errors.New("no profile found with specified name")
	}

	credentials := fmt.Sprintf("[default]\naws_access_key_id=%s\naws_secret_access_key=%s", aws_access_key_id, aws_secret_access_key)
	err = os.WriteFile(dir+CREDENTIALS_FILE, []byte(credentials), 0644)
	if err != nil {
		return err
	}

	data := fmt.Sprintf("default;%s", profile)
	err = os.WriteFile(dir+DEFAULT_FILE, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetDefault() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(dir + DEFAULT_FILE)
	if err != nil {
		return "", errors.New("no default profile")
	}

	return string(data), nil
}

func RemoveProfile(profile string) error {
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	pdata, err := GetProfiles()
	if err != nil {
		return err
	}
	profiles := strings.Split(pdata, "\n")
	profilesModified := ""
	profileFound := false
	for _, element := range profiles {
		if strings.Contains(element, profile) {
			profileFound = true
			continue
		}
		profilesModified += element + "\n"
	}
	if !profileFound {
		return errors.New("no profile found with specified name")
	}

	err = os.WriteFile(dir+PROFILE_FILE, []byte(profilesModified), 0644)
	if err != nil {
		return err
	}
	return nil
}
