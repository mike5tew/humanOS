package barriers

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/mike5tew/humanos/internal/etp"
)

// BarrierDetector analyzes student input to identify barriers
type BarrierDetector struct {
	barriers []etp.StudentBarrier
}

// DetectedBarrier represents a barrier with confidence score
type DetectedBarrier struct {
	Barrier    etp.StudentBarrier `json:"barrier"`
	Confidence float64            `json:"confidence"`
	Reasoning  []string           `json:"reasoning"`
}

// NewBarrierDetector creates a barrier detector from JSON schema
func NewBarrierDetector(barriersPath string) (*BarrierDetector, error) {
	data, err := os.ReadFile(barriersPath)
	if err != nil {
		return nil, err
	}

	var barriersData struct {
		Barriers []etp.StudentBarrier `json:"barriers"`
	}

	if err := json.Unmarshal(data, &barriersData); err != nil {
		return nil, err
	}

	return &BarrierDetector{
		barriers: barriersData.Barriers,
	}, nil
}

// DetectBarriers analyzes input for barrier patterns
func (d *BarrierDetector) DetectBarriers(input string, context etp.StudentContext) []DetectedBarrier {
	detected := []DetectedBarrier{}

	// Check for "I don't know" - primary avoidance
	if d.containsIDontKnow(input) {
		// This searches for barrier with id="lack_of_motivation"
		if barrier := d.findBarrierByID("lack_of_motivation"); barrier != nil {
			detected = append(detected, DetectedBarrier{
				Barrier:    *barrier,
				Confidence: 0.8,
				Reasoning:  []string{"Student used 'I don't know' - primary avoidance tactic"},
			})
		}
	}

	// Check for confrontational language
	if d.isConfrontational(input) {
		if barrier := d.findBarrierByID("confrontational_showoff"); barrier != nil {
			detected = append(detected, DetectedBarrier{
				Barrier:    *barrier,
				Confidence: 0.7,
				Reasoning:  []string{"Confrontational or dismissive language detected"},
			})
		}
	}

	// Check for minimal response (silent avoider)
	if d.isMinimalResponse(input) {
		if barrier := d.findBarrierByID("silent_avoider"); barrier != nil {
			detected = append(detected, DetectedBarrier{
				Barrier:    *barrier,
				Confidence: 0.6,
				Reasoning:  []string{"Minimal engagement, very short response"},
			})
		}
	}

	// Check for playful avoidance
	if d.isPlayfulAvoider(input) {
		if barrier := d.findBarrierByID("quiet_playful_avoider"); barrier != nil {
			detected = append(detected, DetectedBarrier{
				Barrier:    *barrier,
				Confidence: 0.65,
				Reasoning:  []string{"Playful or off-topic responses"},
			})
		}
	}

	// Check for high achiever boredom
	if d.indicatesBoredom(input, context) {
		if barrier := d.findBarrierByID("high_achiever_underengaged"); barrier != nil {
			detected = append(detected, DetectedBarrier{
				Barrier:    *barrier,
				Confidence: 0.7,
				Reasoning:  []string{"Indicates boredom or unchallenging material"},
			})
		}
	}

	return detected
}

func (d *BarrierDetector) containsIDontKnow(input string) bool {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)i don'?t know`),
		regexp.MustCompile(`(?i)idk`),
		regexp.MustCompile(`(?i)dunno`),
		regexp.MustCompile(`(?i)no idea`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(input) {
			return true
		}
	}
	return false
}

func (d *BarrierDetector) isConfrontational(input string) bool {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)this is (stupid|dumb|boring)`),
		regexp.MustCompile(`(?i)why (do|should) i`),
		regexp.MustCompile(`(?i)i don'?t (care|want to)`),
		regexp.MustCompile(`(?i)whatever`),
		regexp.MustCompile(`(?i)so what`),
		regexp.MustCompile(`(?i)make me`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(input) {
			return true
		}
	}
	return false
}

func (d *BarrierDetector) isMinimalResponse(input string) bool {
	trimmed := strings.TrimSpace(input)
	return len(trimmed) < 10 && !d.containsIDontKnow(input)
}

func (d *BarrierDetector) isPlayfulAvoider(input string) bool {
	playfulMarkers := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(haha|lol|lmao)`),
		regexp.MustCompile(`(?i)can we (play|do something else)`),
		regexp.MustCompile(`😀|😂|🎮|🎲`),
	}

	for _, marker := range playfulMarkers {
		if marker.MatchString(input) {
			return true
		}
	}
	return false
}

func (d *BarrierDetector) indicatesBoredom(input string, context etp.StudentContext) bool {
	boredomMarkers := []*regexp.Regexp{
		regexp.MustCompile(`(?i)this is (too )?easy`),
		regexp.MustCompile(`(?i)i (already )?know this`),
		regexp.MustCompile(`(?i)when do we do something interesting`),
		regexp.MustCompile(`(?i)can i do something else`),
	}

	for _, marker := range boredomMarkers {
		if marker.MatchString(input) {
			return true
		}
	}
	return false
}

func (d *BarrierDetector) findBarrierByID(id string) *etp.StudentBarrier {
	for i := range d.barriers {
		if d.barriers[i].ID == id { // This looks for the "id" field in JSON
			return &d.barriers[i]
		}
	}
	return nil
}
