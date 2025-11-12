package etp

type EngagementMode int

const (
    // Momentum-based: following routine to avoid discomfort of uncertainty
    ComplianceMode EngagementMode = iota
    
    // Curiosity-driven: actively thinking and problem-solving
    EngagementMode
    
    // Resistance: actively avoiding both compliance and engagement
    ResistanceMode
)

type RoutineProfile struct {
    RoutineDependency float64  // 0 (self-directed) to 1 (needs complete structure)
    ThinkingAtrophy   float64  // 0 (thinking is easy) to 1 (thinking is painful)
    FenceVoltage      float64  // Pain level when routine is disrupted
}



type InterventionSequence struct {
    Phase1_EmotionalGrounding []InterventionLever  // Lower baseline anxiety
    Phase2_ConfidenceBuilding []InterventionLever  // Create success experiences
    Phase3_ChallengeIntroduction []InterventionLever  // Gradual independence
}

// Phase 1: Emotional grounding
// Goal: Lower fence voltage before attempting fence crossing
func EmotionalGroundingLevers() []InterventionLever {
    return []InterventionLever{
        {
            Name: "strategic_care_signaling",
            Description: "Show you care about student as person, not rule enforcer",
            Steps: []string{
                "Relax enforcement of arbitrary rules (uniform, etc)",
                "Use student's name frequently",
                "Show interest in their life outside class",
                "Acknowledge when they're having a bad day",
            },
            ETPreduction: []string{"threat_avoidance", "fear", "rejection"},
            BrainStateTarget: "lower_primal_level",  // Reduce baseline stress
        },
        {
            Name: "classroom_climate_shift",
            Description: "Make space feel emotionally safe",
            Steps: []string{
                "Smile, use humor, be visibly relaxed",
                "Normalize struggle ('everyone finds this hard')",
                "Never publicly shame wrong answers",
                "Create 'safe to fail' environment",
            },
            ETPreduction: []string{"fear", "shame", "rejection"},
            BrainStateTarget: "enable_rational_engagement",
        },
    }
}

type StudentBarrier struct {
    Name              string
    Category          BarrierCategory
    ActivatedETPs     []string
    AvoidanceTactics  []string
    EffectiveLevers   []InterventionLever
    UnderlyingCause   string  // Why this barrier exists
}

type BarrierCategory int

const (
    // Acute barriers: specific fears/anxieties about current task
    AcuteBarrier BarrierCategory = iota
    
    // Chronic barriers: long-term learned patterns
    ChronicBarrier
    
    // Structural barriers: system-created obstacles (NEW)
    StructuralBarrier
)

// Your routine trap is a STRUCTURAL barrier
var RoutineDependencyBarrier = StudentBarrier{
    Name:     "routine_dependency_trap",
    Category: StructuralBarrier,
    ActivatedETPs: []string{
        "comfort",      // Routine feels safe
        "belonging",    // Everyone else following routine
        "fear",         // When routine disrupted
        "frustration",  // When thinking required
    },
    UnderlyingCause: "Months/years of compliance training → thinking atrophy → independent thought triggers pain",
    AvoidanceTactics: []string{
        "wait_for_instructions",  // Won't start without being told exactly what to do
        "copy_others",            // Follow peers rather than think independently
        "ask_unnecessary_questions", // "What do I do?" when instructions are clear
        "learned_helplessness",   // "I can't do it" before attempting
    },
    EffectiveLevers: []InterventionLever{
        {Name: "micro_autonomy_injection", /* ... */},
        {Name: "thinking_desensitization", /* ... */},
    },
}



// Bubble: Student's self-created control zone within imposed structure
type Bubble struct {
    Type            BubbleType
    Description     string
    ControlValue    float64    // How much autonomy this provides (0-1)
    SocialValue     float64    // Peer connection value (0-1)
    DistractionRisk float64    // Risk of getting carried away (0-1)
}

type BubbleType string

const (
    SocialBubble    BubbleType = "social"      // Sitting with friends
    SensoryBubble   BubbleType = "sensory"     // Music, fidget tools
    AutonomyBubble  BubbleType = "autonomy"    // Choice of task order, method
    StatusBubble    BubbleType = "status"      // Advanced work, "secret knowledge"
)

// Common student bubbles
var CommonBubbles = []Bubble{
    {
        Type:            SocialBubble,
        Description:     "Sitting with friends/peer group",
        ControlValue:    0.6,
        SocialValue:     0.9,
        DistractionRisk: 0.7,
    },
    {
        Type:            SensoryBubble,
        Description:     "Listening to music while working",
        ControlValue:    0.7,
        SocialValue:     0.2,
        DistractionRisk: 0.4,
    },
    {
        Type:            AutonomyBubble,
        Description:     "Choice of task order or approach method",
        ControlValue:    0.8,
        SocialValue:     0.1,
        DistractionRisk: 0.2,
    },
    {
        Type:            StatusBubble,
        Description:     "Access to 'advanced' or 'grown-up' content",
        ControlValue:    0.5,
        SocialValue:     0.6, // "I get to tell my parents about DNA"
        DistractionRisk: 0.1,
    },
}



type TrustBargain struct {
    BubbleGranted       Bubble
    ExpectedBehavior    string
    RiskMitigation      []string
    EmotionalIntelConvo string  // EQ conversation if they get carried away
}

// Your core lever: Give back bubbles as trust, create forward momentum
var StrategicTrustBargains = []InterventionLever{
    {
        Name:        "social_bubble_grant",
        Description: "Allow friend groups if work gets done",
        Steps: []string{
            "Explicitly frame as trust: 'I'm taking heat to relax this rule'",
            "Set clear boundary: 'Work must get done'",
            "Monitor for critical mass (enough students working)",
            "Use plate-spinning approach (keep momentum, avoid confrontation)",
        },
        Prerequisites: []string{
            "emotional_grounding_complete",  // Must feel safe first
            "some_compliance_momentum",      // Need a few students working
        },
        SuccessIndicators: []string{
            "ripple_effect_visible",  // Working students influence others
            "reduced_active_resistance",
            "increased_task_attempts",
        },
        FailureMode: "students_get_carried_away",
        FailureResponse: "emotional_intelligence_conversation",
        ETPreduction: []string{
            "control_need",     // Addresses need for autonomy
            "belonging",        // Satisfies social need
            "resentment",       // Reduces feeling of being controlled
        },
        BrainStateTarget: "reduce_emotional_resistance",
    },
    {
        Name:        "status_bubble_grant",
        Description: "Give access to 'advanced' content as privilege",
        Steps: []string{
            "Frame as 'secret knowledge' or 'year above' material",
            "Tie to coherent narrative (not isolated facts)",
            "Create 'big reveal' moments (tie new info to familiar things)",
            "Explicit recognition: 'That's A-level stuff, well done'",
        },
        Prerequisites: []string{
            "basic_confidence_established",
            "some_recent_success",
        },
        SuccessIndicators: []string{
            "students_share_with_parents",  // "I learned about DNA today!"
            "increased_engagement_with_advanced_topics",
            "penny_drop_moments_visible",   // Intellectual fairground ride
        },
        ETPreduction: []string{
            "achievement",
            "status",
            "curiosity",
            "excitement",  // The reveal moment
        },
        BrainStateTarget: "trigger_reward_seeking",
    },
}

// Emotional Intelligence Conversation Script
// Used when students get carried away with granted bubbles
type EQConversation struct {
    Trigger   string
    Framework string
    Goal      string
}

var TrustViolationEQScript = EQConversation{
    Trigger: "Student gets carried away with granted bubble (too loud, off-task)",
    Framework: `"You need to understand why people shut you down. 
    They are taking a risk and showing you trust. 
    When you abuse that trust, they have to protect themselves by removing privileges.
    This isn't about control—it's about making the space work for everyone."`,
    Goal: "Build meta-awareness of social contracts, not just compliance through fear",
}



// Your observation: Not about individual motivation, but GROUP DYNAMICS
type ClassroomMomentum struct {
    WorkingStudents      int
    ActiveResisters      int
    ComplianceMode       int
    CriticalMassReached  bool
    PlateSpinningActive  bool
}

// Critical mass threshold
func (cm *ClassroomMomentum) CheckCriticalMass(totalStudents int) bool {
    workingRatio := float64(cm.WorkingStudents) / float64(totalStudents)
    
    // Your observation: Need a few working to create ripple effect
    // Then just maintain momentum
    if workingRatio > 0.4 {  // 40% working creates ripple
        cm.CriticalMassReached = true
        return true
    }
    return false
}

// Plate spinning strategy
type PlateSpinningAction string

const (
    FixIssue          PlateSpinningAction = "fix_issue"           // Help stuck student
    NudgeForward      PlateSpinningAction = "nudge_forward"       // Keep momentum
    DistractResister  PlateSpinningAction = "distract_resister"   // Engage to prevent disruption
    AvoidConfrontation PlateSpinningAction = "avoid_confrontation" // Don't trigger resistance
)

// Your behavior management formula:
// Encourage good + Dissuade bad + Avoid confrontation
var BehaviorManagementPrinciples = []string{
    "Encourage the good things",
    "Dissuade the bad ones", 
    "Avoid confrontation as much as possible",
}



// Your controversial approach: Extend knowledge beyond comfort zone
// Traditional: Spoon-feed to avoid discomfort
// Yours: Stretch → Anxiety → Big Reveal → Penny Drop (intellectual fairground ride)

type KnowledgeExtensionStrategy struct {
    StartPoint          string   // What they know
    ExtensionSteps      []string // How far you stretch
    BigReveal           string   // Connection to familiar thing
    ReassurancePattern  string   // How you manage anxiety
}

var HolisticKnowledgeApproach = InterventionLever{
    Name: "holistic_knowledge_extension",
    Description: "Extend beyond comfort zone with coherent narrative",
    Steps: []string{
        "Start with familiar concept",
        "Extend to 'advanced' material (A-level, university)",
        "Frame as coherent narrative (not isolated facts)",
        "Allow anxiety to build (they're stretching)",
        "Big reveal: Tie to something they see every day",
        "Reassure: 'Not expected to get this all at once'",
        "Repeat over time to build familiarity",
    },
    Benefits: []string{
        "Builds resilience (comfort with not-knowing)",
        "Enables innovation (making conceptual jumps)",
        "Creates 'penny drop' moments (intellectual reward)",
        "Develops adult-level thinking patterns",
    },
    TraditionalObjection: "Students get uncomfortable/confused",
    YourCounter: "Discomfort is WHERE learning happens. Spoon-feeding harms long-term.",
    ETPreduction: []string{
        "curiosity",
        "challenge",
        "excitement",     // The reveal
        "achievement",    // "I understood A-level stuff!"
    },
}

