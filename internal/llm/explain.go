package llm

import (
	"appguard/internal/rules"
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

func ExplainFinding(
	ctx context.Context,
	client *genai.Client,
	code string,
	triggerType string,
) (string, error) {
	prompt := fmt.Sprintf(`
		You are a security code reviewer.

		Analyze this code for security risk in less than 20 words.

		Trigger Type: %s

		Code:
		%s
		`, triggerType, code)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return result.Text(), nil

}

func EnrichFindings(
	ctx context.Context,
	client *genai.Client,
	findings []rules.Finding,
) ([]rules.Finding, error) {

	const workers = 4

	jobs := make(chan int, len(findings))
	results := make(chan rules.Finding, len(findings))

	// Start workers
	for w := 0; w < workers; w++ {
		go func() {
			for i := range jobs {

				f := findings[i]

				exp, err := ExplainFinding(
					ctx,
					client,
					f.Code,
					f.Type,
				)

				if err != nil {
					fmt.Printf("LLM error for %s: %v\n", f.Code, err)
					results <- f
					continue
				}

				f.LLMExplanation = exp
				results <- f
			}
		}()
	}

	// Send jobs
	for i := range findings {
		jobs <- i
	}
	close(jobs)

	// Collect results
	enriched := make([]rules.Finding, 0, len(findings))

	for i := 0; i < len(findings); i++ {
		enriched = append(enriched, <-results)
	}

	return enriched, nil
}
