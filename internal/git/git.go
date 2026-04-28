package git

import (
	"fmt"
	"os/exec"
)

func GetStagedDiff() (string, error) {
	out, err := exec.Command("git", "diff", "--staged").Output()

	if err != nil {
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}

func Commit(message string) error {
	_, err := exec.Command("git", "commit", "-m", message).CombinedOutput()

	if err != nil {
		return err
	}
	return nil
}