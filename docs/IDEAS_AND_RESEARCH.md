# Ideas & Research Notes

This document contains research ideas, future features, and theoretical explorations for HumanOS.

## Table of Contents
1. [Barrier-Less Savant Hypothesis](#barrier-less-savant-hypothesis)
2. [Neurodiversity & Autism Framework](#neurodiversity--autism-framework)
3. [Options Philosophy](#options-philosophy)
4. [Fiction as AI Training Data](#fiction-as-ai-training-data)
5. [The Knowing-Doing Gap](#the-knowing-doing-gap)
6. [Book Concept: "The Human Operating System"](#book-concept)

// ...existing content from project/ideas.md...

## Implementation Roadmap

### Future Type Definitions

These types are conceptualized but not yet implemented:

```go
// BarrierProfile models which emotional barriers suppress cognitive domains
// Theory: Savant hypothesis - reduced barriers → specialized excellence
type BarrierProfile struct {
	Domain           string   // "math", "music", "visual-spatial", "language"
	BarrierStrength  float64  // 0 (no barrier) to 1 (fully suppressed)
	EmotionalSources []string // Which ETPs create barrier
}

// NeuroProfile models neurological hardware differences (neurodiversity)
type NeuroProfile struct {
	NeurologyType    string             // "neurotypical", "autistic", "adhd"
	SensoryFenceMap  map[string]float64 // Sensory triggers
	ProcessingStyle  string             // "detail-focused", "big-picture"
	InnateStrengths  []string
	DevelopmentNeeds []string
}

// KeyringProfile models available capabilities (options philosophy)
type KeyringProfile struct {
	AcademicKeys   []string
	SocialKeys     []string
	EmotionalKeys  []string
	CreativeKeys   []string
	OptionsBreadth float64  // 0-1: breadth of life paths accessible
}

// PersonalityProfile combines all profile dimensions
type PersonalityProfile struct {
	DominantETPs     []string
	BarrierProfiles  []BarrierProfile
	PowerNeedBalance [2]float64
	NeuroProfile     NeuroProfile
	Keyring          KeyringProfile
}
```

### Research Questions

// ...existing research questions from ideas.md...
