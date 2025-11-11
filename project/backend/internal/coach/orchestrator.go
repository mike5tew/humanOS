package coach
package coach

import (
    "github.com/michaelstewart/humanos/internal/barriers"
    "github.com/michaelstewart/humanos/internal/etp"
    "github.com/michaelstewart/humanos/internal/safeguarding"
)

type CoachOrchestrator struct {
    barrierDetector *barriers.BarrierDetector
    traumaDetector  *safeguarding.TraumaDetector
    // Add other components as needed
}

func NewCoachOrchestrator(barriersPath, traumaPath string) (*CoachOrchestrator, error) {
    bd, err := barriers.NewBarrierDetector(barriersPath)
    if err != nil {
        return nil, err
    }

    td, err := safeguarding.NewTraumaDetector(traumaPath)
    if err != nil {
        return nil, err
    }

    return &CoachOrchestrator{
        barrierDetector: bd,
        traumaDetector:  td,
    }, nil
}

func (co *CoachOrchestrator) ProcessStudentMessage(
    studentID string,
    message string,
    context etp.StudentContext,
) (etp.CoachResponse, error) {
    reasoning := []string{}

    // 1. Check for trauma/safeguarding issues (highest priority)
    traumaResult, err := co.traumaDetector.Scan(message, context.Age)
    if err != nil {
        return etp.CoachResponse{}, err
    }

    if traumaResult.Severity >= 3 {
        reasoning = append(reasoning, "Safeguarding concern detected - escalating to human team")
        return etp.CoachResponse{
            Message:           co.generateSafeguardingResponse(context.Age),
            Intervention:      nil,
            DetectedBarriers:  []etp.StudentBarrier{},
            SafeguardingAlert: true,
            RewardEarned:      false,
            Reasoning:         reasoning,
        }, nil
    }

    // 2. Detect barriers
    detectedBarriers := co.barrierDetector.DetectBarriers(message, context)
    if len(detectedBarriers) > 0 {
        reasoning = append(reasoning, "Detected barrier: "+detectedBarriers[0].Barrier.Name)
    }

    // 3. Select appropriate intervention
    intervention := co.selectIntervention(detectedBarriers, context)
    if intervention != nil {
        reasoning = append(reasoning, "Selected intervention: "+intervention.Name)
    }

    // 4. Generate response
    responseMessage := co.generateResponse(intervention, context, detectedBarriers)
    
    // 5. Check if reward earned
    rewardEarned := co.checkRewardEarned(message, context, intervention)
    if rewardEarned {
        reasoning = append(reasoning, "Student earned play break reward")
    }

    barriers := make([]etp.StudentBarrier, len(detectedBarriers))
    for i, db := range detectedBarriers {
        barriers[i] = db.Barrier
    }

    return etp.CoachResponse{
        Message:           responseMessage,
        Intervention:      intervention,
        DetectedBarriers:  barriers,
        SafeguardingAlert: traumaResult.Severity > 0,
        RewardEarned:      rewardEarned,
        Reasoning:         reasoning,
    }, nil
}

func (co *CoachOrchestrator) selectIntervention(
    detectedBarriers []etp.DetectedBarrier,
    context etp.StudentContext,
) *etp.InterventionLever {
    if len(detectedBarriers) == 0 {
        return nil
    }

    topBarrier := detectedBarriers[0].Barrier

    // Check brain state - if emotional voltage too high, prioritize calming
    if context.BrainState.EmotionalLevel > 0.7 {
        // Find calming intervention
        for _, lever := range topBarrier.EffectiveLevers {
            if strings.Contains(lever.BrainStateTarget, "lower") ||
               strings.Contains(lever.BrainStateTarget, "calm") {
                return &lever
            }
        }
    }

    // Return first effective lever for the barrier
    if len(topBarrier.EffectiveLevers) > 0 {
        return &topBarrier.EffectiveLevers[0]
    }

    return nil
}

func (co *CoachOrchestrator) generateResponse(
    intervention *etp.InterventionLever,
    context etp.StudentContext,
    barriers []etp.DetectedBarrier,
) string {
    if intervention == nil {
        return "I'm here to help. What would you like to work on?"
    }

    if len(intervention.Steps) > 0 {
        return intervention.Steps[0]
    }

    return intervention.Description
}

func (co *CoachOrchestrator) generateSafeguardingResponse(age int) string {
    // Age-appropriate safeguarding response
    if age < 10 {
        return "Let's talk to a trusted adult about this."
    }
    return "I think it would be helpful to talk to someone who can support you better."
}

func (co *CoachOrchestrator) checkRewardEarned(
    message string,
    context etp.StudentContext,
    intervention *etp.InterventionLever,
) bool {
    // Simple heuristic: longer messages = more engagement = reward
    return len(message) > 50 && !strings.Contains(strings.ToLower(message), "i don't know")
}
