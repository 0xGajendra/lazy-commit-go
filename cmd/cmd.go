package cmd

import (
	"fmt"
	"lazy-commit-go/internal/config"
	"lazy-commit-go/internal/git"
	"lazy-commit-go/internal/groq"
	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
)

func Run(){
	if(!git.IsGitRepo()){
		fmt.Println("working directory is not a git repository")
		prompt := promptui.Select{
		Label: "would you like to initialize it as a git repository?",
		Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}
		if result == "Yes" {
			err = git.InitRepo()
			if err != nil {
				fmt.Println("Error initializing git repository:", err)
				return
			}
			fmt.Println("Git repository initialized successfully.")
		} else {
			fmt.Println("Exiting without initializing git repository.")
			return
		}

		
		return
	}
	var err error
	files, err := git.GetChangedFiles()
	if err != nil {
		fmt.Println("Error getting changed files:", err)
		return
	}
	fmt.Println("Changed files:")
	for _, file := range files {
		fmt.Println("- ", file)
	}

	files = append([]string{"All files"}, files...)
	filesSelected := []string{}
	filePrompt := &survey.MultiSelect{
		Message: "What files do you want to commit?",
		Options: files,
	}
	survey.AskOne(filePrompt, &filesSelected)

	if len(filesSelected) > 0 && filesSelected[0] == "All files" {
		err = git.StageSelectedFiles([]string{"."})
	} else {
		err = git.StageSelectedFiles(filesSelected)
	}
	if err != nil {
		fmt.Println("Error staging files:", err)
		return
	}
	var groqAPI string
	if(!config.IsConfigFileExist()){
		fmt.Println("Enter groq api key: ")		
		fmt.Scan(&groqAPI)
		config.SaveAPIKey(groqAPI)
	}
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
		Label:   "Edit your commit message(enter to keep unchanged)",
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

	branch, err:= git.GetCurrentBranch()
	if err != nil{
		fmt.Println("Error getting current branch:", err)
		return
	}
	fmt.Printf("Changes committed to branch %s with message: %s\n", branch, result)
	prompt = promptui.Select{
		Label: "Do you want to push the changes? (y/n)",
		Items: []string{"y", "n"},
	}
	_, result, err = prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}
	if result == "y" {
		err = git.PushChanges()
		if err != nil {
			fmt.Println("Error pushing changes:", err)
			return
		}
		fmt.Println("Changes pushed successfully.")
	} else {
		fmt.Println("Changes committed but not pushed.")
	}
}