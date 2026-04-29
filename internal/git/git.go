package git

import (
	"fmt"
	"os/exec"
	"strings"
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

func GetChangedFiles() ([]string, error){
	out, err := exec.Command("git", "status", "--short").Output()
	if err!=nil{
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	var files []string
	for _, line := range lines {
		if(len(line))>3{
			files = append(files, strings.TrimSpace(line[3:]))
		}
	}
	return files, nil
}

func StageSelectedFiles(files []string) error {
	args:=append([]string{"add"}, files...)
	err := exec.Command("git", args...).Run()
	if err != nil {
		return err
	}
	return nil
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

func InitRepo() error{
	err := exec.Command("git", "init").Run()
	return err
}