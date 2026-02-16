package scanner

type Finding struct {
	File string
	Line int
	Code string
	Type string

	LLMRisk   string
	LLMReason string
	LLMFix    string
}
