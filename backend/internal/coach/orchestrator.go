// filepath: /Users/mike5tew/Coding/humanOS/backend/internal/coach/orchestrator.go
package coach

import (
	"fmt"
	"strings"
	"time"

	"github.com/mike5tew/humanos/internal/barriers"
	"github.com/mike5tew/humanos/internal/etp"
	"github.com/mike5tew/humanos/internal/safeguarding"
)

// Orchestrator coordinates all HumanOS components
type Orchestrator struct {
	barrierDetector *barriers.BarrierDetector
	traumaDetector  *safeguarding.TraumaDetector
	ageFilter       *barriers.AgeAppropriateness
}

// CoachResponse is what gets sent back to frontend
type CoachResponse struct {
	Message           string                 `json:"message"`
	Intervention      *etp.InterventionLever `json:"intervention,omitempty"`
	DetectedBarriers  []string               `json:"detected_barriers"`
	SafeguardingAlert bool                   `json:"safeguarding_alert"`
	RewardEarned      bool                   `json:"reward_earned"`
	Reasoning         []string               `json:"reasoning"`
	Timestamp         string                 `json:"timestamp"`
}

// NewOrchestrator creates orchestrator with all components
func NewOrchestrator(barriersPath, traumaPath, agePath string) (*Orchestrator, error) {
	bd, err := barriers.NewBarrierDetector(barriersPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load barriers: %w", err)
	}

	td, err := safeguarding.NewTraumaDetector(traumaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load trauma detector: %w", err)
	}

	af, err := barriers.NewAgeAppropriateness(agePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load age filter: %w", err)
	}

	return &Orchestrator{
		barrierDetector: bd,
		traumaDetector:  td,
		ageFilter:       af,
	}, nil
}

