package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Core types
type ETPFramework struct {
	CoreETPs map[string][]string
	Weights  map[string]float64
}

type EmotionalContext struct {
	Situation     string
	ActivatedETPs []string
	Scores        map[string]float64
}

type LLMResponse struct {
	Text         string
	Score        float64
	ETPAlignment map[string]float64
}

type BrainState struct {
	PrimalLevel    float64 // 0-1
	EmotionalLevel float64 // 0-1
	RationalLevel  float64 // 0-1
	OverrideRisk   float64 // 0-1
}

// TODO: BarrierProfile for savant hypothesis modeling
// Models which emotional barriers suppress which cognitive domains.
// Theory: reduced barriers in specific areas → specialized excellence (savant-like ability).
// Future: track barrier strength per domain, design interventions to selectively reduce.
type BarrierProfile struct {
	Domain           string   // "math", "music", "visual-spatial", "language", etc.
	BarrierStrength  float64  // 0 (no barrier, savant-like) to 1 (fully suppressed)
	EmotionalSources []string // Which ETPs create this barrier: "fear_of_failure", "social_pressure"
}

// TODO: NeuroProfile for neurodiversity-aware modeling
// Models neurological hardware differences (nature) vs learned software (nurture).
// Autism example: heightened sensory fence voltage, exceptional pattern recognition.
type NeuroProfile struct {
	NeurologyType    string             // "neurotypical", "autistic", "adhd", etc.
	SensoryFenceMap  map[string]float64 // Which stimuli trigger overwhelm: "loud_sounds": 0.9
	ProcessingStyle  string             // "detail-focused", "big-picture", "sequential"
	InnateStrengths  []string           // Default "keys" in keyring: "pattern_recognition"
	DevelopmentNeeds []string           // Missing keys: "social_navigation", "executive_function"
}

// TODO: KeyringProfile for options-based education modeling
// Models available capabilities (keys) that open life paths (doors).
// Success = breadth of accessible options, not standardized outcomes.
type KeyringProfile struct {
	AcademicKeys   []string // "math", "reading", "science_reasoning"
	SocialKeys     []string // "empathy", "negotiation", "group_work"
	EmotionalKeys  []string // "self_regulation", "resilience", "emotional_awareness"
	CreativeKeys   []string // "problem_solving", "artistic_expression", "innovation"
	OptionsBreadth float64  // 0-1 score: how many life paths are accessible with current keyring
}

// TODO: implement PersonalityProfile with barrier configurations
// type PersonalityProfile struct {
//     DominantETPs     []string
//     BarrierProfiles  []BarrierProfile  // Savant hypothesis: map emotional→cognitive constraints
//     PowerNeedBalance [2]float64        // [power_focus, need_focus] 0-1 each
//     NeuroProfile     NeuroProfile      // Neurodiversity: hardware + processing style
//     Keyring          KeyringProfile    // Options philosophy: current capabilities
// }

// NewETPFramework: initialize a minimal ETP taxonomy and weights
func NewETPFramework() *ETPFramework {
	return &ETPFramework{
		CoreETPs: map[string][]string{
			"reward_seeking":     {"achievement", "pleasure", "excitement", "hope", "curiosity"},
			"threat_avoidance":   {"fear", "anxiety", "stress", "dread", "overwhelm"},
			"social_bonding":     {"belonging", "love", "empathy", "rejection", "loneliness"},
			"status_competition": {"power", "pride", "envy", "shame", "competition", "aggression"},
			"physical_drives":    {"hunger", "tiredness", "pain", "lust", "comfort"},
		},
		Weights: map[string]float64{
			"threat_avoidance":   0.25,
			"physical_drives":    0.20,
			"social_bonding":     0.15,
			"reward_seeking":     0.12,
			"status_competition": 0.08,
		},
	}
}

