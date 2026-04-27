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

		result:=getCommitMessages(string(out), groqAPI)
		fmt.Println("result", result)

	}

}


func getCommitMessages(diff string, apiKey string) string{
	prompt:=`You are a git commit message expert. Analyze the following git diff and suggest exactly 3 commit messages, each with a different style.

Message 1 — Concise: short and to the point, under 50 chars, conventional commit format (feat:, fix:, chore: etc.)
Message 2 — Detailed: descriptive, explains what and why, under 72 chars, conventional commit format
Message 3 — Casual: informal, human tone, no strict format, something a developer would actually type at 2am

Rules:
- Return ONLY the 3 messages, numbered 1. 2. 3.
- No extra explanation, no headers, no other text whatsoever

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
