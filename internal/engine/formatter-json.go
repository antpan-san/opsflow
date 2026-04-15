package engine

import (
	"encoding/json"
	"time"
)

// JSONFormatter JSON 格式化器
type JSONFormatter struct{}

// NewJSONFormatter 创建 JSON 格式化器
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// CheckResult JSON 格式检测结果
type CheckResult struct {
	Name     string                 `json:"name"`
	Success  bool                   `json:"success"`
	Message  string                 `json:"message"`
	Duration string                 `json:"duration,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

// DiagnosticReport JSON 格式诊断报告
type DiagnosticReport struct {
	Scenario    string                 `json:"scenario"`
	Target      string                 `json:"target"`
	Timestamp   string                 `json:"timestamp"`
	Results     map[string]CheckResult `json:"results"`
	Conclusion  string                 `json:"conclusion"`
	Suggestion  string                 `json:"suggestion"`
	Duration    string                 `json:"duration,omitempty"`
}

// FormatReport 格式化诊断报告为 JSON
func (f *JSONFormatter) FormatReport(report DiagnosticReport) (string, error) {
	jsonBytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// CreateReport 创建诊断报告
func (f *JSONFormatter) CreateReport(scenario, target string, results map[string]CheckResult, conclusion, suggestion string, duration time.Duration) DiagnosticReport {
	return DiagnosticReport{
		Scenario:   scenario,
		Target:     target,
		Timestamp:  time.Now().Format(time.RFC3339),
		Results:    results,
		Conclusion: conclusion,
		Suggestion: suggestion,
		Duration:   duration.String(),
	}
}