// Model students (ready-made)
type ModelStudentProfile struct {
    HomeEnvironmentSupport bool
    BaselineMotivation     float64
    InterventionNeeded     string
}

var ModelStudentIntervention = InterventionLever{
    Name: "ego_polishing_advanced_feeding",
    Description: "For ready-made students: Polish ego, feed advanced content",
    Steps: []string{
        "Recognize existing motivation",
        "Don't shield from advanced material",
        "Frame as 'grown-up treatment'",
        "Provide intellectual challenge",
        "Let them show off knowledge to parents",
    },
    ETPreduction: []string{"status", "achievement", "mastery"},
}


// Your hammer-home concept: Mistakes are WHERE we learn
type FailurePhilosophy struct {
    TraditionalView string
    YourReframe     string
    Implementation  []string
}

var FailureAsLearning = FailurePhilosophy{
    TraditionalView: "Avoid mistakes, get answers right, pat yourself on back",
    YourReframe:     "Making mistakes is where we learn. Find what you DON'T know.",
    Implementation: []string{
        "Normalize struggle: 'Everyone finds this hard'",
        "Celebrate attempts, not just success",
        "Frame testing as discovery: 'Let's find what you don't know yet'",
        "Safe-to-fail environment essential",
        "Mistakes are DATA, not judgments",
    },
}

// This ties to your holistic approach:
// Extend knowledge → Create anxiety → Inevitable mistakes → Learning happens
// vs. Spoon-feed → No discomfort → No mistakes → No real learning




type InterventionEngine struct {
    // Phase 1: Emotional grounding (from before)
    EmotionalGrounding []InterventionLever
    
    // Phase 2: Confidence building via micro-successes (from before)
    ConfidenceBuilding []InterventionLever
    
    // Phase 2.5: Trust bargains & bubble granting (NEW)
    TrustBargains []TrustBargain
    
    // Phase 3: Holistic knowledge extension (NEW - replaces simple "challenge")
    KnowledgeExtension KnowledgeExtensionStrategy
    
    // Ongoing: Momentum maintenance
    MomentumTracking ClassroomMomentum
    PlateSpinning    []PlateSpinningAction
}

// Decision logic: When to grant bubbles?
func (ie *InterventionEngine) ShouldGrantBubble(
    ctx *StudentContext,
    classroomState ClassroomMomentum,
) (bool, Bubble) {
    // Prerequisite: Emotional grounding complete
    if ctx.BrainState.EmotionalLevel > 0.6 {
        return false, Bubble{} // Too emotionally loaded
    }
    
    // Check classroom state
    if !classroomState.CriticalMassReached {
        // Not enough students working yet
        // Focus on getting a few more before granting bubbles
        return false, Bubble{}
    }
    
    // Identify which bubble would be most effective
    if ctx.SocialNeed > 0.7 {
        return true, CommonBubbles[0] // Social bubble (friends)
    }
    
    if ctx.AutonomyResistance > 0.6 {
        return true, CommonBubbles[1] // Sensory bubble (music)
    }
    
    if ctx.StatusSeeking > 0.5 {
        return true, CommonBubbles[3] // Status bubble (advanced content)
    }
    
    return false, Bubble{}
}



// Group reward structure: Motivate individuals through collective accountability
type GroupRewardStructure struct {
    GroupID              string
    Members              []string  // Student IDs
    CollectiveTarget     string    // "Everyone completes task X"
    RewardType           BubbleType // What bubble is granted as reward
    IndividualProgress   map[string]float64 // Track each member
    GroupProgressPercent float64
    SocialPressureActive bool
}

// The genius move: Group reward converts peer pressure from NEGATIVE to POSITIVE
// Instead of: "Stop working, you're making us look bad"
// You get: "Come on, we need you to finish so we all get the reward"

type SocialPressureDynamics struct {
    Direction      string  // "positive" or "negative"
    Intensity      float64 // 0-1
    Source         string  // "peers", "teacher", "self"
    TargetBehavior string  // What behavior is being pressured
}

// Traditional classroom: Negative peer pressure dominates
// "Don't try too hard" / "You're such a nerd" / "Teacher's pet"
var TraditionalPeerPressure = SocialPressureDynamics{
    Direction:      "negative",
    Intensity:      0.8,
    Source:         "peers",
    TargetBehavior: "academic_effort",
}

// Your intervention: Flip the pressure direction
var GroupRewardPeerPressure = SocialPressureDynamics{
    Direction:      "positive",
    Intensity:      0.7,
    Source:         "peers",
    TargetBehavior: "task_completion",
}

// Group reward intervention lever
var GroupRewardIntervention = InterventionLever{
    Name:        "collective_target_bubble_grant",
    Description: "Grant bubble to entire group if ALL members reach target",
    Steps: []string{
        "Define clear, achievable collective target",
        "Frame as group challenge: 'Can your table all finish this?'",
        "Make reward meaningful (social bubble, status bubble, etc)",
        "Track progress visibly (everyone can see who's done)",
        "Celebrate when group succeeds collectively",
    },
    Benefits: []string{
        "Converts peer pressure from negative to positive",
        "Hiders can't hide anymore (group sees their progress)",
        "High-achievers become tutors instead of targets",
        "Creates organic accountability without teacher policing",
    },
    AvoidanceTacticAddressed: "group_hiding",
    ETPreduction: []string{
        "belonging",         // Be part of team success
        "social_pressure",   // Peers motivate you
        "achievement",       // Group accomplishment
        "status",           // "Our table finished first"
    },
    RiskMitigation: []string{
        "Ensure targets are achievable (can't punish group for one struggling student)",
        "Allow peer tutoring as legitimate strategy",
        "Monitor for negative pressure (bullying weaker students)",
    },
}

// The hiding avoidance tactic you mentioned
type AvoidanceTactic struct {
    Name            string
    Description     string
    Observable      []string  // How you spot it
    Countermeasure  InterventionLever
}

var GroupHidingTactic = AvoidanceTactic{
    Name:        "group_hiding",
    Description: "Student uses group context to avoid individual accountability",
    Observable: []string{
        "Always sits with higher-performing students",
        "Lets others do the work while appearing engaged",
        "Copies answers without understanding",
        "Contributes minimally but claims group credit",
    },
    Countermeasure: GroupRewardIntervention, // Group reward makes hiding visible
}



// Your critical insight: These patterns developed over YEARS
// Cover teachers get short-term compliance, but pathways persist
type LearnedBehaviorPattern struct {
    PatternName        string
    YearsDeveloping    int      // How long this has been reinforced
    TriggerConditions  []string // What activates this pattern
    DefaultResponse    string   // Automatic behavior
    ChangeResistance   float64  // 0-1: How hard to modify
    RequiredReinforcement int   // How many positive experiences needed to shift
}

var ComplianceWithoutThinking = LearnedBehaviorPattern{
    PatternName:       "routine_compliance",
    YearsDeveloping:   5, // Elementary → Middle school
    TriggerConditions: []string{"classroom_setting", "teacher_authority_present"},
    DefaultResponse:   "follow_routine_without_thinking",
    ChangeResistance:  0.7,
    RequiredReinforcement: 20, // Rough estimate: weeks of consistent experience
}

var TestingBoundaries = LearnedBehaviorPattern{
    PatternName:       "trust_boundary_testing",
    YearsDeveloping:   10, // Lifetime of adult interactions
    TriggerConditions: []string{"new_authority_figure", "relaxed_rules"},
    DefaultResponse:   "push_limits_to_find_real_boundaries",
    ChangeResistance:  0.9, // Very persistent
    RequiredReinforcement: 50, // Many cycles of boundary → consequence
}

// Your observation: Cover teachers get away with it SHORT-TERM
type TemporaryComplianceWindow struct {
    DurationMinutes   int     // How long before testing starts
    NoveltyFactor     float64 // New person = temporary deference
    LimitTestingDelay int     // Minutes before first boundary push
}

var CoverTeacherWindow = TemporaryComplianceWindow{
    DurationMinutes:   30, // One lesson
    NoveltyFactor:     0.8,
    LimitTestingDelay: 15, // Mid-lesson they start testing
}

// Long-term relationship: Different dynamics
var EstablishedTeacherDynamics = TemporaryComplianceWindow{
    DurationMinutes:   0,  // No novelty buffer
    NoveltyFactor:     0.0,
    LimitTestingDelay: 5, // Immediate testing if rules change
}



// Your approach: Pull privileges temporarily, frame as cooling-off period
type PrivilegeWithdrawal struct {
    BubbleRevoked      Bubble
    Duration           string        // "rest_of_lesson", "one_day", "until_discussion"
    Framing            string        // How you explain it
    ReinstatementPath  []string      // What student needs to do to regain
    EmotionalTone      string        // Critical: not punishment, but reset
}

var TrustViolationResponse = PrivilegeWithdrawal{
    BubbleRevoked: CommonBubbles[0], // Social bubble
    Duration:      "rest_of_lesson",
    Framing: `"I know changing behavior patterns is hard—these habits developed over years.
    I'm giving things a chance to cool down, not punishing you.
    We're trying to extend trust, but we need to reset when it's not working.
    Tomorrow we can try again."`,
    ReinstatementPath: []string{
        "Complete work independently for rest of lesson",
        "Have brief check-in conversation at end",
        "Fresh start next lesson with bubble re-granted",
    },
    EmotionalTone: "supportive_boundary_enforcement",
}

// Key distinction your approach makes:
type PrivilegeManagementPhilosophy struct {
    TraditionalApproach string
    YourApproach        string
    CriticalDifference  string
}

var PrivilegePhilosophy = PrivilegeManagementPhilosophy{
    TraditionalApproach: "You broke the rule → You lose privileges permanently → Punishment",
    YourApproach:        "You tested boundaries → We reset temporarily → We try again tomorrow",
    CriticalDifference: `Acknowledges that years-old patterns won't change instantly.
    Frames withdrawal as COOLING OFF, not punishment.
    Explicitly states expectation of gradual change, not immediate perfection.
    Maintains relationship: "I'm still on your side, we're working on this together."`,
}

// Implementation in code
type PrivilegeState int