// AnalyzeEmotionalContext: keyword-based mapping to a small set of ETPs.
// TODO: replace with NLP/embedding-backed extractor.
func (e *ETPFramework) AnalyzeEmotionalContext(situation string) *EmotionalContext {
	ctx := &EmotionalContext{
		Situation: situation,
		Scores:    make(map[string]float64),
	}

	// Physiological amplifiers
	amplifiers := e.detectPhysiologicalAmplifiers(situation)
	for _, a := range amplifiers {
		ctx.Scores[a] = 0.6
	}

	switch {
	case contains(situation, []string{"failed", "rejected", "lost", "missed"}):
		ctx.ActivatedETPs = []string{"rejection", "sadness", "fear"}
		ctx.Scores["rejection"] = 0.8
		ctx.Scores["sadness"] = 0.6
		ctx.Scores["fear"] = 0.4

	case contains(situation, []string{"succeeded", "won", "achieved", "completed"}):
		ctx.ActivatedETPs = []string{"achievement", "happiness", "excitement"}
		ctx.Scores["achievement"] = 0.9
		ctx.Scores["happiness"] = 0.7
		ctx.Scores["excitement"] = 0.5

	case contains(situation, []string{"danger", "threat", "risk", "scared"}):
		ctx.ActivatedETPs = []string{"fear", "protective_instinct", "bravery"}
		ctx.Scores["fear"] = 0.9
		ctx.Scores["protective_instinct"] = 0.6
		ctx.Scores["bravery"] = 0.3

	default:
		// If physiological amplifiers exist, include them as activated ETPs
		if len(amplifiers) > 0 {
			ctx.ActivatedETPs = append(ctx.ActivatedETPs, amplifiers...)
			// also add a neutral fallback
			ctx.ActivatedETPs = append(ctx.ActivatedETPs, "neutral")
			if _, ok := ctx.Scores["neutral"]; !ok {
				ctx.Scores["neutral"] = 0.4
			}
		} else {
			ctx.ActivatedETPs = []string{"neutral"}
			ctx.Scores["neutral"] = 0.5
		}
	}

	return ctx
}

// detectPhysiologicalAmplifiers: basic checks for hunger/tiredness
func (e *ETPFramework) detectPhysiologicalAmplifiers(situation string) []string {
	amps := []string{}
	if contains(situation, []string{"hungry", "starving", "hangry", "need food", "empty stomach"}) {
		amps = append(amps, "hunger")
	}
	if contains(situation, []string{"tired", "exhausted", "sleepy", "burned out", "fatigued"}) {
		amps = append(amps, "tiredness")
	}
	return amps
}

// CalculateBrainState: collapse ETP scores into primal/emotional/rational values
func (e *ETPFramework) CalculateBrainState(ctx *EmotionalContext) BrainState {
	state := BrainState{}

	// primitive/primal ETPs
	primalETPs := []string{"hunger", "tiredness", "pain", "physical_discomfort"}
	for _, etp := range primalETPs {
		if s, ok := ctx.Scores[etp]; ok {
			state.PrimalLevel += s
		}
	}

	// emotional/lignbic ETPs
	emotionalETPs := []string{"fear", "anger", "frustration", "overwhelm", "sadness"}
	for _, etp := range emotionalETPs {
		if s, ok := ctx.Scores[etp]; ok {
			state.EmotionalLevel += s
		}
	}

	// clamp levels to 0..1
	state.PrimalLevel = math.Min(1.0, state.PrimalLevel)
	state.EmotionalLevel = math.Min(1.0, state.EmotionalLevel)

	// rational capacity inversely related to load
	state.RationalLevel = math.Max(0.0, 1.0-(state.PrimalLevel+state.EmotionalLevel)*0.8)
	// override risk heuristic
	state.OverrideRisk = (state.PrimalLevel * 0.6) + (state.EmotionalLevel * 0.4)

	return state
}

// GetResponseStrategy: pick high-level strategy from brain state
func (e *ETPFramework) GetResponseStrategy(bs BrainState) string {
	if bs.OverrideRisk > 0.7 {
		return "primal_first"
	} else if bs.OverrideRisk > 0.4 {
		return "emotional_first"
	}
	return "rational_engagement"
}

