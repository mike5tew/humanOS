// filepath: /Users/mike5tew/Coding/humanOS/backend/internal/age/filter.go
package age

import (
	"encoding/json"
	"os"
	"strings"
)

type AgeGroup struct {
	Name               string        `json:"name"`
	AgeRange           [2]int        `json:"ageRange"`
	DevelopmentalStage string        `json:"developmentalStage"`
	LanguageGuidelines LanguageGuide `json:"languageGuidelines"`
	OffenseRisks       []OffenseRisk `json:"offenseRisks"`
}

type LanguageGuide struct {
	Vocabulary        VocabularyGuide `json:"vocabulary"`
	SentenceStructure SentenceGuide   `json:"sentenceStructure"`
	Concepts          ConceptGuide    `json:"concepts"`
}

type VocabularyGuide struct {
	Level        string              `json:"level"`
	MaxSyllables int                 `json:"maxSyllables,omitempty"`
	Examples     map[string][]string `json:"examples"`
}

type SentenceGuide struct {
	MaxWordsPerSentence interface{}       `json:"maxWordsPerSentence"` // can be int or "no_limit"
	Structure           string            `json:"structure"`
	Examples            map[string]string `json:"examples"`
}

type ConceptGuide struct {
	Allowed  string            `json:"allowed"`
	Examples map[string]string `json:"examples"`
}

type OffenseRisk struct {
	Risk       string `json:"risk"`
	Trigger    string `json:"trigger"`
	Severity   string `json:"severity,omitempty"`
	Prevention string `json:"prevention"`
}

type AgeFilter struct {
	ageGroups []AgeGroup
}

func NewAgeFilter(schemaPath string) (*AgeFilter, error) {
	data, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	var schema struct {
		AgeGroups []AgeGroup `json:"ageGroups"`
	}

	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, err
	}

	return &AgeFilter{ageGroups: schema.AgeGroups}, nil
}

// AdjustForAge takes a response and student age, returns age-appropriate version
func (af *AgeFilter) AdjustForAge(response string, age int) (string, error) {
	group := af.getAgeGroup(age)
	if group == nil {
		return response, nil // Default: no adjustment if age group not found
	}

	adjusted := response

	// Step 1: Simplify vocabulary if needed
	adjusted = af.simplifyVocabulary(adjusted, group)

	// Step 2: Shorten sentences if needed
	adjusted = af.adjustSentenceLength(adjusted, group)

	// Step 3: Check for offense risks
	if af.detectOffenseRisk(adjusted, group) {
		// Log warning and provide safer alternative
		adjusted = af.generateSaferAlternative(adjusted, group)
	}

	return adjusted, nil
}

func (af *AgeFilter) getAgeGroup(age int) *AgeGroup {
	for _, group := range af.ageGroups {
		if age >= group.AgeRange[0] && age <= group.AgeRange[1] {
			return &group
		}
	}
	return nil
}

func (af *AgeFilter) simplifyVocabulary(text string, group *AgeGroup) string {
	// For early primary (5-7), replace complex words with simple alternatives
	if group.LanguageGuidelines.Vocabulary.Level == "simple_everyday_words_only" {
		replacements := map[string]string{
			"consider":    "think about",
			"evaluate":    "look at",
			"demonstrate": "show",
			"attempt":     "try",
			"capability":  "what you can do",
		}

		for complex, simple := range replacements {
			text = strings.ReplaceAll(text, complex, simple)
		}
	}

	return text
}

func (af *AgeFilter) adjustSentenceLength(text string, group *AgeGroup) string {
	// Implementation would split long sentences, simplify structure
	// This is a placeholder - full implementation would use NLP

	maxWords := af.getMaxWordsPerSentence(group)
	if maxWords == 0 {
		return text // No limit
	}

	// TODO: Implement sentence splitting logic
	return text
}

func (af *AgeFilter) getMaxWordsPerSentence(group *AgeGroup) int {
	switch v := group.LanguageGuidelines.SentenceStructure.MaxWordsPerSentence.(type) {
	case float64:
		return int(v)
	case string:
		if v == "no_limit" {
			return 0
		}
	}
	return 15 // default
}

func (af *AgeFilter) detectOffenseRisk(text string, group *AgeGroup) bool {
	// Check for common offense triggers
	lowerText := strings.ToLower(text)

	for _, risk := range group.OffenseRisks {
		// Check for condescension markers
		if strings.Contains(risk.Risk, "condescension") || strings.Contains(risk.Risk, "Talking down") {
			condescensionMarkers := []string{
				"good job!",
				"well done!",
				"you're such a big kid",
				"can you do this for me?",
			}

			for _, marker := range condescensionMarkers {
				if strings.Contains(lowerText, marker) {
					return true
				}
			}
		}
	}

	return false
}

func (af *AgeFilter) generateSaferAlternative(text string, group *AgeGroup) string {
	// In a full implementation, this would use the barrier intervention
	// to regenerate the response with stricter age guidelines

	// For now, return original with warning logged
	return text
}

// CheckOffenseRisk returns specific offense risks for this age + response
func (af *AgeFilter) CheckOffenseRisk(response string, age int) []OffenseRisk {
	group := af.getAgeGroup(age)
	if group == nil {
		return []OffenseRisk{}
	}

	risks := []OffenseRisk{}

	for _, risk := range group.OffenseRisks {
		if af.responseMatchesRisk(response, risk) {
			risks = append(risks, risk)
		}
	}

	return risks
}

func (af *AgeFilter) responseMatchesRisk(response string, risk OffenseRisk) bool {
	// Implement risk detection logic based on trigger patterns
	lowerResponse := strings.ToLower(response)
	lowerTrigger := strings.ToLower(risk.Trigger)

	// Simple substring match - full implementation would be more sophisticated
	return strings.Contains(lowerResponse, lowerTrigger)
}
