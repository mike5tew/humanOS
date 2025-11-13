package coach

import (
	"regexp"
	"sync"
	"time"
)

// Interest represents a detected student interest
type Interest struct {
	Category      string
	Specific      string
	Confidence    float64
	LastMentioned time.Time
	MentionCount  int
}

// PersonalizationEngine tracks student interests (in-memory for MVP)
type PersonalizationEngine struct {
	studentInterests map[string][]Interest
	mu               sync.RWMutex
}

// NewPersonalizationEngine creates personalization engine
func NewPersonalizationEngine() *PersonalizationEngine {
	return &PersonalizationEngine{
		studentInterests: make(map[string][]Interest),
	}
}

// DetectInterests extracts interests from message
func (pe *PersonalizationEngine) DetectInterests(message string) []Interest {
	detected := []Interest{}

	patterns := map[string][]struct {
		pattern *regexp.Regexp
		name    string
	}{
		"games": {
			{regexp.MustCompile(`(?i)minecraft`), "Minecraft"},
			{regexp.MustCompile(`(?i)fortnite`), "Fortnite"},
			{regexp.MustCompile(`(?i)roblox`), "Roblox"},
		},
		"sports": {
			{regexp.MustCompile(`(?i)football|soccer`), "Football"},
			{regexp.MustCompile(`(?i)basketball`), "Basketball"},
		},
		"hobbies": {
			{regexp.MustCompile(`(?i)drawing|art`), "Drawing"},
			{regexp.MustCompile(`(?i)music`), "Music"},
		},
	}

	for category, categoryPatterns := range patterns {
		for _, p := range categoryPatterns {
			if p.pattern.MatchString(message) {
				detected = append(detected, Interest{
					Category:      category,
					Specific:      p.name,
					Confidence:    0.8,
					LastMentioned: time.Now(),
					MentionCount:  1,
				})
			}
		}
	}

	return detected
}

// TrackInterests stores interests in memory
func (pe *PersonalizationEngine) TrackInterests(studentID string, interests []Interest) {
	pe.mu.Lock()
	defer pe.mu.Unlock()

	existing := pe.studentInterests[studentID]

	for _, newInterest := range interests {
		found := false
		for i, existingInterest := range existing {
			if existingInterest.Category == newInterest.Category &&
				existingInterest.Specific == newInterest.Specific {
				existing[i].MentionCount++
				existing[i].LastMentioned = time.Now()
				existing[i].Confidence = min(existing[i].Confidence+0.1, 1.0)
				found = true
				break
			}
		}
		if !found {
			existing = append(existing, newInterest)
		}
	}

	pe.studentInterests[studentID] = existing
}

// GetStudentInterests retrieves tracked interests
func (pe *PersonalizationEngine) GetStudentInterests(studentID string) []Interest {
	pe.mu.RLock()
	defer pe.mu.RUnlock()
	return pe.studentInterests[studentID]
}

// min helper
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
