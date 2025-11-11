// filepath: /Users/michaelstewart/Coding/humanOS/backend/internal/coach/orchestrator.go
package coach

import (
	"github.com/michaelstewart/humanos/internal/age"
	"github.com/michaelstewart/humanos/internal/barriers"
	"github.com/michaelstewart/humanos/internal/etp"
	"github.com/michaelstewart/humanos/internal/safeguarding"
)

type CoachOrchestrator struct {
	barrierDetector *barriers.BarrierDetector
	traumaDetector  *safeguarding.TraumaDetector
	ageFilter       *age.AgeFilter // NEW
}

func NewCoachOrchestrator(barriersPath, traumaPath, agePath string) (*CoachOrchestrator, error) {
	bd, err := barriers.NewBarrierDetector(barriersPath)
	if err != nil {
		return nil, err
	}

	td, err := safeguarding.NewTraumaDetector(traumaPath)
	if err != nil {
		return nil, err
	}

	af, err := age.NewAgeFilter(agePath) // NEW
	if err != nil {
		return nil, err
	}

	return &CoachOrchestrator{
		barrierDetector: bd,
		traumaDetector:  td,
		ageFilter:       af, // NEW
	}, nil
}

func (co *CoachOrchestrator) ProcessStudentMessage(
	studentID string,
	message string,
	context etp.StudentContext,
) (etp.CoachResponse, error) {
	reasoning := []string{}

	// ...existing code for trauma detection and barrier detection...

	// Generate response (existing code)
	responseMessage := co.generateResponse(intervention, context, detectedBarriers)

	// NEW: Age-appropriate adjustment BEFORE sending
	adjustedMessage, err := co.ageFilter.AdjustForAge(responseMessage, context.Age)
	if err != nil {
		reasoning = append(reasoning, "Age adjustment failed, using original")
		adjustedMessage = responseMessage
	} else {
		reasoning = append(reasoning, "Response adjusted for age "+string(context.Age))
	}

	// NEW: Check for offense risks
	offenseRisks := co.ageFilter.CheckOffenseRisk(adjustedMessage, context.Age)
	if len(offenseRisks) > 0 {
		reasoning = append(reasoning, "WARNING: Potential offense risk detected")
		// In production: regenerate response or escalate to human review
	}

	return etp.CoachResponse{
		Message:           adjustedMessage, // Use adjusted, not original
		Intervention:      intervention,
		DetectedBarriers:  barriers,
		SafeguardingAlert: traumaResult.Severity > 0,
		RewardEarned:      rewardEarned,
		Reasoning:         reasoning,
	}, nil
}
