package etp

import "time"

// EngagementMode represents student's current engagement state
type EngagementMode string

const (
	ComplianceModeStr EngagementMode = "compliance" // Following routine to avoid discomfort
	EngagementModeStr EngagementMode = "engagement" // Curiosity-driven active thinking
	ResistanceModeStr EngagementMode = "resistance" // Actively avoiding both
)

// RoutineProfile tracks student's dependency on structure
type RoutineProfile struct {
	RoutineDependency float64 `json:"routine_dependency"` // 0 (self-directed) to 1 (needs complete structure)
	ThinkingAtrophy   float64 `json:"thinking_atrophy"`   // 0 (thinking is easy) to 1 (thinking is painful)
	FenceVoltage      float64 `json:"fence_voltage"`      // Pain level when routine is disrupted
}

// BrainState represents primal/emotional/rational levels
type BrainState struct {
	PrimalLevel    float64 `json:"primal_level"`    // 0-1: Physical needs (hunger, tiredness)
	EmotionalLevel float64 `json:"emotional_level"` // 0-1: Fear, frustration, overwhelm
	RationalLevel  float64 `json:"rational_level"`  // 0-1: Cortical thinking capacity
	OverrideRisk   float64 `json:"override_risk"`   // 0-1: Risk of brain stem takeover
}

// ETP represents an Emotional Trigger Point
type ETP struct {
	Name      string  `json:"name"`
	Category  string  `json:"category"`  // pain, pleasure, social, goal
	Intensity float64 `json:"intensity"` // 0-1
}

// StudentContext contains full student state for intervention selection
type StudentContext struct {
	StudentID          string         `json:"student_id"`
	Age                int            `json:"age"`
	BrainState         BrainState     `json:"brain_state"`
	ActivatedETPs      []ETP          `json:"activated_etps"`
	RoutineProfile     RoutineProfile `json:"routine_profile"`
	SocialNeed         float64        `json:"social_need"`
	AutonomyResistance float64        `json:"autonomy_resistance"`
	StatusSeeking      float64        `json:"status_seeking"`
}

// InterventionLever represents a specific intervention strategy
type InterventionLever struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Steps            []string `json:"steps"`
	Prerequisites    []string `json:"prerequisites,omitempty"`
	Benefits         []string `json:"benefits,omitempty"`
	ETPReduction     []string `json:"etp_reduction"`
	BrainStateTarget string   `json:"brain_state_target"`
	WhenToUse        []string `json:"when_to_use,omitempty"`
	TimeInvestment   string   `json:"time_investment,omitempty"`
	ExpectedTimeline string   `json:"expected_timeline,omitempty"`
}

// BarrierCategory classifies barrier types
type BarrierCategory string

const (
	AcuteBarrier      BarrierCategory = "acute"
	ChronicBarrier    BarrierCategory = "chronic"
	StructuralBarrier BarrierCategory = "structural"
	EnrichmentBarrier BarrierCategory = "enrichment"
)

// StudentBarrier represents a learning obstacle
type StudentBarrier struct {
	ID               string              `json:"id"` // REQUIRED: unique identifier for barrier matching
	Name             string              `json:"name"`
	Category         BarrierCategory     `json:"category"`
	Description      string              `json:"description"`
	ActivatedETPs    []string            `json:"activated_etps"`
	AvoidanceTactics []string            `json:"avoidance_tactics"`
	EffectiveLevers  []InterventionLever `json:"effective_levers"`
	UnderlyingCause  string              `json:"underlying_cause"`
}

// BubbleType represents types of student control zones
type BubbleType string

const (
	SocialBubble   BubbleType = "social"
	SensoryBubble  BubbleType = "sensory"
	AutonomyBubble BubbleType = "autonomy"
	StatusBubble   BubbleType = "status"
)

// Bubble represents student's self-created control zone
type Bubble struct {
	Type            BubbleType `json:"type"`
	Description     string     `json:"description"`
	ControlValue    float64    `json:"control_value"`    // 0-1
	SocialValue     float64    `json:"social_value"`     // 0-1
	DistractionRisk float64    `json:"distraction_risk"` // 0-1
}

