package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main()  {
	root, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	newUser, err := os.Stat(root + "/.lzc_config")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(newUser)

	var groqAPI string
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Enter groq api key: ")
		fmt.Scan(&groqAPI)
		err := os.WriteFile(root + "/.lzc_config", []byte(groqAPI), 0644)
		if err != nil{
			fmt.Println("There is a fucking error in setting up groq api key first time")
		}
		
	} else {
		res, err :=os.ReadFile(root + "/.lzc_config")
		groqAPI=string(res)
		if err != nil{
			fmt.Println("Error in reading file")
		}
		println("groqAPI", groqAPI)
		out, err := exec.Command("git", "diff").Output()

		if err != nil {
			fmt.Println("there is an error in running git diff")
		}

		println(string(out))
	}

}
