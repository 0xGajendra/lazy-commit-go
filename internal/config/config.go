package config

import (
	"os"
)
func IsConfigFileExist() bool {
	root, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	_, err = os.Stat(root + "/.lzc_config")
	return !os.IsNotExist(err)
}

func LoadAPIKey() (string, error) {
    root, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
	res, err :=os.ReadFile(root + "/.lzc_config")
	if err != nil{
		return "", err
	}
	return string(res), nil
}

func SaveAPIKey(groqAPI string) error {
	root, err := os.UserHomeDir()

	if err != nil {
		return err
	}
	err = os.WriteFile(root+"/.lzc_config", []byte(groqAPI), 0644)
	return err
}