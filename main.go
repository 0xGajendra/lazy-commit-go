package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
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
	}

	fmt.Println("groqAPI", groqAPI)
	out, err := exec.Command("git", "diff").Output()

	if err != nil {
		fmt.Println("there is an error in running git diff")
	}


	result:=getCommitMessages(string(out), groqAPI)

	commitMessages:=strings.Split(result, "\n")
	for _, message := range commitMessages {
		fmt.Println(message)
	}

	prompt := promptui.Select{
		Label: "Select commit message",
		Items: commitMessages,
	}

	_, result, err = prompt.Run()	

	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	fmt.Println("you choose commit message: ", result);

	
	exec.Command("git", "add", ".").Run()

	out, err = exec.Command("git", "commit", "-m", result).CombinedOutput()

	fmt.Println(string(out))
	
	if err != nil {
		fmt.Println("there is an error in running git commit", err)
	}

}


func getCommitMessages(diff string, apiKey string) string{
	prompt:=`You are a git commit message expert. Analyze the following git diff and suggest exactly 3 professional commit messages.

All 3 must follow conventional commit format (feat:, fix:, chore:, refactor:, docs:, etc.)

Message 1 — Short: under 50 chars, bare minimum but clear
Message 2 — Medium: under 72 chars, adds a bit more context
Message 3 — Descriptive: under 72 chars, explains what and why

Examples:
1. feat: add Groq API integration
2. feat: add Groq API to generate commit messages from diff
3. feat: integrate Groq API to auto-suggest commit messages based on git diff

Rules:
- Return ONLY the 3 messages, numbered 1. 2. 3.
- No extra explanation, no headers, no labels, no other text
- Do NOT include style labels like "Short:" or "Medium:"

Git diff:
` + diff + ``
	url:="https://api.groq.com/openai/v1/chat/completions"
	payload:= map[string]interface{}{"model": "llama-3.1-8b-instant", "messages": []map[string]string{{"role": "user", "content": prompt}}}

	jsonValues, err := json.Marshal(payload)

	if err!= nil{
		fmt.Println("error in json payload")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValues))

	req.Header.Set("Authorization", "Bearer "+apiKey);
	req.Header.Set("Content-Type", "application/json");

	client := &http.Client{}
	resp, err := client.Do(req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading response")
	}
	type Message struct {
		Content string `json:"content"`
	}
	type Choice struct {
		Message Message `json:"message"`
	}
	type GroqResponse struct {
		Choices []Choice `json:"choices"`
	}
	var groqResponse GroqResponse
	err = json.Unmarshal(body, &groqResponse)
	if err != nil {
		fmt.Println("error in unmarshaling response")
	}
	return groqResponse.Choices[0].Message.Content
}
