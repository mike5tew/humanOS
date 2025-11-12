// filepath: /Users/mike5tew/Coding/humanOS/backend/internal/coach/orchestrator.go
package coach

import (
	"strings"
	"time"

	"github.com/mike5tew/humanos/internal/barriers"
	"github.com/mike5tew/humanos/internal/etp"
	"github.com/mike5tew/humanos/internal/safeguarding"
)

// Orchestrator combines all systems for complete coaching response
type Orchestrator struct {
	barrierDetector *barriers.BarrierDetector
	traumaDetector  *safeguarding.TraumaDetector
	ageFilter       *barriers.AgeAppropriateness
}

// NewOrchestrator creates complete coaching system
func NewOrchestrator(
	barriersPath string,
	traumaPath string,
	agePath string,
) (*Orchestrator, error) {
	bd, err := barriers.NewBarrierDetector(barriersPath)
	if err != nil {
		return nil, err
	}

	td, err := safeguarding.NewTraumaDetector(traumaPath)
	if err != nil {
		return nil, err
	}

	af, err := barriers.NewAgeAppropriateness(agePath)
	if err != nil {
		return nil, err
	}

	return &Orchestrator{
		barrierDetector: bd,
		traumaDetector:  td,
		ageFilter:       af,
	}, nil
}

// ProcessMessage handles complete student interaction
func (o *Orchestrator) ProcessMessage(
	studentID string,
	message string,
	context etp.StudentContext,
) (etp.CoachResponse, error) {
	reasoning := []string{}

	// 1. CRITICAL: Check for trauma/safeguarding issues first
	traumaResult, err := o.traumaDetector.Scan(message, context.Age)
	if err != nil {
		return etp.CoachResponse{}, err
	}

	if traumaResult.Severity >= 3 {
		reasoning = append(reasoning, "Safeguarding concern detected - escalating to human team")
		response := o.ageFilter.SafeguardingResponse(context.Age)

		return etp.CoachResponse{
			Message:           response,
			Intervention:      nil,
			DetectedBarriers:  []etp.StudentBarrier{},
			SafeguardingAlert: true,
			RewardEarned:      false,
			Reasoning:         reasoning,
			Timestamp:         time.Now(),
		}, nil
	}

	// 2. Detect barriers
	detectedBarriers := o.barrierDetector.DetectBarriers(message, context)
	if len(detectedBarriers) > 0 {
		reasoning = append(reasoning, "Detected barrier: "+detectedBarriers[0].Barrier.Name)
	}

	// 3. Select intervention
	intervention := o.selectIntervention(detectedBarriers, context)
	if intervention != nil {
		reasoning = append(reasoning, "Selected intervention: "+intervention.Name)
	}

	// 4. Generate base response
	responseMessage := o.generateResponse(intervention, context, detectedBarriers)

	// 5. CRITICAL: Adjust for age appropriateness
	responseMessage = o.ageFilter.AdjustLanguage(responseMessage, context.Age)

	// 6. Check for offense risks
	offenseRisks := o.ageFilter.CheckOffenseRisk(responseMessage, context.Age)
	if len(offenseRisks) > 0 {
		reasoning = append(reasoning, "Warning: "+strings.Join(offenseRisks, "; "))
	}

	// 7. Check if reward earned
	rewardEarned := o.checkRewardEarned(message, context, intervention)
	if rewardEarned {
		reasoning = append(reasoning, "Student earned play break reward")
	}

	// Extract barriers for response
	barrierList := make([]etp.StudentBarrier, len(detectedBarriers))
	for i, db := range detectedBarriers {
		barrierList[i] = db.Barrier
	}

	return etp.CoachResponse{
		Message:           responseMessage,
		Intervention:      intervention,
		DetectedBarriers:  barrierList,
		SafeguardingAlert: traumaResult.Severity > 0,
		RewardEarned:      rewardEarned,
		Reasoning:         reasoning,
		Timestamp:         time.Now(),
	}, nil
}

func (o *Orchestrator) selectIntervention(
	detectedBarriers []barriers.DetectedBarrier,
	context etp.StudentContext,
) *etp.InterventionLever {
	if len(detectedBarriers) == 0 {
		return nil
	}

	topBarrier := detectedBarriers[0].Barrier

	// Check brain state - if emotional voltage too high, prioritize calming
	if context.BrainState.EmotionalLevel > 0.7 {
		// Find calming intervention
		for i := range topBarrier.EffectiveLevers {
			if strings.Contains(topBarrier.EffectiveLevers[i].BrainStateTarget, "lower") ||
				strings.Contains(topBarrier.EffectiveLevers[i].BrainStateTarget, "calm") {
				return &topBarrier.EffectiveLevers[i]
			}
		}
	}

	// Return first effective lever for the barrier
	if len(topBarrier.EffectiveLevers) > 0 {
		return &topBarrier.EffectiveLevers[0]
	}

	return nil
}

func (o *Orchestrator) generateResponse(
	intervention *etp.InterventionLever,
	context etp.StudentContext,
	barriers []barriers.DetectedBarrier,
) string {
	if intervention == nil {
		return "I'm here to help. What would you like to work on?"
	}

	// Generate response based on intervention steps
	if len(intervention.Steps) > 0 {
		return intervention.Steps[0]
	}

	return intervention.Description
}

func (o *Orchestrator) checkRewardEarned(
	message string,
	context etp.StudentContext,
	intervention *etp.InterventionLever,
) bool {
	// Simple heuristic: longer messages = more engagement = reward
	// In real implementation, track task completion
	return len(message) > 50 && !strings.Contains(strings.ToLower(message), "i don't know")
}
