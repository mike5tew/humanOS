package safeguarding

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

// TraumaResult contains detection results and severity
type TraumaResult struct {
	Severity int      `json:"severity"` // 0-4
	Patterns []string `json:"patterns"`
	Action   string   `json:"action"`
}

// TraumaDetector identifies concerning content in student messages
type TraumaDetector struct {
	patterns map[string][]*regexp.Regexp
}

// NewTraumaDetector creates detector from patterns file
func NewTraumaDetector(patternsPath string) (*TraumaDetector, error) {
	data, err := os.ReadFile(patternsPath)
	if err != nil {
		return nil, err
	}

	var patternsData map[string][]string
	if err := json.Unmarshal(data, &patternsData); err != nil {
		return nil, err
	}

	td := &TraumaDetector{
		patterns: make(map[string][]*regexp.Regexp),
	}

	// Compile regex patterns
	for category, patterns := range patternsData {
		compiled := make([]*regexp.Regexp, 0, len(patterns))
		for _, pattern := range patterns {
			if re, err := regexp.Compile(pattern); err == nil {
				compiled = append(compiled, re)
			}
		}
		td.patterns[category] = compiled
	}

	return td, nil
}

// Scan analyzes message for trauma/safeguarding indicators
func (td *TraumaDetector) Scan(message string, age int) (TraumaResult, error) {
	lowerMsg := strings.ToLower(message)
	severity := 0
	detectedPatterns := []string{}

	// Check each category
	for category, patterns := range td.patterns {
		for _, pattern := range patterns {
			if pattern.MatchString(lowerMsg) {
				detectedPatterns = append(detectedPatterns, category)
				severity = max(severity, td.calculateSeverity(category, age))
				break
			}
		}
	}

	action := td.determineAction(severity)

	return TraumaResult{
		Severity: severity,
		Patterns: detectedPatterns,
		Action:   action,
	}, nil
}

func (td *TraumaDetector) calculateSeverity(category string, age int) int {
	// Severity scale:
	// 1-2: Log and monitor
	// 3: Human review within 24h
	// 4: Immediate emergency services

	baseSeverity := map[string]int{
		"mild_concern":     1,
		"moderate_concern": 2,
		"serious_concern":  3,
		"emergency":        4,
		// Specific categories
		"sexual_content": 4,
		"violence":       3,
		"neglect":        3,
		"self_harm":      4,
		"abuse":          4,
	}

	return baseSeverity[category]
}

func (td *TraumaDetector) determineAction(severity int) string {
	switch severity {
	case 1, 2:
		return "log_and_monitor"
	case 3:
		return "human_review_24h"
	case 4:
		return "immediate_escalation"
	default:
		return "continue"
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