const (
    Granted PrivilegeState = iota
    Withdrawn
    ProbationaryReinstated
)

type StudentPrivilegeTracker struct {
    StudentID         string
    ActiveBubbles     map[BubbleType]PrivilegeState
    ViolationHistory  []PrivilegeViolation
    SuccessHistory    []PrivilegeSuccess
    CurrentTrajectory string // "improving", "stable", "regressing"
}

type PrivilegeViolation struct {
    Timestamp      time.Time
    BubbleType     BubbleType
    Behavior       string  // What they did
    Response       PrivilegeWithdrawal
    StudentReaction string // How they took it
}

type PrivilegeSuccess struct {
    Timestamp        time.Time
    BubbleType       BubbleType
    DurationMinutes  int    // How long they maintained appropriate use
    Reinforcement    string // How you acknowledged it
}

// Gradual trust rebuilding
func (spt *StudentPrivilegeTracker) CalculateTrustLevel() float64 {
    if len(spt.ViolationHistory) == 0 && len(spt.SuccessHistory) == 0 {
        return 0.5 // Neutral starting point
    }
    
    recentSuccesses := countRecent(spt.SuccessHistory, 7) // Last 7 days
    recentViolations := countRecent(spt.ViolationHistory, 7)
    
    // Weight recent more heavily, but don't ignore long-term pattern
    trustLevel := float64(recentSuccesses) / float64(recentSuccesses + recentViolations + 1)
    
    // Consider trajectory
    if spt.CurrentTrajectory == "improving" {
        trustLevel += 0.2
    }
    
    return math.Min(1.0, math.Max(0.0, trustLevel))
}



type InterventionEngine struct {
    // Previous phases
    EmotionalGrounding []InterventionLever
    ConfidenceBuilding []InterventionLever
    TrustBargains      []TrustBargain
    KnowledgeExtension KnowledgeExtensionStrategy
    
    // NEW: Group dynamics management
    GroupRewardStructures []GroupRewardStructure
    SocialPressureTracking map[string]SocialPressureDynamics
    
    // Ongoing: Momentum & privilege management
    MomentumTracking ClassroomMomentum
    PlateSpinning    []PlateSpinningAction
    PrivilegeStates  map[string]StudentPrivilegeTracker // studentID → tracker
}

// Decision logic: Should we use group reward?
func (ie *InterventionEngine) ShouldApplyGroupReward(
    classroomState ClassroomMomentum,
    avoidanceTactics []string,
) bool {
    // Use group reward if:
    // 1. Some students are hiding in groups
    if contains(avoidanceTactics, "group_hiding") {
        return true
    }
    
    // 2. Negative peer pressure is dominant
    if ie.detectNegativePeerPressure(classroomState) {
        return true
    }
    
    // 3. Critical mass not yet reached (need peer motivation)
    if !classroomState.CriticalMassReached {
        return true
    }
    
    return false
}

// Privilege withdrawal decision
func (ie *InterventionEngine) ShouldWithdrawPrivilege(
    studentID string,
    violation PrivilegeViolation,
) (bool, PrivilegeWithdrawal) {
    tracker := ie.PrivilegeStates[studentID]
    
    // Consider: Is this first violation or pattern?
    if len(tracker.ViolationHistory) == 0 {
        // First time: Warning + explicit boundary reminder
        return false, PrivilegeWithdrawal{}
    }
    
    // Pattern emerging: Temporary withdrawal with cooling-off framing
    if len(tracker.ViolationHistory) < 3 {
        return true, TrustViolationResponse
    }
    
    // Persistent pattern: Longer withdrawal + structured conversation
    return true, PrivilegeWithdrawal{
        Duration: "until_discussion",
        Framing: "We need to talk about why this pattern keeps happening",
    }
}


// Your insight: Targets must be personal, constant upward pressure risks blowout
type TargetCalibration struct {
    StudentID         string
    CurrentCapacity   float64  // Where they are now (0-1)
    TargetLevel       float64  // Where you're pushing them (0-1)
    PressureIntensity float64  // How hard you're pushing (0-1)
    BlowoutRisk       float64  // Risk of overwhelm/shutdown (0-1)
    HonestConvoNeeded bool     // Does this require explicit conversation?
}

// Calculate blowout risk
func (tc *TargetCalibration) CalculateBlowoutRisk() float64 {
    // Gap between current and target
    pressureGap := tc.TargetLevel - tc.CurrentCapacity
    
    // Risk increases non-linearly with pressure
    baseRisk := pressureGap * tc.PressureIntensity
    
    // Struggling students have higher baseline risk
    if tc.CurrentCapacity < 0.3 {
        baseRisk *= 1.5
    }
    
    return math.Min(1.0, baseRisk)
}

// Your honest conversation framing
var UphillPressureConversation = InterventionLever{
    Name:        "honest_upward_pressure_framing",
    Description: "Explicit conversation about pushing them + why it matters",
    Steps: []string{
        "Acknowledge difficulty: 'I know this is tough'",
        "Explain necessity: 'You need to understand how much this progress will help'",
        "Commit to support: 'I will always try to give you space'",
        "Frame as investment: 'But we can't stop pushing—it's too important'",
        "Monitor for blowout constantly",
    },
    WhenToUse: []string{
        "Struggling student needs consistent upward pressure",
        "Risk of overwhelm is high",
        "Student might interpret pressure as harassment",
    },
    ETPreduction: []string{
        "trust",         // "Teacher has my back"
        "purpose",       // "This matters for my future"
        "resilience",    // "I can handle tough things"
    },
}

// The central dichotomy you identified
type TeacherStudentDichotomy struct {
    Problem    string
    TradeOff   string
    YourSolution string
}

var PainAvoidanceDichotomy = TeacherStudentDichotomy{
    Problem: "Staying with struggling student is painful for teacher. If student can make teacher move on (via that pain), they win = do nothing.",
    TradeOff: "Teacher pain vs. student avoidance success",
    YourSolution: "Accept the pain, don't move on. Make it clear you care more about their future than your comfort.",
}



// Your insight: Rewards build momentum but KILL self-motivation
type RewardsParadox struct {
    ShortTermEffect  string
    LongTermDanger   string
    OptimalStrategy  string
}

var RewardsDangerZone = RewardsParadox{
    ShortTermEffect: "Rewards create momentum, get initial engagement",
    LongTermDanger:  "Making it about rewards destroys intrinsic motivation",
    OptimalStrategy: "Use rewards minimally, shift to internal drivers ASAP",
}

// Your limited reward types
type AcceptableReward struct {
    RewardType        string
    WhyItWorks        string
    WhyItsLimited     string
    TransitionPath    string
}

var StayWithGroupReward = AcceptableReward{
    RewardType:    "stay_with_friends",
    WhyItWorks:    "Social connection is genuine need, not artificial reward",
    WhyItsLimited: "Still extrinsic—compliance for social access",
    TransitionPath: "Eventually group work becomes norm, not reward",
}

var NoHomeworkReward = AcceptableReward{
    RewardType:    "no_homework_if_class_work_done",
    WhyItWorks:    "Respects their time, shows you value effort efficiency",
    WhyItsLimited: "Can create resentment if classwork incomplete",
    TransitionPath: "Shift to self-directed practice (intrinsic)",
}

var TestTargetReward = AcceptableReward{
    RewardType:    "hit_test_target_for_benefit",
    WhyItWorks:    "Performance-based, visible progress marker",
    WhyItsLimited: "Still outcome-focused rather than process-focused",
    TransitionPath: "Student starts valuing mastery over target-hitting",
}

// The critical transition: Structured homework → Self-directed practice
type HomeworkTransition struct {
    TraditionalApproach string
    Problem             string
    YourApproach        string
    Implementation      []string
}

var HomeworkToSelfPractice = HomeworkTransition{
    TraditionalApproach: "Assign homework, expect compliance, check completion",
    Problem:             "Token effort under duress, no real learning, builds resentment",
    YourApproach:        "Provide high-quality optional materials, respect their time",
    Implementation: []string{
        "Create range of materials: quick-win to deep-dive",
        "Maximize gains per minute invested",
        "Frame as: 'I respect your time, here's what helps most'",
        "Make practice feel like their choice, not your requirement",
        "Celebrate when they use materials voluntarily",
    },
}

// Respect-for-time principle
var TimeRespectPrinciple = InterventionLever{
    Name:        "time_respect_through_efficiency",
    Description: "Show you value their time by providing high-efficiency materials",
    Steps: []string{
        "Design materials for maximum impact per minute",
        "Provide range of options (5-min quick, 30-min deep)",
        "Explicitly state: 'I know your time is valuable'",
        "Track which materials students actually use",
        "Iterate based on their choices",
    },
    Benefits: []string{
        "Builds trust: you respect their life outside school",
        "Reduces resentment: not wasting their time",
        "Increases voluntary engagement: materials are genuinely useful",
        "Shifts motivation: from compliance to self-improvement",
    },
    ETPreduction: []string{
        "respect",       // "Teacher values my time"
        "autonomy",      // "My choice to engage"
        "competence",    // "Materials actually help me improve"
    },
}

// Mental health override for high-achievers
type HighAchieverManagement struct {
    StudentType        string
    DangerPattern      string
    InterventionNeeded string
}

var PanickedOverworker = HighAchieverManagement{
    StudentType:   "always_working_student",
    DangerPattern: "Panic if not constantly working, burnout trajectory",
    InterventionNeeded: "Push mental health maintenance explicitly",
}

var MentalHealthPush = InterventionLever{
    Name:        "mental_health_override_for_overworkers",
    Description: "Explicitly push back against constant work panic",
    Steps: []string{
        "Identify students in constant work panic",
        "Have direct conversation: 'You need to stop sometimes'",
        "Frame rest as performance enhancer, not laziness",
        "Provide permission to not work: 'I'm telling you to take a break'",
        "Monitor for burnout signs",
    },
    WhenToUse: []string{
        "High-achiever showing signs of overwhelm",
        "Student can't disengage from work",
        "Panic response to any downtime",
    },
}



// Your observation: Hiders require massive teacher time investment
type HiderProfile struct {
    StudentID         string
    AvoidanceTactic   string        // "stay_quiet_dont_work"
    TeacherTimeNeeded float64       // Hours per week required
    ConfidenceBuild   ProgressStage // Where they are in journey
    WorkEthicBuild    ProgressStage
}

type ProgressStage int

