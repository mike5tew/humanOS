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

	// STEP 1: Simplify vocabulary for younger ages
	if age < 10 {
		adjusted = aa.simplifyVocabulary(adjusted, ageGroup)
	}

	// STEP 2: Shorten sentences if needed
	adjusted = aa.adjustSentenceLength(adjusted, ageGroup)

	// STEP 3: Remove abstract concepts for very young
	if age < 8 {
		adjusted = aa.removeAbstractConcepts(adjusted)
	}

	return adjusted
}

// simplifyVocabulary replaces complex words with simpler alternatives
func (aa *AgeAppropriateness) simplifyVocabulary(text string, ageGroup *AgeGroup) string {
	// Replace complex words with simpler alternatives
	replacements := map[string]string{
		"evaluate":    "look at",
		"consider":    "think about",
		"demonstrate": "show",
		"analyze":     "look at carefully",
		"synthesize":  "put together",
		"hypothesis":  "idea to test",
		"implement":   "try out",
		"facilitate":  "help with",
		"comprehend":  "understand",
		"utilize":     "use",
		"commence":    "start",
		"terminate":   "stop",
		"substantial": "big",
		"sufficient":  "enough",
		"challenge":   "hard thing",
		"struggle":    "having trouble",
	}

	result := text
	for complex, simple := range replacements {
		// Case-insensitive replacement
		result = strings.ReplaceAll(result, complex, simple)
		result = strings.ReplaceAll(result, strings.ToUpper(complex[:1])+complex[1:], strings.ToUpper(simple[:1])+simple[1:])
	}

	return result
}

// adjustSentenceLength breaks long sentences into shorter ones
func (aa *AgeAppropriateness) adjustSentenceLength(text string, ageGroup *AgeGroup) string {
	maxWords := ageGroup.LanguageGuidelines.SentenceStructure.MaxWordsPerSentence

	// Split into sentences (simple approach - split on period + space)
	sentences := strings.Split(text, ". ")
	adjusted := []string{}

	for _, sentence := range sentences {
		words := strings.Fields(strings.TrimSpace(sentence))

		if len(words) > maxWords {
			// Split long sentence into shorter ones
			// Split at natural break points (conjunctions, commas)
			chunks := aa.splitSentence(words, maxWords)
			adjusted = append(adjusted, chunks...)
		} else {
			adjusted = append(adjusted, sentence)
		}
	}

	return strings.Join(adjusted, ". ") + "."
}

// splitSentence breaks a word slice into chunks at conjunction points
func (aa *AgeAppropriateness) splitSentence(words []string, maxWords int) []string {
	if len(words) <= maxWords {
		return []string{strings.Join(words, " ")}
	}

	chunks := []string{}
	current := []string{}

	for i, word := range words {
		current = append(current, word)

		// Check if we've hit a conjunction or reached max length
		isConjunction := strings.Contains(word, ",") ||
			strings.ToLower(word) == "and" ||
			strings.ToLower(word) == "but" ||
			strings.ToLower(word) == "or"

		shouldBreak := len(current) >= maxWords || (isConjunction && len(current) >= maxWords/2)

		if shouldBreak && i < len(words)-1 {
			// Remove trailing punctuation from chunk for recombining
			lastWord := current[len(current)-1]
			if strings.HasSuffix(lastWord, ",") {
				current[len(current)-1] = strings.TrimSuffix(lastWord, ",")
			}
			chunks = append(chunks, strings.Join(current, " "))
			current = []string{}
		}
	}

	// Add remaining words
	if len(current) > 0 {
		chunks = append(chunks, strings.Join(current, " "))
	}

	return chunks
}

// removeAbstractConcepts removes or replaces abstract phrases
func (aa *AgeAppropriateness) removeAbstractConcepts(text string) string {
	// Remove abstract phrases that confuse very young children
	abstractPhrases := []string{
		"in other words,",
		"metaphorically speaking,",
		"from a theoretical perspective,",
		"conceptually,",
		"theoretically,",
		"hypothetically,",
		"essentially,",
		"arguably,",
	}

	result := text
	for _, phrase := range abstractPhrases {
		result = strings.ReplaceAll(result, phrase, "")
		result = strings.ReplaceAll(result, strings.ToUpper(phrase[:1])+phrase[1:], "")
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

	// RISK 1: Talking down (overly simple for age)
	if age >= 10 {
		// Check for baby-talk patterns
		babyTalkPatterns := []string{
			"super duper",
			"really really",
			"yay!",
			"good job!",
			"well done!",
		}

		textLower := strings.ToLower(response)
		for _, pattern := range babyTalkPatterns {
			if strings.Contains(textLower, pattern) {
				risks = append(risks, fmt.Sprintf("Potentially condescending language for age %d: '%s'", age, pattern))
				break
			}
		}
	}

	// RISK 2: Too complex vocabulary (overestimating capability)
	if age < 8 {
		complexWords := []string{
			"evaluate", "analyze", "synthesize", "hypothesis",
			"implementation", "facilitate", "comprehend", "utilize",
		}

		textLower := strings.ToLower(response)
		foundComplex := []string{}
		for _, word := range complexWords {
			if strings.Contains(textLower, word) {
				foundComplex = append(foundComplex, word)
			}
		}

		if len(foundComplex) > 0 {
			risks = append(risks, fmt.Sprintf("Too complex vocabulary for age %d: %v", age, foundComplex))
		}
	}

	// RISK 3: Overwhelming complexity (too many clauses/ideas)
	if age < 12 {
		sentences := strings.Split(response, ".")
		for _, sentence := range sentences {
			words := strings.Fields(strings.TrimSpace(sentence))
			if len(words) > ageGroup.LanguageGuidelines.SentenceStructure.MaxWordsPerSentence*2 {
				risks = append(risks, fmt.Sprintf("Sentence too complex for age %d: %d words (max %d)",
					age, len(words), ageGroup.LanguageGuidelines.SentenceStructure.MaxWordsPerSentence))
				break
			}
		}
	}

	// RISK 4: Abstract concepts for very young
	if age < 8 {
		abstractMarkers := []string{
			"consider", "imagine", "suppose", "what if",
			"theoretically", "in theory", "conceptually",
		}

		textLower := strings.ToLower(response)
		for _, marker := range abstractMarkers {
			if strings.Contains(textLower, marker) {
				risks = append(risks, fmt.Sprintf("Abstract concept for age %d: '%s'", age, marker))
				break
			}
		}
	}

	return risks
}

// SafeguardingResponse provides age-appropriate safeguarding message
func (aa *AgeAppropriateness) SafeguardingResponse(age int) string {
	if age < 10 {
		return "Let's talk to a trusted adult about this. A teacher or parent can help."
	} else if age < 13 {
		return "I think it would be helpful to talk to someone who can support you better, like a teacher, parent, or counselor."
	} else {
		return "I think it would be helpful to talk to a trusted adult or counselor about this. Your wellbeing is important."
	}
}
