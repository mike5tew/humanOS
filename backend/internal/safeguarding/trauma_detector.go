package safeguarding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
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
	Timestamp time.Time
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

// SafeguardingLog stores incident for human review
type SafeguardingLog struct {
	Timestamp   time.Time
	StudentID   string
	Age         int
	Message     string
	Severity    int
	Category    string
	Response    string
	HumanReview bool
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
					Timestamp: time.Now(),
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
		Detected:  false,
		Severity:  0,
		Timestamp: time.Now(),
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

	// No adjustment for older ages - they have more context
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

// GetEscalationMessage returns appropriate response based on severity
func (td *TraumaDetector) GetEscalationMessage(severity int, age int) string {
	if severity >= 4 {
		if age < 10 {
			return "I'm very concerned about what you've shared. Let's get a trusted adult right now - a teacher, parent, or counselor. Your safety is the most important thing."
		} else if age < 13 {
			return "I'm very concerned about what you've shared. Your safety is the most important thing right now. I'm going to make sure you talk with someone who can help immediately - like a school counselor, trusted teacher, or parent."
		} else {
			return "I'm very concerned about what you've shared. Your safety is the most important thing. I'm connecting you with someone trained to help - please reach out to a trusted adult, counselor, or contact a crisis helpline immediately."
		}
	}

	if severity >= 3 {
		if age < 10 {
			return "I think it would be really helpful for you to talk with a trusted adult about this - like a teacher or parent. They can help you better than I can."
		} else if age < 13 {
			return "I think it would be really helpful for you to talk with someone who can support you better, like a school counselor, teacher, or parent. I'm going to let them know you need to talk."
		} else {
			return "I think it would be helpful to talk with a trusted adult or counselor about this. Your wellbeing is important. Please reach out to a school counselor, trusted teacher, parent, or contact a crisis helpline if you need immediate help."
		}
	}

	return "Let's take a break from the work for now. I want to make sure you get some support."
}

// LogIncident creates audit trail for safeguarding concern
func (td *TraumaDetector) LogIncident(studentID string, age int, message string, result TraumaResult, aiResponse string) SafeguardingLog {
	return SafeguardingLog{
		Timestamp:   result.Timestamp,
		StudentID:   studentID,
		Age:         age,
		Message:     message,
		Severity:    result.Severity,
		Category:    result.Category,
		Response:    aiResponse,
		HumanReview: result.Severity >= 3,
	}
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
	return time.Now().Unix()
}