const (
    NotStarted ProgressStage = iota
    EarlyStage
    BuildingMomentum
    SelfSustaining
)

var HiderManagementStrategy = InterventionLever{
    Name:        "shoulder_sitting_intensive_support",
    Description: "Massive time investment for silent avoiders",
    Steps: []string{
        "Sit on their shoulder (metaphorically): constant proximity",
        "Start with tiniest tasks: 'Write your name'",
        "Celebrate micro-progress: 'You wrote one sentence'",
        "Build confidence molecule by molecule",
        "Accept this takes months, not weeks",
        "Never give up, never move on to avoid pain",
    },
    Prerequisites: []string{
        "teacher_has_capacity_for_intensive_investment",
        "relationship_building_in_progress",
    },
    TimeInvestment: "Huge—potentially 10+ hours per week per student",
    ExpectedTimeline: "Months to years for meaningful progress",
    SuccessMetrics: []string{
        "Student writes one line without prompting",
        "Student asks a question voluntarily",
        "Student completes small task independently",
    },
    ETPreduction: []string{
        "safety",        // "Teacher won't abandon me"
        "competence",    // "I can do small things"
        "belonging",     // "I'm worth teacher's time"
    },
}

// The pain trade-off you identified
var TeacherPainAcceptance = struct {
    Reality string
    Trap    string
    Solution string
}{
    Reality:  "Staying with struggling student is painful (no visible progress, frustrating)",
    Trap:     "If you move on to avoid pain, student learns avoidance wins",
    Solution: "Accept pain, don't move on. Show you care more than they do (right now).",
}



// Gamification: The vegetable-mixing problem
type GamificationRisk struct {
    Strategy    string
    Problem     string
    StudentHack string
    WhenItWorks string
}

var GamificationParadox = GamificationRisk{
    Strategy:    "Hide learning in game mechanics",
    Problem:     "Competitive students game the system without learning",
    StudentHack: "Like spitting out vegetables mixed in dinner—they find the learning and avoid it",
    WhenItWorks: "When game mechanics require actual learning to progress (can't bypass)",
}

// Your performing arts approach: Script-format material
type ContentDeliveryInnovation struct {
    Name            string
    Context         string
    Method          string
    WhyItWorked     string
    ContentContact  string
}

var ScriptFormatDelivery = ContentDeliveryInnovation{
    Name:    "script_format_content",
    Context: "Performing arts school students",
    Method:  "Write curriculum content in script format",
    WhyItWorked: "Students read through multiple times doing silly voices (fun), got more subject contact than copying while unmotivated",
    ContentContact: "Half dozen read-throughs > one unmotivated copy session",
}

// Skills Map Project: Skill practice + content learning
type SkillsMapApproach struct {
    ProjectName       string
    CoreIdea          string
    MotivationalShift string
    BoredomReduction  string
}

var SkillsMapProject = SkillsMapApproach{
    ProjectName: "ESP Skills Map",
    CoreIdea:    "Students learn content while practicing a skill",
    MotivationalShift: "Feel they're improving at something (skill), not just memorizing facts",
    BoredomReduction:  "Skill challenge reduces boredom from content repetition",
}

// General principle: Change motivational balance
var MotivationalBalanceShift = InterventionLever{
    Name:        "skill_practice_content_wrapper",
    Description: "Wrap content learning in skill development activity",
    Steps: []string{
        "Identify skill student wants to develop",
        "Design activity where skill practice requires content engagement",
        "Frame as skill development, not content learning",
        "Student engages more because motivated by skill improvement",
        "Content learning happens as side effect",
    },
    Examples: []string{
        "Script reading: acting skill + content",
        "Debate: argumentation skill + content",
        "Teaching others: communication skill + content",
        "Video creation: production skill + content",
    },
    Benefits: []string{
        "Changes 'why': from 'learn facts' to 'improve skill'",
        "Reduces boredom: challenge comes from skill, not content",
        "Builds transferable capability",
    },
}



// Your question: What drives up voltage for some students and not others?
type VoltageProfile struct {
    StudentID       string
    BaselineVoltage float64  // Resting fence voltage (0-1)
    Triggers        []VoltageTrigger
    GeneticFactor   float64  // Innate sensitivity (0-1)
    LearnedFactor   float64  // Learned responses (0-1)
}

type VoltageTrigger struct {
    Stimulus        string
    VoltageIncrease float64  // How much it raises voltage
    LearnedOrInnate string   // "learned", "genetic", "both"
}

// Autism parallel: Hypersensitivity
var AutismHypersensitivity = VoltageProfile{
    BaselineVoltage: 0.8,  // High resting voltage
    GeneticFactor:   0.9,  // Primarily innate
    LearnedFactor:   0.3,  // Some learned amplification
}

// Brain stem takeover: When voltage exceeds threshold
type BrainStemTakeover struct {
    VoltageThreshold  float64  // Above this = brain stem control
    Observable        []string
    InterventionWindow string
    PostTakeoverOptions []string
}

var TemperTantrum = BrainStemTakeover{
    VoltageThreshold: 0.9,
    Observable: []string{
        "Rational thought offline",
        "Fight-or-flight engaged",
        "No access to cortical reasoning",
        "Physical aggression possible",
    },
    InterventionWindow: "Before threshold—once past, too late for reasoning",
    PostTakeoverOptions: []string{
        "Remove from situation immediately",
        "Let emotions cool (physiological process)",
        "Rebuild relationship after",
    },
}

// Your physical presence: Unspoken authority frame
type AuthorityFraming struct {
    PhysicalPresence   string
    Credentials        string
    Purpose            string
    RiskIfMisused      string
}

var PhysicalPresenceFrame = AuthorityFraming{
    PhysicalPresence: "6ft 1, 95kg, former British Kickboxing champion",
    Credentials:      "Worked with two Fellows of Royal Society",
    Purpose:          "Create aura that keeps things to level, show allegiance through strength",
    RiskIfMisused:    "Can raise threat perception—must use to show you're ON their side",
}

var AllegianceSignaling = InterventionLever{
    Name:        "strength_as_allegiance",
    Description: "Use authority/strength to signal you're fighting FOR them",
    Steps: []string{
        "Establish credibility: they know you're capable",
        "Use strength to show investment: 'Not putting all this in to watch you throw it away'",
        "Frame as: 'We've come so far, why slip back now?'",
        "Explicit: 'I care more about your future than you do right now'",
        "Neurophysiology backs this: prefrontal cortex not fully online in adolescence",
    },
    WhenToUse: []string{
        "Student at risk of throwing away progress",
        "Need to break through apathy/resistance",
        "Relationship strong enough to handle direct confrontation",
    },
    ETPreduction: []string{
        "safety",        // "Strong person is on my side"
        "value",         // "I'm worth this effort"
        "purpose",       // "My future matters"
    },
}

// Context-dependent frames
type ContextualAuthority struct {
    Context      string
    FrameUsed    string
    Purpose      string
}

var ToughGroupFrame = ContextualAuthority{
    Context:   "Tough/challenging student groups",
    FrameUsed: "Physical presence + battle partnership",
    Purpose:   "Show you're tough enough to be on their side, celebrate wins authentically",
}

var PrivilegedGroupFrame = ContextualAuthority{
    Context:   "Over-privileged students",
    FrameUsed: "Intellectual superiority (Fellows of Royal Society)",
    Purpose:   "Win respect through knowledge authority, show you can help them",
}



// Your best reinstatement strategy: Frame in terms of your own mistakes
var SelfDeprecationReinstatement = InterventionLever{
    Name:        "reinstate_through_shared_fallibility",
    Description: "Frame issues as normal, reference your own mistakes",
    Steps: []string{
        "Approach student after cooling off",
        "Reference your own mistakes: 'I've screwed up like this before'",
        "Normalize the behavior: 'Issues are to be expected'",
        "Remove shame: 'We all do this'",
        "Fresh start: 'Let's try again'",
    },
    Benefits: []string{
        "Reduces shame response",
        "Reduces resentment (you're not above them)",
        "Humanizes teacher",
        "Creates permission to fail and recover",
    },
    ETPreduction: []string{
        "shame",       // Reduced
        "resentment",  // Reduced
        "belonging",   // "We're both human"
        "safety",      // "Mistakes are okay"
    },
}



// Your critical observation: Timeline varies wildly, NOT linear
type ProgressTimeline struct {
    StudentID         string
    TimeInvested      int      // Weeks/months
    BattlesCount      int      // Number of conflicts
    EmotionalToll     string   // "embarrassing, wearing, stressful"
    LongTermOutcome   string   // What happened years later
    StudentAwareness  string   // "They know they need to work"
}

var NonLinearProgressReality = struct {
    Myth    string
    Reality string
    Example string
}{
    Myth:    "Consistent effort → linear progress → steady improvement",
    Reality: "Battles over months, embarrassing setbacks, massive stress, then... they appreciate it years later",
    Example: "Students years later: 'You made me work' (said positively)",
}

// Student awareness vs. behavioral control
type StudentAwarenessGap struct {
    Awareness     string
    BehaviorForce string
    Result        string
}

var KnowButCantAct = StudentAwarenessGap{
    Awareness:     "Students are aware they need to work (social norm)",
    BehaviorForce: "Force pushing them NOT to work is overwhelmed by peers/other factors",
    Result:        "Knowing ≠ doing. Peer addiction-level reaction overrides intent.",
}

// The peer addiction concept
var PeerAddictionForce = struct {
    Description  string
    Strength     string
    YourRole     string
}{
    Description: "Peers are their life—addiction-level need for peer connection/status",
    Strength:    "Overwhelms rational knowledge of needing to work",
    YourRole:    "You're trying to break into that world—competing with addiction-level force",
}

// Celebration in tough contexts
var AuthenticCelebrationPrinciple = InterventionLever{
    Name:        "authentic_celebration_in_battle",
    Description: "Tougher the group, more you can celebrate good things",
    Steps: []string{
        "In challenging groups: small wins = big celebrations",
        "Feels more authentic (they're used to battling)",
        "Shows you're on their side through the struggle",
        "Counts for a lot: 'Teacher is fighting with us'",
    },
    WhenToUse: []string{
        "Tough/resistant student groups",
        "Long-term battles over engagement",
        "When small progress is hard-won",
    },
}



