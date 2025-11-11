package etp
package etp

import "time"

type EngagementMode string

const (
    ComplianceModeStr EngagementMode = "compliance"
    EngagementModeStr EngagementMode = "engagement"
    ResistanceModeStr EngagementMode = "resistance"
)

type BarrierCategory string

const (
    AcuteBarrier     BarrierCategory = "acute"
    ChronicBarrier   BarrierCategory = "chronic"
    StructuralBarrier BarrierCategory = "structural"
)

type BubbleType string

const (
    SocialBubble   BubbleType = "social"
    SensoryBubble  BubbleType = "sensory"
    AutonomyBubble BubbleType = "autonomy"
    StatusBubble   BubbleType = "status"
)

type ETP struct {
    Name      string  `json:"name"`
    Category  string  `json:"category"` // pain, pleasure, social, goal
    Intensity float64 `json:"intensity"` // 0-1
}

type BrainState struct {
    PrimalLevel    float64 `json:"primal_level"`    // 0-1
    EmotionalLevel float64 `json:"emotional_level"` // 0-1
    RationalLevel  float64 `json:"rational_level"`  // 0-1
    CurrentMode    string  `json:"current_mode"`    // primal, emotional, rational
}

type RoutineProfile struct {
    RoutineDependency float64 `json:"routine_dependency"` // 0-1
    ThinkingAtrophy   float64 `json:"thinking_atrophy"`   // 0-1
    FenceVoltage      float64 `json:"fence_voltage"`      // 0-1
}

type StudentContext struct {
    StudentID         string         `json:"student_id"`
    Age               int            `json:"age"`
    BrainState        BrainState     `json:"brain_state"`
    ActivatedETPs     []ETP          `json:"activated_etps"`
    RoutineProfile    RoutineProfile `json:"routine_profile"`
    SocialNeed        float64        `json:"social_need"`        // 0-1
    AutonomyResistance float64       `json:"autonomy_resistance"` // 0-1
    StatusSeeking     float64        `json:"status_seeking"`     // 0-1
}

type InterventionLever struct {
    Name            string   `json:"name"`
    Description     string   `json:"description"`
    Steps           []string `json:"steps"`
    Prerequisites   []string `json:"prerequisites,omitempty"`
    Benefits        []string `json:"benefits,omitempty"`
    ETPReduction    []string `json:"etp_reduction"`
    BrainStateTarget string  `json:"brain_state_target"`
    WhenToUse       []string `json:"when_to_use,omitempty"`
}

type StudentBarrier struct {
    ID               string              `json:"id"`
    Name             string              `json:"name"`
    Category         BarrierCategory     `json:"category"`
    Description      string              `json:"description"`
    ActivatedETPs    []string            `json:"activated_etps"`
    AvoidanceTactics []string            `json:"avoidance_tactics"`
    EffectiveLevers  []InterventionLever `json:"effective_levers"`
    UnderlyingCause  string              `json:"underlying_cause"`
}

type Bubble struct {
    Type            BubbleType `json:"type"`
    Description     string     `json:"description"`
    ControlValue    float64    `json:"control_value"`    // 0-1
    SocialValue     float64    `json:"social_value"`     // 0-1
    DistractionRisk float64    `json:"distraction_risk"` // 0-1
}

type TrustBargain struct {
    BubbleGranted       Bubble   `json:"bubble_granted"`
    ExpectedBehavior    string   `json:"expected_behavior"`
    RiskMitigation      []string `json:"risk_mitigation"`
    EmotionalIntelConvo string   `json:"emotional_intel_convo"`
}

type DetectedBarrier struct {
    Barrier    StudentBarrier `json:"barrier"`
    Confidence float64        `json:"confidence"` // 0-1
    Reasoning  []string       `json:"reasoning"`
}

type CoachResponse struct {
    Message           string              `json:"message"`
    Intervention      *InterventionLever  `json:"intervention"`
    DetectedBarriers  []StudentBarrier    `json:"detected_barriers"`
    SafeguardingAlert bool                `json:"safeguarding_alert"`
    RewardEarned      bool                `json:"reward_earned"`
    Reasoning         []string            `json:"reasoning"`
}