// TrustBargain represents granting bubbles strategically
type TrustBargain struct {
	BubbleGranted       Bubble   `json:"bubble_granted"`
	ExpectedBehavior    string   `json:"expected_behavior"`
	RiskMitigation      []string `json:"risk_mitigation"`
	EmotionalIntelConvo string   `json:"emotional_intel_convo"`
}

// CoachResponse is the complete AI coach output
type CoachResponse struct {
	Message           string             `json:"message"`
	Intervention      *InterventionLever `json:"intervention"`
	DetectedBarriers  []StudentBarrier   `json:"detected_barriers"`
	SafeguardingAlert bool               `json:"safeguarding_alert"`
	RewardEarned      bool               `json:"reward_earned"`
	Reasoning         []string           `json:"reasoning"`
	Timestamp         time.Time          `json:"timestamp"`
}

// PlayBreakStage represents progression through play break graduation
type PlayBreakStage int

const (
	ConcentrationStage PlayBreakStage = iota // 5 min work → 5 min play
	RewardsStage                             // Work to break → play reward
	ExamPeriodStage                          // Sustained concentration
	PraiseStage                              // Intrinsic motivation
)

// PlayBreakProfile tracks student's progression through play break system
type PlayBreakProfile struct {
	StudentID       string         `json:"student_id"`
	CurrentStage    PlayBreakStage `json:"current_stage"`
	TimeInStage     int            `json:"time_in_stage"`    // Days
	SuccessRate     float64        `json:"success_rate"`     // 0-1
	WorkDuration    int            `json:"work_duration"`    // Minutes per session
	RewardsEarned   int            `json:"rewards_earned"`   // Count
	LastProgression time.Time      `json:"last_progression"` // When moved to current stage
	ReadyToAdvance  bool           `json:"ready_to_advance"` // Calculated flag
}

// TODO: Future types from project/backend/main.go ideas

// BarrierProfile models which emotional barriers suppress cognitive domains
// Theory: Savant hypothesis - reduced barriers → specialized excellence
type BarrierProfile struct {
	Domain           string   // "math", "music", "visual-spatial", "language"
	BarrierStrength  float64  // 0 (no barrier) to 1 (fully suppressed)
	EmotionalSources []string // Which ETPs create barrier: "fear_of_failure"
}

// VoltageProfile tracks emotional sensitivity
type VoltageProfile struct {
	BaselineVoltage float64 `json:"baseline_voltage"` // Resting fence voltage
	GeneticFactor   float64 `json:"genetic_factor"`   // Innate sensitivity
	LearnedFactor   float64 `json:"learned_factor"`   // Learned responses
}

// RelationshipPhase tracks trust building
type RelationshipPhase string

const (
	PhaseEarly       RelationshipPhase = "early"       // No relationship yet
	PhaseBuilding    RelationshipPhase = "building"    // Some trust forming
	PhaseEstablished RelationshipPhase = "established" // Trust exists
)

// NeuroProfile models neurological hardware differences (neurodiversity)
type NeuroProfile struct {
	NeurologyType    string             // "neurotypical", "autistic", "adhd"
	SensoryFenceMap  map[string]float64 // Sensory triggers: "loud_sounds": 0.9
	ProcessingStyle  string             // "detail-focused", "big-picture"
	InnateStrengths  []string           // Default capabilities
	DevelopmentNeeds []string           // Missing skills to develop
}

// KeyringProfile models available capabilities (options philosophy)
// Success = breadth of accessible life paths
type KeyringProfile struct {
	AcademicKeys   []string // "math", "reading", "science_reasoning"
	SocialKeys     []string // "empathy", "negotiation", "group_work"
	EmotionalKeys  []string // "self_regulation", "resilience"
	CreativeKeys   []string // "problem_solving", "innovation"
	OptionsBreadth float64  // 0-1: how many life paths accessible
}

// PersonalityProfile combines all profile dimensions
type PersonalityProfile struct {
	DominantETPs     []string
	BarrierProfiles  []BarrierProfile
	PowerNeedBalance [2]float64 // [power_focus, need_focus]
	NeuroProfile     NeuroProfile
	Keyring          KeyringProfile
}