type InterventionEngine struct {
    // Phase-based interventions
    EmotionalGrounding    []InterventionLever
    ConfidenceBuilding    []InterventionLever
    TrustBargains         []TrustBargain
    KnowledgeExtension    KnowledgeExtensionStrategy
    
    // Group dynamics
    GroupFormation        GroupFormationEngine
    GroupRewardStructures []GroupRewardStructure
    SocialPressureTracking map[string]SocialPressureDynamics
    BellwetherTracking    map[string]BellwetherStrategy
    
    // Energy & voltage management
    EnergyManagement      map[string]EnergyManagementStrategy
    VoltageProfiles       map[string]VoltageProfile
    
    // Target & rewards
    TargetCalibration     map[string]TargetCalibration
    RewardsStrategy       RewardsParadox
    
    // Special cases
    HiderManagement       map[string]HiderProfile
    HighAchieverOverride  []string  // Student IDs needing mental health push
    
    // Ongoing systems
    MomentumTracking      ClassroomMomentum
    PlateSpinning         []PlateSpinningAction
    PrivilegeStates       map[string]StudentPrivilegeTracker
    ProgressTimelines     map[string]ProgressTimeline
    
    // Authority framing
    TeacherAuthority      AuthorityFraming
    ContextualFrames      map[string]ContextualAuthority
}

// Master decision function: What intervention now?
func (ie *InterventionEngine) SelectIntervention(
    studentID string,
    context StudentContext,
    classroomState ClassroomMomentum,
) InterventionLever {
    
    // Check 1: Brain stem takeover imminent?
    voltage := ie.VoltageProfiles[studentID]
    if voltage.BaselineVoltage > 0.8 {
        return ie.preventBrainStemTakeover(studentID)
    }
    
    // Check 2: Is this a hider requiring intensive support?
    if hider, exists := ie.HiderManagement[studentID]; exists {
        if hider.ConfidenceBuild < BuildingMomentum {
            return HiderManagementStrategy
        }
    }
    
    // Check 3: High achiever needing mental health override?
    if contains(ie.HighAchieverOverride, studentID) {
        return MentalHealthPush
    }
    
    // Check 4: Blowout risk from target pressure?
    target := ie.TargetCalibration[studentID]
    if target.CalculateBlowoutRisk() > 0.7 {
        return UphillPressureConversation
    }
    
    // Check 5: Group dynamics intervention needed?
    if classroomState.ActiveResisters > 5 {
        return ie.handleGroupDynamics(classroomState)
    }
    
    // Standard progression through phases
    return ie.selectPhaseAppropriateIntervention(context)
}



// Your core insight: Chaos is not the problem—it's the path to engagement
type ChaosParadox struct {
    TraditionalView    string
    YourView           string
    WhatOthersHate     string
    WhatYouBelieve     string
}

var ChaosAsPathway = ChaosParadox{
    TraditionalView:    "Chaos = failure. Suppress into compliance for order.",
    YourView:           "Chaos = energy. Ride it through to engagement on the other side.",
    WhatOthersHate:     "Noisy, chaotic, looks like loss of control",
    WhatYouBelieve:     "There's something valuable on the other side of chaos",
}

// The fundamental split in educational philosophy
type EducationalPhilosophySplit struct {
    TraditionalAuthority   TeachingApproach
    YourApproach          TeachingApproach
    InstitutionalConflict bool
}

type TeachingApproach struct {
    Goal                string
    Method              string
    AppearanceInClass   string
    PowerDynamics       string
    LongTermOutcome     string
}

var TraditionalSuppressionApproach = TeachingApproach{
    Goal:              "Maintain order and control",
    Method:            "Suppress student energy into compliance",
    AppearanceInClass: "Quiet, orderly, 'professional'",
    PowerDynamics:     "Teacher authority maintained, students subdued",
    LongTermOutcome:   "Compliance mode, thinking atrophy, resentment",
}

var ChaosSurfingApproach = TeachingApproach{
    Goal:              "Channel student energy into engagement",
    Method:            "Give students lead, ride the energy",
    AppearanceInClass: "Noisy, chaotic, looks 'out of control'",
    PowerDynamics:     "Shared agency, student-led momentum",
    LongTermOutcome:   "Independent thinking, self-motivation, years-later appreciation",
}

var PhilosophicalConflict = EducationalPhilosophySplit{
    TraditionalAuthority:   TraditionalSuppressionApproach,
    YourApproach:          ChaosSurfingApproach,
    InstitutionalConflict: true,
}



// Why power structures oppose chaos-surfing approach
type InstitutionalResistance struct {
    ResistanceSource  string
    PerceivedThreat   string
    ActualThreat      string
    DefenseReaction   string
}

var PowerStructureResistance = InstitutionalResistance{
    ResistanceSource: "Other staff, especially those in power",
    PerceivedThreat:  "Classroom appears out of control, undermines authority norms",
    ActualThreat:     "Exposes that suppression-compliance is about teacher comfort, not student learning",
    DefenseReaction:  "Criticize as unprofessional, demand traditional order",
}

// The visibility problem
type VisibilityBias struct {
    TraditionalClassroom  string
    YourClassroom         string
    WhatAdminsObserve     string
    WhatTheyMiss          string
}

var AppearanceVsReality = VisibilityBias{
    TraditionalClassroom: "Quiet, orderly, students writing in silence",
    YourClassroom:        "Noisy, chaotic, students talking and moving",
    WhatAdminsObserve:    "Disorder, 'poor classroom management'",
    WhatTheyMiss:         "Engagement, thinking, collaboration, energy channeling",
}

// The measurement problem
type OutcomeMeasurementGap struct {
    ShortTermMetrics   []string
    LongTermOutcomes   []string
    WhatGetsRewarded   string
    WhatActuallyMatters string
}

var MeasurementMisalignment = OutcomeMeasurementGap{
    ShortTermMetrics: []string{
        "Quiet classroom",
        "Compliance to instructions",
        "No behavior incidents",
        "Work completed on time",
    },
    LongTermOutcomes: []string{
        "Independent thinking",
        "Self-motivation",
        "Years-later appreciation ('you made me work')",
        "Genuine engagement",
    },
    WhatGetsRewarded:   "Short-term order and compliance",
    WhatActuallyMatters: "Long-term student development",
}

// Risk profile for teachers using this approach
type TeacherRiskProfile struct {
    ProfessionalRisks    []string
    RequiredAttributes   []string
    ProtectiveFactors    []string
    SurvivalStrategies   []string
}

var ChaosSurfingRisks = TeacherRiskProfile{
    ProfessionalRisks: []string{
        "Admin criticism for 'poor classroom management'",
        "Peer judgment ('unprofessional', 'can't control class')",
        "Parent complaints about noise/chaos",
        "Performance reviews flagging 'disorder'",
        "Career advancement blocked",
    },
    RequiredAttributes: []string{
        "Thick skin (withstand criticism)",
        "Confidence in approach (won't cave to pressure)",
        "Long-term thinking (value years-later outcomes over today's appearance)",
        "Willingness to be unpopular with power structures",
        "Physical/intellectual authority (your kickboxing + Royal Society credentials)",
    },
    ProtectiveFactors: []string{
        "Strong results (students do learn, eventually)",
        "Student testimonials (years later)",
        "Parent support (if you can explain approach)",
        "Alternative credibility (your credentials)",
    },
    SurvivalStrategies: []string{
        "Document long-term outcomes",
        "Build parent relationships (bypass admin)",
        "Accept being 'controversial'",
        "Find ally teachers (you're not alone)",
        "Focus on student outcomes, not admin approval",
    },
}



// What you believe is on the other side of chaos
type ChaosTransformationTheory struct {
    Phase1_Suppression    string
    Phase2_Chaos          string
    Phase3_OtherSide      string
    TransformationProcess string
}

var ThroughChaosToEngagement = ChaosTransformationTheory{
    Phase1_Suppression: "Traditional approach: Suppress energy → Compliance mode → Thinking atrophy",
    Phase2_Chaos:       "Your approach: Release energy → Chaotic period → Riding the energy",
    Phase3_OtherSide:   "Emergence: Self-directed engagement, independent thinking, genuine motivation",
    TransformationProcess: "Chaos is not destination—it's transition phase between suppression and engagement",
}

// Chaos as developmental stage
type ChaosAsGrowthStage struct {
    Stage           string
    Duration        string
    TeacherRole     string
    StudentState    string
    Breakthrough    string
}

var ChaosStageMapping = []ChaosAsGrowthStage{
    {
        Stage:        "Pre-chaos (Suppression)",
        Duration:     "Years of prior schooling",
        TeacherRole:  "Maintain order through authority",
        StudentState: "Compliant but disengaged",
        Breakthrough: "None—stuck in routine dependency",
    },
    {
        Stage:        "Chaos Entry",
        Duration:     "First weeks of new approach",
        TeacherRole:  "Give students lead, tolerate noise",
        StudentState: "Testing boundaries, high energy, disorganized",
        Breakthrough: "Students start to believe they have agency",
    },
    {
        Stage:        "Chaos Peak",
        Duration:     "Weeks to months",
        TeacherRole:  "Ride the energy, channel without suppressing",
        StudentState: "Noisy, chaotic, but starting to self-organize",
        Breakthrough: "Critical mass reaches engagement",
    },
    {
        Stage:        "Other Side (Emergence)",
        Duration:     "Months onward",
        TeacherRole:  "Facilitate, guide, step back",
        StudentState: "Self-directed, collaborative, genuinely engaged",
        Breakthrough: "Students work because they want to, not because forced",
    },
}

// The teacher's emotional journey through chaos
type TeacherChaosJourney struct {
    Emotion          string
    WhenItOccurs     string
    WhatHelps        string
}

var TeacherEmotionalArc = []TeacherChaosJourney{
    {
        Emotion:      "Anxiety (Am I losing control?)",
        WhenItOccurs: "Chaos entry—class is loud, admin is watching",
        WhatHelps:    "Trust the process, remember long-term goal",
    },
    {
        Emotion:      "Doubt (Maybe I should just suppress them)",
        WhenItOccurs: "Chaos peak—weeks of noise, peer criticism",
        WhatHelps:    "Look for micro-signs of engagement, celebrate small wins",
    },
    {
        Emotion:      "Exhaustion (This is so much work)",
        WhenItOccurs: "Throughout—plate spinning, shoulder sitting, battles",
        WhatHelps:    "Accept the pain (care more than they do), don't give up",
    },
    {
        Emotion:      "Vindication (It's working!)",
        WhenItOccurs: "Other side—students self-directing, years-later testimonials",
        WhatHelps:    "Document outcomes, share with other teachers considering this",
    },
}



