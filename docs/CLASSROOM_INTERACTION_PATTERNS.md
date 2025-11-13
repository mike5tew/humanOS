# Classroom Interaction Patterns (Extracted for AI Tutor)

**Source**: Classroom workflow discussion - "idealised tablet scenario"
**Purpose**: Document proven interaction patterns for AI tutor implementation
**Status**: Reference for individual coaching design

## Core Interaction Sequence

### 1. Voltage Reduction Through Familiarity
**Classroom Pattern**:
- Recap previous content first
- Make connections to known material
- "Reduce voltage by making as much as possible familiar"

**AI Tutor Adaptation**:
```go
// Pattern: Start with familiar ground
func (coach *AICoach) StartSession(studentID string) {
    // 1. Acknowledge student with warmth
    greeting := "Appreciate the effort as always"
    
    // 2. Quick recap of last session
    previousContext := coach.GetPreviousSession(studentID)
    
    // 3. Connect to student interests
    interests := coach.GetStudentInterests(studentID)
    
    // 4. Lower voltage before new content
    startEasy := coach.SelectMicroSuccess(studentID)
}
```

### 2. Engagement Hook ("Welcome Ladies and Gentlemen")
**Classroom Pattern**:
- Loud but friendly greeting
- Set positive tone immediately
- Break through the hubbub
- "As many smiles as possible = better lesson"

**AI Tutor Adaptation**:
- Personalized greeting using interests
- Warm acknowledgment of student's presence
- Set positive emotional tone before content
- Reference previous successes

**Implementation**:
```markdown
Examples:
- "Great to see you back! Remember that awesome answer you gave last time about..."
- "Welcome! I know you're into [interest] - we're going to connect that to today's topic"
- "Hey there! Before we dive in, how are you feeling about this subject today?"
```

### 3. Question Killer Game (No-Hands-Up Alternative)
**Classroom Pattern**:
- Grid on board = seating chart
- Everyone must answer by end of lesson
- Can't answer again until everyone's participated
- Guided struggling students toward easy answers
- Wrong answers OK if not obviously random

**Why This Works**:
- Avoids public spotlight terror (no hands up)
- Prevents hiding behind enthusiastic students
- Creates safety through structure
- Builds participation muscle memory

**AI Tutor Adaptation**:
```go
// Pattern: Safe, structured question progression
type QuestionProgression struct {
    StartDifficulty  float64  // 0-1, start impossibly easy
    CurrentStreak    int      // Track success
    LastAnswerQuality string  // "excellent", "good", "struggling"
}

func (qp *QuestionProgression) NextQuestion() Question {
    if qp.LastAnswerQuality == "struggling" {
        // Guide toward easier question (guaranteed success)
        return qp.FindGuaranteedWin()
    }
    
    if qp.CurrentStreak >= 3 {
        // Slightly increase difficulty
        return qp.IncreaseDifficultyGradually()
    }
    
    return qp.MaintainCurrentLevel()
}
```

### 4. Plate Spinning (Teacher Bandwidth Management)
**Classroom Pattern**:
- Keep momentum with whole class
- Provide deep individual support where needed
- Accept that chaos develops when focusing on individual
- "Most effective learning = fixing individual misunderstandings"

**AI Tutor Advantage**:
- No bandwidth limit (AI can focus 100% on individual)
- No plate spinning needed
- Can provide depth without chaos tradeoff

**Implementation Note**:
```markdown
This is where AI tutoring has fundamental advantage:
- Classroom: 1 teacher, 30 students = 3% attention per student
- AI Tutor: 1 AI, 1 student = 100% attention
- Result: Can provide "shoulder sitting" intensity for everyone
```

## ESP Behaviour Lite Patterns

### Spatial Interface (Seating Chart as Control)
**Pattern**: Physical layout maps to digital interface
**Key Insight**: Zero cognitive load - flick while maintaining teaching flow

**AI Tutor Parallel**:
Instead of seating chart → **Concept map navigation**

```typescript
interface ConceptMapInterface {
  // Student sees: Visual map of topic
  // Student interacts: Flick gestures on concepts
  
  flickUp: "I've mastered this concept",
  flickDown: "I'm really struggling here",
  flickLeft: "This doesn't connect for me",
  flickRight: "Making good progress"
}
```

### Four-Dimensional Behavior Capture
**Original Four Behaviors**:
- Up: Positive engagement/Excellent answer
- Down: Needs help/Struggling
- Left: Off-task/Disruption
- Right: Good progress/On track

**AI Tutor Mapping**:
```go
type StudentState struct {
    // Vertical axis: Confidence
    Confidence float64  // 0 (down/struggling) to 1 (up/mastered)
    
    // Horizontal axis: Focus
    Focus float64  // 0 (left/disconnected) to 1 (right/progressing)
    
    // Derived metrics
    VoltageLevel  float64  // High if low confidence + low focus
    InterventionNeeded bool // Trigger when both < 0.4
}
```

### Team Competition → Virtual Teams
**Classroom Pattern**:
- Teams compete for positive behavior points
- Bar chart shows live standings
- Social pressure becomes positive force

**AI Tutor Adaptation** (Solo Learning):
```go
type VirtualTeam struct {
    Members []TeamMember
}

type TeamMember struct {
    Role string
    // "Current You" - today's performance
    // "Past You" - progress tracker
    // "Target You" - goal setter
    // "AI Tutor" - coach/cheerleader
}

// Competition mechanics
func (vt *VirtualTeam) CompeteAgainstSelf() {
    // Yesterday vs Today
    // Last week vs This week
    // Personal best vs Current attempt
}
```

## The Learning Journey Progression

### Stage 1: Quick Recap
**Methods**:
- AWS Polly (text-to-speech)
- Spreeder (speed reading tool)
- YouTube video summaries

**Purpose**: Refresh key points, lower voltage

### Stage 2: Semantic Links (Ultra-short)
**Purpose**: Get connections, familiarize with keywords
**Format**: Minimal text, maximum connectivity

Example:
