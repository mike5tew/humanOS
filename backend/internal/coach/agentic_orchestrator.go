package coach

import (
	"strings"

	"github.com/mike5tew/humanos/internal/barriers"
	"github.com/mike5tew/humanos/internal/etp"
	"github.com/mike5tew/humanos/internal/integration"
	"github.com/mike5tew/humanos/internal/safeguarding"
)

type AgenticOrchestrator struct {
	barrierDetector *barriers.BarrierDetector // Changed from BarrierDetector
	traumaDetector  *safeguarding.TraumaDetector
	chisgClient     *integration.CHISGClient
}

func NewAgenticOrchestrator(barriersPath, traumaPath string) (*AgenticOrchestrator, error) {
	bd, err := barriers.NewBarrierDetector(barriersPath) // Changed from NewBarrierDetector
	if err != nil {
		return nil, err
	}

	td, err := safeguarding.NewTraumaDetector(traumaPath)
	if err != nil {
		return nil, err
	}

	return &AgenticOrchestrator{
		barrierDetector: bd,
		traumaDetector:  td,
		chisgClient:     integration.NewCHISGClient(),
	}, nil
}

type AgenticResponse struct {
	Message           string                        `json:"message"`
	Intervention      *etp.InterventionLever        `json:"intervention"`
	DetectedBarriers  []etp.StudentBarrier          `json:"detected_barriers"`
	KnowledgeContext  *integration.KnowledgeContext `json:"knowledge_context"`
	FramingStrategy   string                        `json:"framing_strategy"`
	SafeguardingAlert bool                          `json:"safeguarding_alert"`
	RewardEarned      bool                          `json:"reward_earned"`
	Reasoning         []string                      `json:"reasoning"`
}

// ProcessStudentMessage orchestrates both emotional and knowledge analysis
func (ao *AgenticOrchestrator) ProcessStudentMessage(
	studentID string,
	message string,
	context etp.StudentContext,
) (AgenticResponse, error) {
	reasoning := []string{}

	// 1. Safeguarding check (highest priority)
	traumaResult, err := ao.traumaDetector.Scan(message, context.Age)
	if err != nil {
		return AgenticResponse{}, err
	}

	if traumaResult.Severity >= 3 {
		reasoning = append(reasoning, "Safeguarding concern - escalating")
		return AgenticResponse{
			Message:           ao.generateSafeguardingResponse(context.Age),
			SafeguardingAlert: true,
			Reasoning:         reasoning,
		}, nil
	}

	// 2. Detect emotional barriers (HumanOS core)
	detectedBarriers := ao.barrierDetector.DetectBarriers(message, context)
	if len(detectedBarriers) > 0 {
		reasoning = append(reasoning, "Barrier detected: "+detectedBarriers[0].Barrier.Name)
	}

	// 3. Extract topic from message for CHISG analysis
	topic := ao.extractTopic(message)
	var knowledgeContext *integration.KnowledgeContext

	if topic != "" {
		// Query CHISG for semantic understanding
		knowledgeContext, err = ao.chisgClient.GetKnowledgeContext(topic, 0.5) // TODO: actual student level
		if err != nil {
			reasoning = append(reasoning, "CHISG unavailable, proceeding with emotion-only analysis")
		} else {
			reasoning = append(reasoning, "CHISG analysis integrated")
		}
	}

	// 4. Select intervention based on barriers + knowledge gaps
	intervention := ao.selectIntervention(detectedBarriers, context, knowledgeContext)

	// 5. Determine framing strategy (which ETP lens to use)
	framingStrategy := ao.determineFraming(detectedBarriers, knowledgeContext, context)
	reasoning = append(reasoning, "Framing: "+framingStrategy)

	// 6. Generate response combining emotional + knowledge context
	responseMessage := ao.generateAgenticResponse(
		intervention,
		knowledgeContext,
		framingStrategy,
		context,
	)

	// 7. Check reward eligibility
	rewardEarned := ao.checkRewardEarned(message, context, intervention)

	barriers := make([]etp.StudentBarrier, len(detectedBarriers))
	for i, db := range detectedBarriers {
		barriers[i] = db.Barrier
	}

	return AgenticResponse{
		Message:          responseMessage,
		Intervention:     intervention,
		DetectedBarriers: barriers,
		KnowledgeContext: knowledgeContext,
		FramingStrategy:  framingStrategy,
		RewardEarned:     rewardEarned,
		Reasoning:        reasoning,
	}, nil
}

