package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/genai"
)

func ExplainRisk(
	ctx context.Context,
	client *genai.Client,
	code string,
	triggerType string,
) (*RiskExplanation, error) {

	prompt := fmt.Sprintf(`
You are a security code reviewer.

Analyze this code for security risk.

Trigger Type: %s

Code:
%s

Return ONLY valid JSON:
{
 "risk_level": "low|medium|high",
 "reason": "...",
 "fix": "...",
 "cwe": "CWE-XXX"
}
`, triggerType, code)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return nil, err
	}

	var explanation RiskExplanation
	err = json.Unmarshal([]byte(result.Text()), &explanation)
	if err != nil {
		return nil, err
	}

	return &explanation, nil
}
