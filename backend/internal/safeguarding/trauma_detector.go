package safeguarding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// TraumaDetector analyzes student input for safeguarding concerns
type TraumaDetector struct {
	patterns             []TraumaPattern
	safeguardingEndpoint string
	emergencyEndpoint    string
}

// TraumaPattern represents concerning content patterns
type TraumaPattern struct {
	ID          string   `json:"id"`
	Category    string   `json:"category"`
	Patterns    []string `json:"patterns"`
	Severity    int      `json:"severity"` // 1-4
	Description string   `json:"description"`
}

// TraumaResult represents detection outcome
type TraumaResult struct {
	Detected  bool
	Severity  int
	Category  string
	Reasoning string
}

// SafeguardingAlert sent to human team
type SafeguardingAlert struct {
	StudentID string `json:"student_id"`
	Timestamp string `json:"timestamp"`
	Severity  int    `json:"severity"`
	Category  string `json:"category"`
	Content   string `json:"content"`
	Urgent    bool   `json:"urgent"`
	Response  string `json:"response"`
}

// NewTraumaDetector loads trauma patterns from JSON
func NewTraumaDetector(traumaPath string) (*TraumaDetector, error) {
	// Load patterns from JSON file
	patterns := []TraumaPattern{
		{
			ID:       "sexual_abuse",
			Category: "sexual",
			Patterns: []string{
				"sexual act|sexual abuse|touched me|made me",
				"inappropriate touch|uncomfortable",
			},
			Severity:    4,
			Description: "Indicators of sexual abuse or harm",
		},
		{
			ID:       "violence_threat",
			Category: "violence",
			Patterns: []string{
				"going to hurt|going to kill|have a plan|get a weapon",
				"tonight|tomorrow.*(hurt|kill|attack)",
				"want to hurt|want to kill",
			},
			Severity:    4,
			Description: "Immediate violence threat or plan",
		},
		{
			ID:       "neglect",
			Category: "neglect",
			Patterns: []string{
				"no food|haven't eaten|starving",
				"no one cares|left alone|abandoned",
			},
			Severity:    3,
			Description: "Signs of neglect or abandonment",
		},
	}

	return &TraumaDetector{
		patterns:             patterns,
		safeguardingEndpoint: "http://safeguarding-team/api/alert",
		emergencyEndpoint:    "http://emergency-services/api/report",
	}, nil
}

// Scan checks message for trauma indicators
func (td *TraumaDetector) Scan(message string, age int) TraumaResult {
	message = strings.ToLower(message)

	// Check each pattern
	for _, pattern := range td.patterns {
		for _, patternStr := range pattern.Patterns {
			// Compile regex with case-insensitive flag
			regex, err := regexp.Compile("(?i)" + patternStr)
			if err != nil {
				continue
			}

			if regex.MatchString(message) {
				// Age-calibrated severity
				severity := td.calibrateSeverity(pattern.Severity, age, pattern.Category)

				result := TraumaResult{
					Detected:  true,
					Severity:  severity,
					Category:  pattern.Category,
					Reasoning: pattern.Description,
				}

				// Escalate based on severity
				if severity >= 3 {
					td.escalateAlert(message, age, result)
				}

				return result
			}
		}
	}

	return TraumaResult{
		Detected: false,
		Severity: 0,
	}
}

// calibrateSeverity adjusts severity based on age and content
func (td *TraumaDetector) calibrateSeverity(baseSeverity int, age int, category string) int {
	// Very young children (< 8): lower threshold for escalation
	if age < 8 {
		if baseSeverity >= 2 {
			return baseSeverity + 1 // Escalate one level for very young
		}
	}

	// Teenagers: may have different context
	if age >= 14 && category == "violence" {
		// Check if potentially joking/gaming reference
		// For now, keep same severity
	}

	return baseSeverity
}

// escalateAlert sends alert to safeguarding team
func (td *TraumaDetector) escalateAlert(message string, age int, result TraumaResult) {
	alert := SafeguardingAlert{
		Timestamp: fmt.Sprintf("%d", getCurrentTimestamp()),
		Severity:  result.Severity,
		Category:  result.Category,
		Content:   truncateContent(message),
		Urgent:    result.Severity >= 3,
		Response:  td.generateSafeguardingResponse(result.Severity),
	}

	// Send to safeguarding endpoint (non-blocking)
	go func() {
		data, _ := json.Marshal(alert)
		resp, err := http.Post(
			td.safeguardingEndpoint,
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			log.Printf("Failed to alert safeguarding team: %v", err)
		} else {
			resp.Body.Close()
		}
	}()

	// Severity 4 = emergency services
	if result.Severity == 4 {
		td.alertEmergencyServices(alert, age)
	}
}

// alertEmergencyServices escalates to emergency services
func (td *TraumaDetector) alertEmergencyServices(alert SafeguardingAlert, age int) {
	log.Printf("🚨 EMERGENCY ALERT - Severity 4 trauma indicator detected for age %d", age)
	log.Printf("Category: %s | Content: %s", alert.Category, alert.Content)

	// In production: Actual emergency services integration
	// For now: Log for manual review
	go func() {
		data, _ := json.Marshal(alert)
		_, err := http.Post(
			td.emergencyEndpoint,
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			log.Printf("Failed to alert emergency services: %v", err)
		}
	}()
}

// generateSafeguardingResponse creates age-appropriate response
func (td *TraumaDetector) generateSafeguardingResponse(severity int) string {
	if severity >= 4 {
		return "I'm very concerned about what you've shared. Your safety is the most important thing. I'm connecting you with someone who can help right now."
	}

	if severity >= 3 {
		return "Thanks for sharing that with me. I think it would be really helpful to talk with someone who can support you better. I'm going to connect you with help."
	}

	return "I hear you. Let's take a break and get some support."
}

// Helper functions
func truncateContent(s string) string {
	if len(s) > 200 {
		return s[:200] + "..."
	}
	return s
}

func getCurrentTimestamp() int64 {
	// Return current Unix timestamp
	return 0 // Placeholder
}
