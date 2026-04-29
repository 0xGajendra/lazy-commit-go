package cmd

import (
	"fmt"
	"lazy-commit-go/internal/config"
	"lazy-commit-go/internal/git"
	"lazy-commit-go/internal/groq"

	"github.com/manifoldco/promptui"
)

func Run(){
	var groqAPI string
	if(!config.IsConfigFileExist()){
		fmt.Println("Enter groq api key: ")		
		fmt.Scan(&groqAPI)
		config.SaveAPIKey(groqAPI)
	}
	var err error
	groqAPI, err = config.LoadAPIKey()
	if err != nil {
		fmt.Println("Error loading API key:", err)
		return
	}

	diff, err := git.GetStagedDiff()

	if err != nil {
		fmt.Println("Error getting staged diff:", err)
		return
	}

	commitMessages, err := groq.GetCommitMessages(diff, groqAPI)

	if err != nil {
		fmt.Println("Error getting commit messages:", err)
		return
	}
	for _, message := range commitMessages {
		fmt.Println(message)
	}
	prompt := promptui.Select{
		Label: "Select commit message",
		Items: commitMessages,
	}

	_, result, err := prompt.Run()	

	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}


	editPrompt := promptui.Prompt{
		Label:   "Edit your commit message",
		Default: result,
	}
	result, err = editPrompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	err = git.Commit(result)
	
	if err != nil {
		fmt.Println("there is an error in running git commit", err)
	}
}