// selectIntervention combines barrier detection + knowledge gaps
func (ao *AgenticOrchestrator) selectIntervention(
	detectedBarriers []barriers.DetectedBarrier,
	context etp.StudentContext,
	knowledgeCtx *integration.KnowledgeContext,
) *etp.InterventionLever {
	// If no barriers, check if there are prerequisite gaps from CHISG
	if len(detectedBarriers) == 0 && knowledgeCtx != nil {
		if len(knowledgeCtx.PrerequisiteGaps) > 0 {
			// Student needs foundational knowledge first
			// Return intervention to address gaps
			return &etp.InterventionLever{
				Name:        "address_prerequisites",
				Description: "Fill foundational knowledge gaps before tackling main topic",
				Steps: []string{
					"Acknowledge the challenge",
					"Identify what's needed first: " + strings.Join(knowledgeCtx.PrerequisiteGaps, ", "),
					"Work on prerequisites before main topic",
				},
				BrainStateTarget: "build_confidence_through_mastery",
			}
		}
	}

	// Standard barrier-based intervention selection
	if len(detectedBarriers) == 0 {
		return nil
	}

	topBarrier := detectedBarriers[0].Barrier

	// High emotional voltage? Prioritize calming
	if context.BrainState.EmotionalLevel > 0.7 {
		for _, lever := range topBarrier.EffectiveLevers {
			if strings.Contains(lever.BrainStateTarget, "lower") ||
				strings.Contains(lever.BrainStateTarget, "calm") {
				return &lever
			}
		}
	}

	if len(topBarrier.EffectiveLevers) > 0 {
		return &topBarrier.EffectiveLevers[0]
	}

	return nil
}

// determineFraming decides which ETP lens to use based on context
func (ao *AgenticOrchestrator) determineFraming(
	barriers []barriers.DetectedBarrier,
	knowledgeCtx *integration.KnowledgeContext,
	context etp.StudentContext,
) string {
	// If student is anxious, use achievement/mastery framing
	if context.BrainState.EmotionalLevel > 0.6 {
		return "achievement_with_small_wins"
	}

	// If knowledge gaps detected, use curiosity framing
	if knowledgeCtx != nil && len(knowledgeCtx.PrerequisiteGaps) > 0 {
		return "curiosity_building_blocks"
	}

	// If confrontational barrier, use status framing
	if len(barriers) > 0 && barriers[0].Barrier.ID == "confrontational_showoff" {
		return "status_through_mastery"
	}

	// Default to achievement framing
	return "achievement_progress_tracking"
}

// generateAgenticResponse combines emotional + knowledge contexts
func (ao *AgenticOrchestrator) generateAgenticResponse(
	intervention *etp.InterventionLever,
	knowledgeCtx *integration.KnowledgeContext,
	framing string,
	context etp.StudentContext,
) string {
	if intervention == nil {
		if knowledgeCtx != nil {
			return "I can help you understand " + knowledgeCtx.Topic + ". What specifically would you like to explore?"
		}
		return "I'm here to help. What would you like to work on?"
	}

	// Combine intervention steps with knowledge context
	baseResponse := intervention.Steps[0]

	// Add CHISG-informed specifics if available
	if knowledgeCtx != nil && len(knowledgeCtx.PrerequisiteGaps) > 0 {
		baseResponse += "\n\nTo get there, let's first make sure you're solid on: " +
			strings.Join(knowledgeCtx.PrerequisiteGaps, ", ")
	}

	// Apply emotional framing
	switch framing {
	case "achievement_with_small_wins":
		baseResponse = "Let's break this into small wins. " + baseResponse
	case "curiosity_building_blocks":
		baseResponse = "This is like building blocks - each piece connects. " + baseResponse
	case "status_through_mastery":
		baseResponse = "When you nail this, you'll know more than most students. " + baseResponse
	}

	return baseResponse
}

// extractTopic attempts to identify the learning topic from student message
func (ao *AgenticOrchestrator) extractTopic(message string) string {
	// Simple keyword extraction (TODO: improve with NLP)
	lower := strings.ToLower(message)

	topics := map[string]string{
		"quadratic": "quadratic equations",
		"algebra":   "algebra",
		"fraction":  "fractions",
		"essay":     "essay writing",
		"dna":       "DNA structure",
		"cell":      "cell biology",
	}

	for keyword, topic := range topics {
		if strings.Contains(lower, keyword) {
			return topic
		}
	}

	return ""
}

func (ao *AgenticOrchestrator) generateSafeguardingResponse(age int) string {
	if age < 10 {
		return "Let's talk to a trusted adult about this."
	}
	return "I think it would be helpful to talk to someone who can support you better."
}

func (ao *AgenticOrchestrator) checkRewardEarned(
	message string,
	context etp.StudentContext,
	intervention *etp.InterventionLever,
) bool {
	return len(message) > 50 && !strings.Contains(strings.ToLower(message), "i don't know")
}