// Track where class/student is in chaos transformation
type ChaosTransformationTracker struct {
    ClassroomID       string
    CurrentStage      ChaosStage
    TimeInStage       int  // Days
    NoiseLevel        float64  // 0 (silent) to 1 (very loud)
    EngagementSignals []string // Micro-indicators of breakthrough
    AdminPressure     float64  // 0 (none) to 1 (high)
    TeacherResilience float64  // 0 (breaking) to 1 (strong)
}

type ChaosStage int

const (
    Suppression ChaosStage = iota
    ChaosEntry
    ChaosPeak
    Emergence
)

// Decision: Is this chaos productive or destructive?
type ChaosQualityAssessment struct {
    IsProductive    bool
    Indicators      []string
    Recommendation  string
}

func (ctt *ChaosTransformationTracker) AssessChaosQuality() ChaosQualityAssessment {
    // Productive chaos indicators:
    productiveSignals := []string{
        "students_talking_about_work",
        "collaborative_problem_solving",
        "questions_being_asked",
        "multiple_approaches_being_tried",
        "peer_tutoring_emerging",
        "laughter_mixed_with_focus",
    }
    
    // Destructive chaos indicators:
    destructiveSignals := []string{
        "no_work_happening",
        "purely_social_conversation",
        "bullying_or_exclusion",
        "students_completely_off_task",
        "no_forward_momentum",
    }
    
    productiveCount := 0
    for _, signal := range ctt.EngagementSignals {
        if contains(productiveSignals, signal) {
            productiveCount++
        }
    }
    
    if productiveCount >= 3 {
        return ChaosQualityAssessment{
            IsProductive: true,
            Indicators: ctt.EngagementSignals,
            Recommendation: "Continue—chaos is transitional, signs of emergence",
        }
    }
    
    return ChaosQualityAssessment{
        IsProductive: false,
        Indicators: ctt.EngagementSignals,
        Recommendation: "Redirect—this chaos is destructive, not developmental",
    }
}

// Support system for teachers navigating chaos
type TeacherSupportSystem struct {
    AdminRelationship    string  // "supportive", "neutral", "hostile"
    AllyTeachers         []string
    ParentBuyIn          float64  // 0-1
    DocumentedOutcomes   []OutcomeEvidence
    ResilienceStrategies []string
}

type OutcomeEvidence struct {
    StudentID       string
    Timestamp       time.Time
    Evidence        string  // "Student independently started advanced work"
    WitnessedBy     string  // "Parent email", "Admin observation", etc.
}

var ChaosNavigationSupport = InterventionLever{
    Name:        "chaos_navigation_teacher_support",
    Description: "Help teachers survive institutional resistance to chaos approach",
    Steps: []string{
        "Document engagement micro-signals daily",
        "Build parent relationships explaining approach",
        "Find ally teachers for mutual support",
        "Prepare for admin criticism with outcome evidence",
        "Remember: Students appreciate it years later (that's what matters)",
        "Accept being 'controversial' as price of effectiveness",
    },
    Prerequisites: []string{
        "teacher_has_strong_conviction",
        "teacher_has_protective_factors", // Credentials, results, etc.
        "teacher_willing_to_accept_criticism",
    },
    LongTermPayoff: "Years-later student testimonials: 'You made me work'",
}

package main

import (
    "fmt"
    "github.com/mike5tew/humanos/internal/etp"
)

// Dashboard to help teachers navigate chaos phase
type ChaosDashboard struct {
    Tracker       etp.ChaosTransformationTracker
    SupportSystem etp.TeacherSupportSystem
}

func (cd *ChaosDashboard) DailyAssessment() {
    // Assess chaos quality
    quality := cd.Tracker.AssessChaosQuality()
    
    fmt.Printf("Chaos Stage: %v\n", cd.Tracker.CurrentStage)
    fmt.Printf("Time in Stage: %d days\n", cd.Tracker.TimeInStage)
    fmt.Printf("Noise Level: %.2f\n", cd.Tracker.NoiseLevel)
    fmt.Printf("Chaos Quality: %v\n", quality.IsProductive)
    fmt.Printf("Recommendation: %s\n", quality.Recommendation)
    
    // Check teacher resilience
    if cd.Tracker.TeacherResilience < 0.3 {
        fmt.Println("⚠️  WARNING: Teacher resilience low. Consider support:")
        fmt.Println("  - Connect with ally teacher")
        fmt.Println("  - Review documented outcomes")
        fmt.Println("  - Remember: Students will appreciate this years later")
    }
    
    // Admin pressure warning
    if cd.Tracker.AdminPressure > 0.7 {
        fmt.Println("⚠️  HIGH ADMIN PRESSURE. Prepare defense:")
        fmt.Println("  - Document engagement signals")
        fmt.Println("  - Gather parent testimonials")
        fmt.Println("  - Present long-term outcome evidence")
    }
    
    // Breakthrough indicators
    if quality.IsProductive && cd.Tracker.CurrentStage == etp.ChaosPeak {
        fmt.Println("✅ BREAKTHROUGH SIGNS: Approaching 'other side'")
        fmt.Println("   Continue current approach—transformation in progress")
    }
}

func (cd *ChaosDashboard) RecordEngagementSignal(signal string) {
    cd.Tracker.EngagementSignals = append(cd.Tracker.EngagementSignals, signal)
    fmt.Printf("📝 Recorded: %s\n", signal)
}

func (cd *ChaosDashboard) RecordOutcomeEvidence(evidence etp.OutcomeEvidence) {
    cd.SupportSystem.DocumentedOutcomes = append(
        cd.SupportSystem.DocumentedOutcomes, 
        evidence,
    )
    fmt.Printf("📋 Evidence logged: %s\n", evidence.Evidence)
}



type InterventionEngine struct {
    // ... (all previous components)
    
    // NEW: Chaos navigation
    ChaosTracker      ChaosTransformationTracker
    TeacherSupport    TeacherSupportSystem
    InstitutionalRisk InstitutionalResistance
}

// Master decision: Is chaos phase sustainable?
func (ie *InterventionEngine) AssessChaosViability() (bool, string) {
    quality := ie.ChaosTracker.AssessChaosQuality()
    
    // Check 1: Is chaos productive?
    if !quality.IsProductive {
        return false, "Chaos is destructive, not developmental—need to redirect"
    }
    
    // Check 2: Can teacher sustain this?
    if ie.ChaosTracker.TeacherResilience < 0.3 {
        return false, "Teacher burnout risk too high—need support or approach change"
    }
    
    // Check 3: Is institutional pressure overwhelming?
    if ie.ChaosTracker.AdminPressure > 0.8 && 
       len(ie.TeacherSupport.DocumentedOutcomes) < 3 {
        return false, "Admin pressure high, insufficient evidence—need to document outcomes"
    }
    
    // Check 4: Are we seeing breakthrough signs?
    breakthroughSignals := []string{
        "students_self_directing",
        "peer_tutoring_emerging",
        "voluntary_practice",
    }
    
    breakthroughCount := 0
    for _, signal := range ie.ChaosTracker.EngagementSignals {
        if contains(breakthroughSignals, signal) {
            breakthroughCount++
        }
    }
    
    if breakthroughCount >= 2 {
        return true, "Breakthrough imminent—chaos is working, stay the course"
    }
    
    return true, "Chaos phase ongoing—productive but not yet breakthrough"
}



// Your actual success indicators (not traditional metrics)
type SuccessIndicators struct {
    CriticalMassWorking      bool
    TechnicalInsightFromResister bool  // The fighter suddenly cares
    CaringAboutResults       bool
    EnjoymentWhileWorking    bool      // "Good at work + enjoy + being yourself"
    SubjectUnderstanding     bool      // "Start understanding what subject is about"
}

// NOT success: Quiet classroom, perfect compliance, orderly appearance
var FalseSuccessMarkers = []string{
    "silent_classroom",
    "perfect_compliance", 
    "no_chaos_visible",
}

// Actual success: Engagement + understanding + authentic self
var RealSuccessMarkers = []string{
    "critical_mass_working",
    "resistant_student_technical_insight",
    "students_care_about_results",
    "enjoyment_while_working",
    "hated_school_now_enjoying",
    "understanding_subject_purpose",
}

// Skills Map Profile indicators
type SkillsMapProfile struct {
    PositiveScore float64  // Engagement, curiosity, effort
    NegativeScore float64  // Avoidance, resistance, apathy
    NetProgress   float64  // Positive - Negative
}

// Success = positive trend over time, not absolute quiet
func (smp *SkillsMapProfile) IsProgressing() bool {
    return smp.NetProgress > 0 && smp.PositiveScore > smp.NegativeScore
}



// Your distinction: Productive chaos vs. lively students taking over
type ChaosQuality struct {
    Type              ChaosType
    CriticalMassState CriticalMassState
    LivelyDominance   float64  // 0 (balanced) to 1 (taken over)
}

type ChaosType string

const (
    ProductiveChaos   ChaosType = "productive"    // Critical mass working, noise from engagement
    LivelyTakeover    ChaosType = "lively_takeover" // Lively students dominate, work stopped
    DestructiveChaos  ChaosType = "destructive"   // No work, pure social/disruption
)

type CriticalMassState string

const (
    CriticalMassWorking    CriticalMassState = "working"     // Enough students engaged
    CriticalMassNotReached CriticalMassState = "not_reached" // Too few working
    CriticalMassLost       CriticalMassState = "lost"        // Was working, now taken over
)

// Detect when lively students have taken over
func (cq *ChaosQuality) DetectLivelyTakeover(
    workingCount int,
    livelyDominantCount int,
    totalStudents int,
) ChaosType {
    workingRatio := float64(workingCount) / float64(totalStudents)
    livelyRatio := float64(livelyDominantCount) / float64(totalStudents)
    
    // Lively takeover: Lively students dominate, working ratio drops
    if livelyRatio > 0.4 && workingRatio < 0.3 {
        return LivelyTakeover
    }
    
    // Productive chaos: Critical mass working, some noise is fine
    if workingRatio > 0.4 {
        return ProductiveChaos
    }
    
    // Destructive: Nothing happening
    return DestructiveChaos
}


// Individual student coaching context (not classroom management)
type IndividualStudentCoach struct {
    StudentProfile    StudentProfile
    BarrierHistory    []StudentBarrier
    InterventionLog   []InterventionRecord
    ProgressTracker   IndividualProgressTracker
    AIPersonality     AICoachPersonality
}