// FilterLLMResponses: score and rank plain responses based on keyword alignment
func (e *ETPFramework) FilterLLMResponses(responses []string, ctx *EmotionalContext) []LLMResponse {
	out := make([]LLMResponse, 0, len(responses))
	for _, r := range responses {
		score, align := e.scoreResponse(r, ctx)
		out = append(out, LLMResponse{
			Text:         r,
			Score:        score,
			ETPAlignment: align,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}

// scoreResponse: simple alignment score by counting keyword matches weighted by context scores
func (e *ETPFramework) scoreResponse(response string, ctx *EmotionalContext) (float64, map[string]float64) {
	alignment := make(map[string]float64)
	total := 0.0
	for _, etp := range ctx.ActivatedETPs {
		weight := 1.0
		if s, ok := ctx.Scores[etp]; ok {
			weight = s
		}
		etpScore := e.calculateETPAlignment(response, etp) * weight
		alignment[etp] = etpScore
		total += etpScore
	}
	// small normalization
	return total, alignment
}

// calculateETPAlignment: baseline keyword map for matching
func (e *ETPFramework) calculateETPAlignment(response string, etp string) float64 {
	keywordMap := map[string][]string{
		"rejection":   {"sorry", "understand", "difficult", "better luck"},
		"achievement": {"congratulations", "well done", "great work", "proud"},
		"fear":        {"safe", "protected", "secure", "it's okay", "you're safe"},
		"sadness":     {"sorry", "understand", "here for you", "support"},
		"happiness":   {"happy", "excited", "celebrate", "wonderful"},
		"hunger":      {"eat", "food", "hungry", "eat something", "grab a snack"},
		"tiredness":   {"rest", "sleep", "take a break", "get some rest"},
	}

	kw, ok := keywordMap[etp]
	if !ok {
		// baseline small score so non-empty responses get some credit
		return 0.05
	}
	matches := 0
	for _, k := range kw {
		if contains(response, []string{k}) {
			matches++
		}
	}
	return float64(matches) / float64(len(kw))
}

// contains helper
func contains(s string, subs []string) bool {
	s = strings.ToLower(s)
	for _, sub := range subs {
		if strings.Contains(s, strings.ToLower(sub)) {
			return true
		}
	}
	return false
}

// Example CLI/demo
func main() {
	etp := NewETPFramework()

	// Example situation
	situation := "I'm exhausted and failed my interview, I feel rejected and can't think straight"
	ctx := etp.AnalyzeEmotionalContext(situation)
	fmt.Printf("Situation: %s\nActivated ETPs: %v\nScores: %v\n\n", ctx.Situation, ctx.ActivatedETPs, ctx.Scores)

	brain := etp.CalculateBrainState(ctx)
	fmt.Printf("BrainState: Primal=%.2f Emotional=%.2f Rational=%.2f OverrideRisk=%.2f\nStrategy: %s\n\n",
		brain.PrimalLevel, brain.EmotionalLevel, brain.RationalLevel, brain.OverrideRisk, etp.GetResponseStrategy(brain))

	candidates := []string{
		"That's terrible news. You must feel awful.",
		"Don't worry, you'll get the next one!",
		"Let me explain why you probably failed...",
		"I understand this is disappointing. Would you like some advice for next time?",
		"Maybe take a rest and then we can plan some actionable next steps.",
		"Grab some food and sleep, you'll be clearer tomorrow.",
	}

	ranked := etp.FilterLLMResponses(candidates, ctx)
	fmt.Println("Ranked responses:")
	for i, r := range ranked {
		fmt.Printf("%d) Score: %.2f - %s\n   Alignment: %v\n", i+1, r.Score, r.Text, r.ETPAlignment)
	}

	// TODOs:
	// - Replace keyword-based functions with NLP/embeddings and vector DB integration.
	// - Add PersonalityProfile, persistent local profiles, and federated update hooks.
	// - Add unit tests and small dataset for evaluation.
	// - Implement BarrierProfile modeling (savant hypothesis): track which emotional barriers suppress cognitive domains.
	// - Research: can targeted ETP work (reducing specific fears/pressures) unlock latent abilities?
	// - Implement NeuroProfile for neurodiversity-aware interventions (voltage regulation, exposure pathways).
	// - Build KeyringProfile assessment tool: measure options breadth as success metric.
	// - Design autism-specific intervention: identify sensory fence map → lower voltage → gradual exposure → leverage strengths.
}