// ProcessMessage is the main workflow
func (o *Orchestrator) ProcessMessage(
	studentID string,
	message string,
	context etp.StudentContext,
) (*CoachResponse, error) {

	reasoning := []string{}

	// STEP 1: Trauma/safeguarding check (HIGHEST PRIORITY) - NOW IN BACKEND
	traumaResult := o.traumaDetector.Scan(message, context.Age)
	if traumaResult.Severity >= 3 {
		// Escalation already handled in trauma_detector
		safeguardingMsg := o.ageFilter.SafeguardingResponse(context.Age)
		reasoning = append(reasoning, "⚠️ Safeguarding concern - human team notified")

		return &CoachResponse{
			Message:           safeguardingMsg,
			SafeguardingAlert: true,
			DetectedBarriers:  []string{},
			Reasoning:         reasoning,
			Timestamp:         time.Now().Format(time.RFC3339),
		}, nil
	}

	// STEP 2: Detect barriers
	detectedBarriers := o.barrierDetector.DetectBarriers(message, context)

	if len(detectedBarriers) > 0 {
		topBarrier := detectedBarriers[0]
		reasoning = append(reasoning,
			fmt.Sprintf("🎯 Detected: %s (%.0f%%)",
				topBarrier.Barrier.Name, topBarrier.Confidence*100))
		reasoning = append(reasoning, topBarrier.Reasoning...)
	}

	// STEP 3: Select intervention based on barrier + brain state
	intervention := o.selectIntervention(detectedBarriers, context)
	if intervention != nil {
		reasoning = append(reasoning,
			fmt.Sprintf("💡 Intervention: %s", intervention.Name))
	}

	// STEP 4: Generate response using intervention strategy
	rawResponse := o.generateResponse(intervention, context, detectedBarriers)

	// STEP 5: Make response age-appropriate
	finalResponse := o.ageFilter.AdjustLanguage(rawResponse, context.Age)

	// STEP 6: Check for offense risk
	offenseRisks := o.ageFilter.CheckOffenseRisk(finalResponse, context.Age)
	if len(offenseRisks) > 0 {
		reasoning = append(reasoning, "⚠️ Regenerating safer response...")
		finalResponse = o.regenerateSafeResponse(context, intervention)
	}

	// STEP 7: Check if reward earned
	rewardEarned := o.checkRewardEarned(message, detectedBarriers)
	if rewardEarned {
		reasoning = append(reasoning, "🎮 Play break earned!")
	}

	return &CoachResponse{
		Message:          finalResponse,
		Intervention:     intervention,
		DetectedBarriers: extractBarrierNames(detectedBarriers),
		RewardEarned:     rewardEarned,
		Reasoning:        reasoning,
		Timestamp:        time.Now().Format(time.RFC3339),
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

	// High emotional voltage → calming intervention
	if context.BrainState.EmotionalLevel > 0.7 {
		for i := range topBarrier.EffectiveLevers {
			lever := &topBarrier.EffectiveLevers[i]
			if strings.Contains(strings.ToLower(lever.BrainStateTarget), "lower") {
				return lever
			}
		}
	}

	// Return first effective lever
	if len(topBarrier.EffectiveLevers) > 0 {
		return &topBarrier.EffectiveLevers[0]
	}

	return nil
}

// generateResponse creates response from intervention strategy
func (o *Orchestrator) generateResponse(
	intervention *etp.InterventionLever,
	context etp.StudentContext,
	detectedBarriers []barriers.DetectedBarrier,
) string {

	// PRIORITY 1: No barriers detected = positive engagement
	if len(detectedBarriers) == 0 {
		return o.generatePositiveEngagementResponse(context)
	}

	// PRIORITY 2: Barrier-specific response (most contextual)
	if len(detectedBarriers) > 0 {
		barrierID := detectedBarriers[0].Barrier.ID
		if response, ok := o.getBarrierSpecificResponse(barrierID); ok {
			return response
		}
	}

	// PRIORITY 3: Intervention-based response
	if intervention != nil {
		if intervention.Description != "" {
			return o.translateStepToResponse(intervention.Description)
		}
		if len(intervention.Steps) > 0 {
			return o.translateStepToResponse(intervention.Steps[0])
		}
	}

	// PRIORITY 4: Fallback generic
	return "I'm here to help. What would you like to work on?"
}

// generatePositiveEngagementResponse handles cases where student is genuinely engaged
func (o *Orchestrator) generatePositiveEngagementResponse(context etp.StudentContext) string {
	responses := []string{
		"That's exactly the kind of thinking I want to see! You're really understanding this.",
		"Excellent question! It shows you're thinking deeply about the material.",
		"You're making real progress. Let's keep building on this momentum.",
		"I can tell you're genuinely engaged. That's how real learning happens!",
		"Perfect! You're asking the right questions. Let's explore this further.",
	}

	// Vary response based on age
	if context.Age < 10 {
		responses = []string{
			"That's great thinking! You're doing really well.",
			"I love that you want to learn more! Keep going!",
			"You're asking smart questions. That's how you get better!",
		}
	}

	// Return a random response from the list
	return responses[len(responses)%len(responses)]
}

// getBarrierSpecificResponse returns contextual response based on barrier type
// Maps barrier IDs directly to responses (no context needed, just the ID)
func (o *Orchestrator) getBarrierSpecificResponse(barrierID string) (string, bool) {
	responses := map[string]string{
		"lack_of_motivation":         "How about we try something really simple first? Just to get warmed up. You might surprise yourself!",
		"confrontational_showoff":    "I can see you're feeling frustrated with this. That's totally okay. Let's find something small you're doing well with.",
		"silent_avoider":             "No pressure at all. I'm right here with you. Can you just try one tiny thing - even just writing one word?",
		"quiet_playful_avoider":      "I like your energy! Let's turn that into something fun AND educational. Ready for a challenge?",
		"high_achiever_underengaged": "You're clearly capable of more. Let me show you something that will actually challenge you. This gets interesting.",
	}

	if response, exists := responses[barrierID]; exists {
		return response, true
	}

	return "", false
}

// translateStepToResponse converts intervention step/description to actual response
func (o *Orchestrator) translateStepToResponse(step string) string {
	if step == "" {
		return "I'm here to help. What would you like to work on?"
	}

	step_lower := strings.ToLower(step)

	// Pattern-based matching for common intervention phrases
	if strings.Contains(step_lower, "find") || strings.Contains(step_lower, "recognize") {
		return "I can see you're working on this. Let's find something small you're doing well!"
	}
	if strings.Contains(step_lower, "lower") || strings.Contains(step_lower, "easy") {
		return "How about we try something super easy first? Just to get warmed up."
	}
	if strings.Contains(step_lower, "show interest") || strings.Contains(step_lower, "chat") {
		return "Before we dive in, how are you feeling today? Anything on your mind?"
	}
	if strings.Contains(step_lower, "guide") || strings.Contains(step_lower, "help") {
		return "Let me help you out here. Can you tell me just one thing you think about this topic?"
	}
	if strings.Contains(step_lower, "celebrate") || strings.Contains(step_lower, "praise") {
		return "That's a great start! You're doing really well just by being here and trying."
	}
	if strings.Contains(step_lower, "game") || strings.Contains(step_lower, "reward") {
		return "Complete this and you'll get a reward. Let's see what you can do!"
	}
	if strings.Contains(step_lower, "micro") || strings.Contains(step_lower, "tiny") {
		return "Let's start with something really easy - just to warm up. You've got this!"
	}
	if strings.Contains(step_lower, "shoulder") || strings.Contains(step_lower, "right here") {
		return "I'm right here with you. Let's start with something tiny - what's one thing you could try?"
	}

	// Fallback: use step text directly with supportive framing
	return "Let's work through this together. " + step
}

func (o *Orchestrator) regenerateSafeResponse(
	context etp.StudentContext,
	intervention *etp.InterventionLever,
) string {

	// Age-appropriate safe fallbacks
	if context.Age < 8 {
		return "Let's try something fun! What would you like to do?"
	} else if context.Age < 13 {
		return "I'm here to help you learn. What part are you finding tricky?"
	} else {
		return "Let's work through this together. Where would you like to start?"
	}
}

func (o *Orchestrator) checkRewardEarned(
	message string,
	barriers []barriers.DetectedBarrier,
) bool {

	// Longer, engaged responses = reward
	if len(message) > 50 {
		// Check if not pure avoidance
		for _, b := range barriers {
			if strings.Contains(b.Barrier.ID, "lack_of_motivation") ||
				strings.Contains(b.Barrier.ID, "confrontational") {
				return false
			}
		}
		return true
	}

	return false
}

func extractBarrierNames(barriers []barriers.DetectedBarrier) []string {
	names := make([]string, len(barriers))
	for i, b := range barriers {
		names[i] = b.Barrier.Name
	}
	return names
}