type AICoachPersonality struct {
    Approach       string  // "supportive", "challenging", "matter-of-fact"
    VoltageReading float64 // How sensitive to student voltage
    PressureLevel  float64 // How much upward pressure to apply
}

// Individual intervention (no group dynamics)
type IndividualIntervention struct {
    Barrier             StudentBarrier
    StudentVoltage      float64
    BrainState          BrainState
    RecommendedAction   string
    Reasoning           string  // "Why this suggestion"
}

// Core individual coaching functions
func (isc *IndividualStudentCoach) AnalyzeBarrier(
    situation string,
) StudentBarrier {
    // Detect barrier from student input
    // No group dynamics, pure individual assessment
}

func (isc *IndividualStudentCoach) RecommendNextStep(
    barrier StudentBarrier,
) IndividualIntervention {
    // What should AI suggest to this specific student right now?
}

func (isc *IndividualStudentCoach) TrackProgress() {
    // Monitor individual trajectory over time
    // Skills Map positive/negative scores
}



// Your concern: "As much opinion as anything else"
type ValidationFramework struct {
    ETPMapping           ETPValidation
    InterventionStrategy InterventionValidation
    ProgressMetrics      MetricsValidation
}

type ETPValidation struct {
    MappedToResearch    []string  // Which psychological theories support this
    EmpiricalEvidence   []string  // What studies validate this
    ExpertReview        []string  // Psychologist feedback
    ClassroomEvidence   []string  // Your 12 years of observations
}

type InterventionValidation struct {
    Strategy            string
    PsychologicalBasis  string  // "Rooted in X theory"
    EvidenceStrength    string  // "Strong", "Moderate", "Opinion-based"
    AlternativeViews    []string
    YourRationale       string
}

// Example: Group reward validation
var GroupRewardValidation = InterventionValidation{
    Strategy: "group_reward_structure",
    PsychologicalBasis: "Social interdependence theory (Deutsch, 1949), Peer influence research",
    EvidenceStrength: "Strong - well-supported in cooperative learning literature",
    AlternativeViews: []string{
        "Some argue individual accountability suffers",
        "Risk of scapegoating weaker students",
    },
    YourRationale: "In my experience, converts negative peer pressure to positive",
}

// Example: Chaos surfing validation
var ChaosSurfingValidation = InterventionValidation{
    Strategy: "chaos_surfing_approach",
    PsychologicalBasis: "Self-determination theory (autonomy support), Constructivist learning",
    EvidenceStrength: "Moderate - supported by autonomy research, but chaotic implementation controversial",
    AlternativeViews: []string{
        "Classroom management literature emphasizes structure",
        "Risk of overwhelming students without support",
        "Institutional norms favor order",
    },
    YourRationale: "Years-later student testimonials validate long-term effectiveness despite short-term appearance",
}

// Opinion vs. Evidence spectrum
type EvidenceLevel string

const (
    StrongEvidence      EvidenceLevel = "strong"       // Research-backed + classroom validated
    ModerateEvidence    EvidenceLevel = "moderate"     // Some research support
    ExperientialEvidence EvidenceLevel = "experiential" // Your observations, needs validation
    OpinionBased        EvidenceLevel = "opinion"      // Educated guess, needs testing
)




// Map your interventions to established psychological research
type ResearchMapping struct {
    YourIntervention    InterventionLever
    RelatedTheories     []string
    SupportingStudies   []string
    Contradictions      []string
    ValidationStatus    EvidenceLevel
}

// We'll build this together—for each of your strategies:
// 1. Find supporting research (I can help search)
// 2. Find contradicting research (need to address)
// 3. Identify gaps (where you're innovating beyond current knowledge)
// 4. Design validation tests (how to prove this works)

package main

import (
    "github.com/mike5tew/humanos/internal/etp"
)

// Individual AI coach MVP
type IndividualCoachApp struct {
    Coach        etp.IndividualStudentCoach
    SkillsMap    etp.SkillsMapProfile
    Validation   etp.ValidationFramework
}

func (app *IndividualCoachApp) ProcessStudentInput(input string) {
    // 1. Analyze barrier from student input
    barrier := app.Coach.AnalyzeBarrier(input)
    
    // 2. Check brain state (primal/emotional/rational)
    brainState := app.Coach.CalculateBrainState(barrier)
    
    // 3. Recommend intervention
    intervention := app.Coach.RecommendNextStep(barrier)
    
    // 4. Update Skills Map profile
    app.SkillsMap.RecordInteraction(input, intervention)
    
    // 5. Show reasoning (transparency)
    fmt.Printf("Detected: %s\n", barrier.Name)
    fmt.Printf("Reasoning: %s\n", intervention.Reasoning)
    fmt.Printf("Suggestion: %s\n", intervention.RecommendedAction)
    
    // 6. Validation status
    validation := app.Validation.GetValidation(intervention)
    fmt.Printf("Evidence level: %s\n", validation.EvidenceStrength)
}


// Individual rewards for AI coaching context
type IndividualReward struct {
    RewardType        string
    Duration          string
    TriggerCondition  string
    TransitionPath    string
    DigitalEquivalent string
}

var FiveMinuteBreakReward = IndividualReward{
    RewardType:       "five_minute_break",
    Duration:         "5 minutes",
    TriggerCondition: "Complete focused work session",
    TransitionPath:   "Eventually student self-regulates breaks",
    DigitalEquivalent: "access_to_games_for_five_minutes",
}

var GameAccessReward = IndividualReward{
    RewardType:       "game_access",
    Duration:         "5 minutes",
    TriggerCondition: "Hit micro-goal, complete challenge",
    TransitionPath:   "Gradually reduce frequency, shift to intrinsic motivation",
    DigitalEquivalent: "unlock_game_content_temporarily",
}

// Implementation for digital coaching
type DigitalRewardSystem struct {
    StudentID           string
    EarnedBreakMinutes  int
    RewardHistory       []RewardEvent
    TransitionProgress  float64 // 0 (reward-dependent) to 1 (intrinsically motivated)
}

type RewardEvent struct {
    Timestamp      time.Time
    RewardType     string
    TriggerReason  string
    StudentResponse string // Did they use it? How did they respond?
}

// Xbox account unlock concept
type GamificationUnlock struct {
    UnlockCode      string
    ValidDuration   time.Duration
    EarnedThrough   string // What student did to earn it
    UsageTracking   bool   // Monitor if/how they use it
}

// Generate temporary unlock for gaming/reward access
func (drs *DigitalRewardSystem) GenerateUnlock(
    earnedThrough string,
) GamificationUnlock {
    return GamificationUnlock{
        UnlockCode:    generateTimeLimitedCode(),
        ValidDuration: 5 * time.Minute,
        EarnedThrough: earnedThrough,
        UsageTracking: true,
    }
}

// Track transition from extrinsic to intrinsic
func (drs *DigitalRewardSystem) AssessMotivationShift() float64 {
    recentRewards := getRecentEvents(drs.RewardHistory, 14) // Last 2 weeks
    
    // Calculate ratio of self-initiated work vs. reward-triggered work
    selfInitiated := 0
    rewardTriggered := 0
    
    for _, event := range recentRewards {
        if strings.Contains(event.TriggerReason, "voluntary") {
            selfInitiated++
        } else {
            rewardTriggered++
        }
    }
    
    if len(recentRewards) == 0 {
        return 0.0
    }
    
    return float64(selfInitiated) / float64(len(recentRewards))
}


// The biggest individual barrier you identified
type ConfrontationalPattern struct {
    Name              string
    Category          BarrierCategory
    UnderlyingCause   string
    Manifestations    []string
    AvailableLevers   []string // Very few
    InterventionPath  InterventionSequence
    VictoryDefinition string   // "Every small goal is a victory"
}

var ConfrontationalShowoffBarrier = StudentBarrier{
    Name:     "confrontational_showoff_pattern",
    Category: ChronicBarrier, // Deeply learned behavior
    ActivatedETPs: []string{
        "status",        // Among peers
        "entertainment", // This is how I fill my day
        "belonging",     // To peer group that values this
        "power",         // Taking teachers to pieces
    },
    UnderlyingCause: "Learned behavior: This is the way to fill your day. Status comes from 'taking teachers to pieces' in front of peers.",
    AvoidanceTactics: []string{
        "straight_up_confrontational",
        "deliberately_obnoxious",
        "blank_you_completely",
        "perform_for_peer_audience",
    },
    EffectiveLevers: []InterventionLever{
        // Very few levers available
        {Name: "micro_goal_celebration"},
        {Name: "remove_peer_audience"},
        {Name: "pile_on_praise"},
        {Name: "non_work_chat_building"},
        {Name: "pattern_awareness_conversation"},
    },
}

// Your intervention strategy for this barrier
type ConfrontationalInterventionStrategy struct {
    Phase1_TinyVictories    []string
    Phase2_AudienceRemoval  []string
    Phase3_RelationshipBuild []string
    Phase4_PatternDiscussion []string
    TimelineExpectation     string
}

var ConfrontationalApproach = ConfrontationalInterventionStrategy{
    Phase1_TinyVictories: []string{
        "Every small goal is a victory",
        "Write one word → celebrate",
        "Make eye contact → acknowledge",
        "Stay in room for lesson → praise",
        "Lower bar to ground level, find ANYTHING to recognize",
    },
    Phase2_AudienceRemoval: []string{
        "Work on people around them",
        "Remove their distractions (peer audience)",
        "Separate from enablers/amplifiers",
        "One-on-one interaction when possible",
    },
    Phase3_RelationshipBuild: []string{
        "Pile on the praise (disproportionate to achievement)",
        "Chat is NOT all about work",
        "Show interest in their life, interests, world",
        "Build human connection, not just teacher-student",
    },
    Phase4_PatternDiscussion: []string{
        "Occasionally chat about patterns of what happens",
        "Explain why this is harming them (not moral judgment)",
        "Discuss how we can try to fix it (collaborative)",
        "Not 'you're bad', but 'this pattern isn't working for you'",
    },
    TimelineExpectation: "Months to years. No quick fixes. Inch by inch progress.",
}

// Individual coaching adaptation
type IndividualConfrontationalCoaching struct {
    RecognizePattern     bool
    AvoidMoralJudgment   bool   // Not "good" or "bad"
    TuneSelfToSituation  bool   // "Tune yourself to the situation"
    ManageOwnEmotions    bool   // "Avoid allowing your own emotions to get in the way"
}

