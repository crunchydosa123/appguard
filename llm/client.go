package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type LLMRequest struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

func AnalyzeRisk(code string, t string) (string, error) {

	body := LLMRequest{
		Code: code,
		Type: t,
	}

	b, _ := json.Marshal(body)

	resp, err := http.Post(
		"https://api.openai.com/v1/...",
		"application/json",
		bytes.NewBuffer(b),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return "risk result", nil
}
