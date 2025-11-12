package barriers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// AgeGroup represents developmental stage characteristics
type AgeGroup struct {
	Name               string        `json:"name"`
	AgeRange           [2]int        `json:"ageRange"`
	DevelopmentalStage string        `json:"developmentalStage"`
	Characteristics    []string      `json:"characteristics"`
	LanguageGuidelines LanguageGuide `json:"languageGuidelines"`
}

// LanguageGuide contains age-appropriate language rules
type LanguageGuide struct {
	Vocabulary        VocabularyGuide `json:"vocabulary"`
	SentenceStructure SentenceGuide   `json:"sentenceStructure"`
	Concepts          ConceptGuide    `json:"concepts"`
	OffenseRisks      []OffenseRisk   `json:"offenseRisks"`
}

type VocabularyGuide struct {
	Level        string              `json:"level"`
	MaxSyllables int                 `json:"maxSyllables"`
	CanIntroduce []string            `json:"canIntroduce,omitempty"`
	Examples     map[string][]string `json:"examples"`
}

type SentenceGuide struct {
	MaxWordsPerSentence int               `json:"maxWordsPerSentence"`
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
	Prevention string `json:"prevention"`
}

// AgeAppropriateness handles age-based language filtering
type AgeAppropriateness struct {
	ageGroups []AgeGroup
}

// NewAgeAppropriateness creates age filter from JSON schema
func NewAgeAppropriateness(schemaPath string) (*AgeAppropriateness, error) {
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

	return &AgeAppropriateness{
		ageGroups: schema.AgeGroups,
	}, nil
}

// AdjustLanguage modifies response to be age-appropriate
func (aa *AgeAppropriateness) AdjustLanguage(response string, age int) string {
	ageGroup := aa.getAgeGroup(age)
	if ageGroup == nil {
		return response // No adjustment if age group not found
	}

	adjusted := response

	// Simplify vocabulary for younger ages
	if age < 10 {
		adjusted = aa.simplifyVocabulary(adjusted, ageGroup)
	}

	// Shorten sentences if needed
	adjusted = aa.adjustSentenceLength(adjusted, ageGroup)

	// Remove abstract concepts for very young
	if age < 8 {
		adjusted = aa.removeAbstractConcepts(adjusted)
	}

	return adjusted
}

func (aa *AgeAppropriateness) getAgeGroup(age int) *AgeGroup {
	for i := range aa.ageGroups {
		if age >= aa.ageGroups[i].AgeRange[0] && age <= aa.ageGroups[i].AgeRange[1] {
			return &aa.ageGroups[i]
		}
	}
	return nil
}

func (aa *AgeAppropriateness) simplifyVocabulary(text string, ageGroup *AgeGroup) string {
	// Replace complex words with simpler alternatives
	replacements := map[string]string{
		"evaluate":    "look at",
		"consider":    "think about",
		"demonstrate": "show",
		"analyze":     "look at carefully",
		"synthesize":  "put together",
		"hypothesis":  "idea to test",
	}

	result := text
	for complex, simple := range replacements {
		result = strings.ReplaceAll(result, complex, simple)
		result = strings.ReplaceAll(result, strings.Title(complex), strings.Title(simple))
	}

	return result
}

func (aa *AgeAppropriateness) adjustSentenceLength(text string, ageGroup *AgeGroup) string {
	maxWords := ageGroup.LanguageGuidelines.SentenceStructure.MaxWordsPerSentence

	// Split into sentences
	sentences := strings.Split(text, ". ")
	adjusted := []string{}

	for _, sentence := range sentences {
		words := strings.Fields(sentence)
		if len(words) > maxWords {
			// Split long sentence into shorter ones
			mid := maxWords / 2
			firstHalf := strings.Join(words[:mid], " ")
			secondHalf := strings.Join(words[mid:], " ")
			adjusted = append(adjusted, firstHalf+".", secondHalf)
		} else {
			adjusted = append(adjusted, sentence)
		}
	}

	return strings.Join(adjusted, ". ")
}

func (aa *AgeAppropriateness) removeAbstractConcepts(text string) string {
	// Remove or replace abstract phrases
	abstractPhrases := []string{
		"in other words",
		"metaphorically speaking",
		"from a theoretical perspective",
		"conceptually",
	}

	result := text
	for _, phrase := range abstractPhrases {
		result = strings.ReplaceAll(result, phrase, "")
	}

	return strings.TrimSpace(result)
}

// CheckOffenseRisk identifies potential age-inappropriate content
func (aa *AgeAppropriateness) CheckOffenseRisk(response string, age int) []string {
	ageGroup := aa.getAgeGroup(age)
	if ageGroup == nil {
		return []string{}
	}

	risks := []string{}

	// Check for talking down (too simple for age)
	if age >= 10 && strings.Contains(strings.ToLower(response), "good job") {
		risks = append(risks, "Potentially condescending language for age "+fmt.Sprintf("%d", age))
	}

	// Check for complexity overload (too complex for age)
	if age < 8 {
		complexWords := []string{"evaluate", "analyze", "synthesize", "hypothesis"}
		for _, word := range complexWords {
			if strings.Contains(strings.ToLower(response), word) {
				risks = append(risks, "Too complex vocabulary for age "+fmt.Sprintf("%d", age))
				break
			}
		}
	}

	return risks
}

// SafeguardingResponse provides age-appropriate safeguarding message
func (aa *AgeAppropriateness) SafeguardingResponse(age int) string {
	if age < 10 {
		return "Let's talk to a trusted adult about this."
	} else if age < 13 {
		return "I think it would be helpful to talk to someone who can support you better, like a teacher or parent."
	} else {
		return "I think it would be helpful to talk to a trusted adult or counselor about this."
	}
}
