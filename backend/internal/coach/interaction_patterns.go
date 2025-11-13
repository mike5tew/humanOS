package coach

import (
	"fmt"

	"github.com/mike5tew/humanos/internal/etp"
)

// InteractionPattern defines proven engagement sequences
type InteractionPattern struct {
	Name        string
	Purpose     string
	Sequence    []InteractionStep
	VoltageGoal string // "lower", "maintain", "raise_gradually"
}

// InteractionStep represents one stage in interaction sequence
type InteractionStep struct {
	Action        string
	Expected      string
	IfSuccess     string
	IfStruggle    string
	VoltageImpact float64 // negative = lowers, positive = raises
}

// SessionOpener handles voltage reduction through familiarity
type SessionOpener struct {
	WarmGreeting       string
	PreviousRecap      string
	InterestConnection string
	MicroSuccessSetup  string
}

// QuestionProgression manages difficulty scaling
type QuestionProgression struct {
	StartDifficulty   float64 // 0-1, deliberately start low
	CurrentStreak     int
	LastAnswerQuality string
	SemanticDistance  int // Steps between question and answer
}

// NextQuestion selects appropriate difficulty level
func (qp *QuestionProgression) NextQuestion(context etp.StudentContext) QuestionSpec {
	spec := QuestionSpec{}

	// If struggling: guide toward guaranteed win
	if qp.LastAnswerQuality == "struggling" {
		spec.Difficulty = 0.2     // Impossibly easy
		spec.SemanticDistance = 1 // Direct recall
		spec.Hint = "Think about what we just covered..."
		return spec
	}

	// If streak of successes: gradually increase
	if qp.CurrentStreak >= 3 {
		spec.Difficulty = minFloat(qp.StartDifficulty+0.1, 0.9)
		// FIX: Cast to int properly
		spec.SemanticDistance = minInt(qp.SemanticDistance+1, 5)
	} else {
		spec.Difficulty = qp.StartDifficulty
		spec.SemanticDistance = qp.SemanticDistance
	}

	return spec
}

// QuestionSpec defines question characteristics
type QuestionSpec struct {
	Difficulty       float64 // 0-1
	SemanticDistance int     // 1-5: jumps between Q and A
	Extrapolation    float64 // 0-1: recall vs novel application
	Hint             string
}

// LearningJourneyStage represents progression through content
type LearningJourneyStage int

const (
	StageRecap LearningJourneyStage = iota
	StageSemanticLinks
	StageKeywordGames
	StageApplication
	StageExtrapolation
)

// LearningPathway defines different learning approaches
type LearningPathway string

const (
	PathwayProjectBased     LearningPathway = "project_based"
	PathwaySensoryRegulated LearningPathway = "sensory_regulated"
	PathwayTraditional      LearningPathway = "traditional"
)

// SensoryProfile configures interface for neurodiversity
type SensoryProfile struct {
	NoisePreference  string // "minimal", "moderate", "high"
	VisualComplexity string // "simple", "moderate", "rich"
	InteractionPace  string // "slow", "medium", "fast"
}

// ConceptMapState tracks student's concept understanding
type ConceptMapState struct {
	ConceptID  string
	Confidence float64 // 0-1: down/struggling to up/mastered
	Focus      float64 // 0-1: left/disconnected to right/progressing
}

// CalculateVoltage derives stress level from state
func (cms *ConceptMapState) CalculateVoltage() float64 {
	// High voltage = low confidence + low focus
	return (1.0 - cms.Confidence) * (1.0 - cms.Focus)
}

// InterventionNeeded checks if AI should step in
func (cms *ConceptMapState) InterventionNeeded() bool {
	return cms.Confidence < 0.4 && cms.Focus < 0.4
}

// VirtualTeam creates solo competition mechanics
type VirtualTeam struct {
	CurrentPerformance float64
	PastPerformance    []float64 // History
	TargetPerformance  float64   // Goal
	ImprovementRate    float64   // Trend
}

// CompeteAgainstSelf generates motivational comparison
func (vt *VirtualTeam) CompeteAgainstSelf() string {
	if len(vt.PastPerformance) == 0 {
		return "Let's set your first baseline!"
	}

	lastSession := vt.PastPerformance[len(vt.PastPerformance)-1]

	if vt.CurrentPerformance > lastSession {
		improvement := ((vt.CurrentPerformance - lastSession) / lastSession) * 100
		return fmt.Sprintf("You're %.1f%% better than last session! 🚀", improvement)
	}

	return "Let's match your personal best today!"
}

// GetInteractionPattern returns proven engagement sequence
func GetInteractionPattern(name string) *InteractionPattern {
	patterns := map[string]InteractionPattern{
		"voltage_reduction": {
			Name:        "voltage_reduction",
			Purpose:     "Lower emotional resistance before content",
			VoltageGoal: "lower",
			Sequence: []InteractionStep{
				{
					Action:        "warm_greeting",
					Expected:      "student_acknowledges",
					VoltageImpact: -0.2,
				},
				{
					Action:        "previous_recap",
					Expected:      "student_remembers",
					IfSuccess:     "connect_to_interest",
					IfStruggle:    "provide_gentle_reminder",
					VoltageImpact: -0.1,
				},
				{
					Action:        "micro_success_setup",
					Expected:      "student_succeeds",
					IfSuccess:     "pile_on_praise",
					VoltageImpact: -0.3,
				},
			},
		},

		"semantic_distance_progression": {
			Name:        "semantic_distance_progression",
			Purpose:     "Gradually extend thinking depth",
			VoltageGoal: "raise_gradually",
			Sequence: []InteractionStep{
				{
					Action:        "direct_recall",
					Expected:      "quick_answer",
					IfSuccess:     "one_inference_question",
					VoltageImpact: 0.1,
				},
				{
					Action:        "one_inference_question",
					Expected:      "thinks_briefly",
					IfSuccess:     "two_inference_question",
					IfStruggle:    "return_to_direct_recall",
					VoltageImpact: 0.2,
				},
				{
					Action:        "novel_application",
					Expected:      "genuine_thinking",
					IfSuccess:     "celebrate_breakthrough",
					IfStruggle:    "provide_scaffolding",
					VoltageImpact: 0.3,
				},
			},
		},
	}

	if pattern, ok := patterns[name]; ok {
		return &pattern
	}
	return nil
}

// minFloat helper for float64
func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// minInt helper for int
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
