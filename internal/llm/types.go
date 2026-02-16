package llm

type RiskExplanation struct {
	RiskLevel string `json:"risk_level"`
	Reason    string `json:"reason"`
	Fix       string `json:"fix"`
	CWE       string `json:"cwe"`
}
