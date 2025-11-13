# Implementation Status: Concepts → Code

**Complete Framework Reference**: `/docs/reference/ETP_COMPLETE_FRAMEWORK.go` (GitHub confirmed)

## Source Files

- **Individual Coaching Framework**: `/docs/reference/ETP_COMPLETE_FRAMEWORK.go` (GitHub)
- **Classroom Management Framework**: NOT in GitHub yet (concepts in earlier attachments)
- **Production Code**: `/backend/internal/` (selective implementation for AI tutor)

## What's Implemented (AI Tutor MVP)

| Concept | Status | Production Location | Reference Location |
|---------|--------|---------------------|-------------------|
| 5 Core Barriers | ✅ | `/backend/internal/barriers/detector.go` | Lines 200-400 in reference |
| BrainState Types | ✅ | `/backend/internal/etp/types.go` | Lines 50-100 |
| StudentContext | ✅ | `/backend/internal/etp/types.go` | Lines 150-200 |
| Age Appropriateness | 🚧 | `/backend/internal/barriers/age_appropriate.go` | Lines 500-800 |

## Extraction Priority (Next Steps)

### High Priority (Week 1-2)
- [ ] **RoutineProfile** (lines 100-150) → Implement in `types.go`
- [ ] **VoltageProfile** (lines 800-900) → Implement in `types.go`
- [ ] **RelationshipPhase tracking** (lines 1200-1300) → New file `/backend/internal/coach/relationship.go`
- [ ] **ConfrontationalPattern detection** (lines 1500-1800) → Extract to `/backend/internal/barriers/confrontational.go`

### Medium Priority (Week 3-4)
- [ ] **IndividualReward system** (lines 1000-1200) → `/backend/internal/coach/rewards.go`
- [ ] **EmotionalNonReactivity** (lines 1800-2000) → Built into orchestrator
- [ ] **SituationalTuning** (lines 2000-2200) → Orchestrator decision logic

## What's NOT Being Implemented (Not Relevant for 1-on-1 AI Tutor)

| Concept | Status | Reason |
|---------|--------|--------|
| Group Dynamics | 🚫 | No classroom context in AI tutor |
| Critical Mass | 🚫 | Requires multiple students |
| Plate Spinning | 🚫 | Classroom management only |
| Chaos Surfing | 🚫 | Institutional/classroom concept |

**Key Distinction**: The GitHub file contains INDIVIDUAL coaching concepts (perfect for AI tutor). The classroom management concepts (group dynamics, chaos surfing) are in separate discussions/attachments.

## Decision Log

**Why keep group dynamics in reference?**
- Future product: "Teacher Classroom Management Tool"
- Different market than AI tutor (schools vs individuals)
- Represents years of experience that shouldn't be lost

**Why extract only individual concepts?**
- AI tutor = 1-on-1 coaching, no peer dynamics
- 80/20 principle: individual barriers give 80% of value
- Simpler MVP, faster to market

**Complete framework preserved in**: `/docs/reference/ETP_COMPLETE_FRAMEWORK.go`
