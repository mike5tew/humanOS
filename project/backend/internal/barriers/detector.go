package barriers
package barriers

import (
    "encoding/json"
    "os"
    "regexp"
    "strings"

    "github.com/michaelstewart/humanos/internal/etp"
)

type BarrierDetector struct {
    barriers []etp.StudentBarrier
}

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

func (bd *BarrierDetector) DetectBarriers(input string, context etp.StudentContext) []etp.DetectedBarrier {
    detected := []etp.DetectedBarrier{}

    // Check for "I don't know" - primary avoidance indicator
    if bd.containsIDontKnow(input) {
        if barrier := bd.findBarrierByID("lack_of_motivation"); barrier != nil {
            detected = append(detected, etp.DetectedBarrier{
                Barrier:    *barrier,
                Confidence: 0.8,
                Reasoning:  []string{"Student used 'I don't know' - primary avoidance tactic"},
            })
        }
    }

    // Check for confrontational language
    if bd.isConfrontational(input) {
        if barrier := bd.findBarrierByID("confrontational_showoff"); barrier != nil {
            detected = append(detected, etp.DetectedBarrier{
                Barrier:    *barrier,
                Confidence: 0.7,
                Reasoning:  []string{"Confrontational or dismissive language detected"},
            })
        }
    }

    // Check for minimal response (potential silent avoider)
    if bd.isMinimalResponse(input) {
        if barrier := bd.findBarrierByID("silent_avoider"); barrier != nil {
            detected = append(detected, etp.DetectedBarrier{
                Barrier:    *barrier,
                Confidence: 0.6,
                Reasoning:  []string{"Minimal engagement, very short response"},
            })
        }
    }

    return detected
}

func (bd *BarrierDetector) containsIDontKnow(input string) bool {
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

func (bd *BarrierDetector) isConfrontational(input string) bool {
    patterns := []*regexp.Regexp{
        regexp.MustCompile(`(?i)this is (stupid|dumb|boring)`),
        regexp.MustCompile(`(?i)why (do|should) i`),
        regexp.MustCompile(`(?i)i don'?t (care|want to)`),
        regexp.MustCompile(`(?i)whatever`),
        regexp.MustCompile(`(?i)so what`),
    }
    
    for _, pattern := range patterns {
        if pattern.MatchString(input) {
            return true
        }
    }
    return false
}

func (bd *BarrierDetector) isMinimalResponse(input string) bool {
    trimmed := strings.TrimSpace(input)
    return len(trimmed) < 10 && !bd.containsIDontKnow(input)
}

func (bd *BarrierDetector) findBarrierByID(id string) *etp.StudentBarrier {
    for _, barrier := range bd.barriers {
        if barrier.ID == id {
            return &barrier
        }
    }
    return nil
}
