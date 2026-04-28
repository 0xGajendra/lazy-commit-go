package groq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Message struct {
	Content string `json:"content"`
}
type Choice struct {
	Message Message `json:"message"`
}
type GroqResponse struct {
	Choices []Choice `json:"choices"`
}
const commitPrompt=`You are a git commit message expert. Analyze the following git diff and suggest exactly 3 professional commit messages.

All 3 must follow conventional commit format (feat:, fix:, chore:, refactor:, docs:, etc.)

Message 1 — Short: under 50 chars, bare minimum but clear
Message 2 — Medium: under 72 chars, adds a bit more context
Message 3 — Descriptive: under 72 chars, explains what and why

Examples:
feat: add Groq API integration
feat: add Groq API to generate commit messages from diff
feat: integrate Groq API to auto-suggest commit messages based on git diff

Rules:
- Return ONLY the 3 messages, numbered 1. 2. 3.
- No extra explanation, no headers, no labels, no other text
- Do NOT include style labels like "Short:" or "Medium:"
- Do NOT include the number prefix (no "1." "2." "3.")
- Maximum 60 characters per message, be concise

Git diff:
`

func GetCommitMessages(diff, apiKey string) ([]string, error) {
		
	url:="https://api.groq.com/openai/v1/chat/completions"
	payload:= map[string]interface{}{"model": "llama-3.3-70b-versatile", "messages": []map[string]string{{"role": "user", "content": commitPrompt + diff}}}

	jsonValues, err := json.Marshal(payload)

	if err!= nil{
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValues))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey);
	req.Header.Set("Content-Type", "application/json");

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var groqResponse GroqResponse
	err = json.Unmarshal(body, &groqResponse)
	if err != nil {
		return nil, err
	}
	commitMessages:=strings.Split(groqResponse.Choices[0].Message.Content, "\n")

	
	return commitMessages, nil
}

