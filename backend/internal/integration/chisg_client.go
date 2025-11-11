package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CHISGClient struct {
	baseURL string
	client  *http.Client
}

func NewCHISGClient() *CHISGClient {
	baseURL := os.Getenv("CHISG_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8082"
	}

	return &CHISGClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// SemanticLink represents a connection in CHISG knowledge graph
type SemanticLink struct {
	FromConcept  string  `json:"from_concept"`
	ToConcept    string  `json:"to_concept"`
	Relationship string  `json:"relationship"`
	Strength     float64 `json:"strength"`
}

// LearningPath represents suggested learning sequence from CHISG
type LearningPath struct {
	Nodes         []LearningNode `json:"nodes"`
	Prerequisites []string       `json:"prerequisites"`
	Difficulty    float64        `json:"difficulty"`
}

type LearningNode struct {
	Concept     string   `json:"concept"`
	Description string   `json:"description"`
	Resources   []string `json:"resources"`
}

// KnowledgeContext represents what CHISG knows about a topic
type KnowledgeContext struct {
	Topic            string         `json:"topic"`
	SemanticLinks    []SemanticLink `json:"semantic_links"`
	SuggestedPath    LearningPath   `json:"suggested_path"`
	PrerequisiteGaps []string       `json:"prerequisite_gaps"`
}

// GetKnowledgeContext queries CHISG for semantic understanding of a topic
func (c *CHISGClient) GetKnowledgeContext(topic string, studentLevel float64) (*KnowledgeContext, error) {
	reqBody := map[string]interface{}{
		"topic":         topic,
		"student_level": studentLevel,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Post(
		c.baseURL+"/api/knowledge/analyze",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CHISG API error: %d", resp.StatusCode)
	}

	var context KnowledgeContext
	if err := json.NewDecoder(resp.Body).Decode(&context); err != nil {
		return nil, err
	}

	return &context, nil
}

// IdentifyPrerequisites asks CHISG what student needs to know first
func (c *CHISGClient) IdentifyPrerequisites(topic string) ([]string, error) {
	reqBody := map[string]string{"topic": topic}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Post(
		c.baseURL+"/api/knowledge/prerequisites",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Prerequisites []string `json:"prerequisites"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Prerequisites, nil
}