var IndividualCoachingPrinciples = IndividualConfrontationalCoaching{
    RecognizePattern:   true,
    AvoidMoralJudgment: true, // "These are not valid descriptions in this context"
    TuneSelfToSituation: true,
    ManageOwnEmotions:  true, // Critical for AI coach: no emotional reactions
}




// Your critical caveat: None of this is magic bullet or clear cut
type InterventionReality struct {
    Myth            string
    Reality         string
    TeacherSkill    string
    AIEquivalent    string
}

var NonMagicBulletReality = InterventionReality{
    Myth:    "Follow intervention X → Get result Y (predictable, mechanical)",
    Reality: "All of this is trying to tune yourself to the situation while avoiding allowing your own emotions or reactions to get in the way",
    TeacherSkill: "Reading situation, adjusting approach, emotional regulation",
    AIEquivalent: "Adaptive reasoning, context sensitivity, no reactive patterns",
}

// Teacher emotional management (critical skill)
type TeacherEmotionalManagement struct {
    Challenge           string
    RequiredSkill       string
    HowItShowsUp        string
    AIImplication       string
}

var EmotionalNonReactivity = TeacherEmotionalManagement{
    Challenge:     "Student is confrontational, obnoxious, or blanking you",
    RequiredSkill: "Avoid letting your own emotions or reactions get in the way",
    HowItShowsUp:  "Stay calm, don't take it personally, don't match their energy",
    AIImplication: "AI has advantage here—no emotional buttons to push. But must model appropriate human response.",
}

// Tuning to situation
type SituationalTuning struct {
    WhatItMeans         string
    HowTeachersDoIt     []string
    HowAICanDoIt        []string
    WhatItIsNot         string
}

var TuningToSituation = SituationalTuning{
    WhatItMeans: "Read student state, context, history → Adjust approach in real-time",
    HowTeachersDoIt: []string{
        "Read body language, tone, energy level",
        "Consider what happened earlier today",
        "Remember previous interactions",
        "Adjust pressure up/down based on signals",
        "Switch tactics if one isn't working",
    },
    HowAICanDoIt: []string{
        "Analyze student input for emotional markers",
        "Track interaction history and patterns",
        "Adjust language, pressure, approach based on voltage",
        "Recognize when to back off vs. push forward",
        "Multiple intervention options, select based on context",
    },
    WhatItIsNot: "Following a script or predetermined path",
}


type IndividualInterventionEngine struct {
    StudentProfile        StudentProfile
    BarrierHistory        []StudentBarrier
    InterventionLog       []InterventionRecord
    ProgressTracker       IndividualProgressTracker
    RewardSystem          DigitalRewardSystem
    EmotionalManagement   EmotionalNonReactivity
    SituationalTuning     SituationalTuning
}

// Detect confrontational/showoff pattern
func (iie *IndividualInterventionEngine) DetectConfrontationalPattern(
    input string,
) (bool, ConfrontationalPattern) {
    // Indicators:
    confrontationalMarkers := []string{
        "confrontational_language",
        "dismissive_tone",
        "ignoring_prompts",
        "performing_resistance",
    }
    
    markerCount := 0
    for _, marker := range confrontationalMarkers {
        if detectMarker(input, marker) {
            markerCount++
        }
    }
    
    if markerCount >= 2 {
        return true, ConfrontationalShowoffBarrier
    }
    
    return false, ConfrontationalPattern{}
}

// Select intervention for confrontational student
func (iie *IndividualInterventionEngine) InterventionForConfrontational(
    barrier ConfrontationalPattern,
    interactionHistory []InterventionRecord,
) IndividualIntervention {
    
    // Check phase: Where are we in the relationship?
    phase := iie.assessRelationshipPhase(interactionHistory)
    
    switch phase {
    case "early": // No relationship yet
        return IndividualIntervention{
            Barrier: barrier,
            RecommendedAction: "Find ANY small thing to recognize. Lower bar to ground level.",
            Reasoning: "Every small goal is a victory. Need to establish that engagement = positive response.",
        }
        
    case "building": // Some relationship forming
        return IndividualIntervention{
            Barrier: barrier,
            RecommendedAction: "Mix praise with non-work chat. Show interest in them as person.",
            Reasoning: "Building human connection beyond work. Chat is not all about the work.",
        }
        
    case "established": // Trust exists
        return IndividualIntervention{
            Barrier: barrier,
            RecommendedAction: "Occasional pattern awareness conversation. 'Let's talk about what's happening and why it's harming you.'",
            Reasoning: "Relationship strong enough for honest discussion. Frame as collaborative problem-solving, not judgment.",
        }
        
    default:
        return IndividualIntervention{
            RecommendedAction: "Assess relationship phase first.",
        }
    }
}

// Tune to situation in real-time
func (iie *IndividualInterventionEngine) TuneToSituation(
    studentInput string,
    context StudentContext,
) IndividualIntervention {
    
    // Check 1: Emotional voltage
    voltage := iie.assessVoltage(studentInput, context)
    
    if voltage > 0.8 {
        return IndividualIntervention{
            RecommendedAction: "Back off. Lower pressure. No confrontation.",
            Reasoning: "Voltage too high—anything demanding will trigger resistance.",
        }
    }
    
    // Check 2: What happened in recent interactions?
    recentPattern := iie.analyzeRecentPattern()
    
    if recentPattern == "consistent_avoidance" {
        return IndividualIntervention{
            RecommendedAction: "Lower goal to absolute minimum. Find tiniest possible success.",
            Reasoning: "Pattern of avoidance—need to break it with guaranteed micro-win.",
        }
    }
    
    if recentPattern == "recent_success" {
        return IndividualIntervention{
            RecommendedAction: "Build on momentum. Slightly raise challenge.",
            Reasoning: "Had success recently—can push a bit more while it's fresh.",
        }
    }
    
    // Check 3: Time of day, previous context, etc.
    // (Would need more context data for full implementation)
    
    return iie.selectPhaseAppropriateIntervention(context)
}

// Manage AI's "emotional" responses (avoid reactive patterns)
func (iie *IndividualInterventionEngine) AvoidReactiveResponse(
    studentInput string,
) (shouldRespond bool, response string) {
    
    // Detect if student is trying to provoke reaction
    isProvocation := detectProvocation(studentInput)
    
    if isProvocation {
        // Don't take bait—respond neutrally or not at all
        return false, "" // Sometimes not responding is the right move
    }
    
    // Detect if AI would have emotional reaction (if it had emotions)
    // e.g., defensive response to criticism, frustration at repeated avoidance
    wouldBeEmotionalResponse := detectEmotionalTrigger(studentInput)
    
    if wouldBeEmotionalResponse {
        // Override with neutral, strategic response
        return true, generateNonReactiveResponse(studentInput)
    }
    
    return true, generateStandardResponse(studentInput)
}




// Top individual barriers (no group dynamics)
var TopIndividualBarriers = []StudentBarrier{
    {
        Name:     "confrontational_showoff_pattern",
        Category: ChronicBarrier,
        UnderlyingCause: "This is the way to fill your day. Status from taking teachers to pieces.",
        AvoidanceTactics: []string{
            "confrontational",
            "obnoxious",
            "blanking",
            "performing_for_peers",
        },
        EffectiveLevers: []InterventionLever{
            {
                Name: "micro_goal_celebration",
                Description: "Every small goal is a victory",
                Steps: []string{
                    "Find ANY tiny thing to recognize",
                    "Lower bar to ground level",
                    "Pile on praise disproportionately",
                },
            },
            {
                Name: "non_work_relationship",
                Description: "Chat is not all about work",
                Steps: []string{
                    "Show interest in their life",
                    "Talk about their interests",
                    "Build human connection",
                },
            },
            {
                Name: "pattern_awareness_conversation",
                Description: "Occasionally discuss the pattern and harm",
                Steps: []string{
                    "Not moral judgment—just pattern observation",
                    "Explain why it's harming them",
                    "Collaborative: 'How can we try to fix it?'",
                },
            },
        },
    },
    // More barriers would go here based on your experience
    // We need to identify the other top 4-5
}

package main

import (
    "fmt"
    "github.com/mike5tew/humanos/internal/etp"
)

type AIStudentCoach struct {
    Engine       etp.IndividualInterventionEngine
    RewardSystem etp.DigitalRewardSystem
}

func (coach *AIStudentCoach) ProcessStudentMessage(msg string) {
    // 1. Detect confrontational pattern?
    isConfrontational, pattern := coach.Engine.DetectConfrontationalPattern(msg)
    
    if isConfrontational {
        fmt.Println("⚠️  Confrontational pattern detected")
        fmt.Println("   Approach: Micro-goal celebration, no moral judgment")
    }
    
    // 2. Tune to situation
    context := coach.Engine.BuildContext(msg)
    intervention := coach.Engine.TuneToSituation(msg, context)
    
    // 3. Avoid reactive response
    shouldRespond, response := coach.Engine.AvoidReactiveResponse(msg)
    
    if !shouldRespond {
        fmt.Println("   [Strategic non-response—not taking bait]")
        return
    }
    
    // 4. Deliver intervention
    fmt.Printf("💡 %s\n", intervention.RecommendedAction)
    fmt.Printf("📝 Reasoning: %s\n", intervention.Reasoning)
    fmt.Printf("💬 Response: %s\n", response)
    
    // 5. Check if reward earned
    if intervention.MicroGoalAchieved {
        unlock := coach.RewardSystem.GenerateUnlock("completed_micro_goal")
        fmt.Printf("🎮 Earned 5-minute game access!\n")
        fmt.Printf("   Code: %s (valid for 5 minutes)\n", unlock.UnlockCode)
    }
}

func main() {
    coach := &AIStudentCoach{
        Engine:       etp.NewIndividualInterventionEngine(),
        RewardSystem: etp.NewDigitalRewardSystem(),
    }
    
    // Example interaction
    studentMsg := "This is boring. Why do I have to do this?"
    
    coach.ProcessStudentMessage(studentMsg)
    
    // Example output:
    // 💡 Acknowledge feeling, find smallest possible engagement point
    // 📝 Reasoning: Voltage moderate, recent avoidance pattern—lower bar
    // 💬 Response: "I hear you. Let's find one thing you're curious about. Just one."
}

