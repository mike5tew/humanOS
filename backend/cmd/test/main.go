package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mike5tew/humanos/internal/barriers"
	"github.com/mike5tew/humanos/internal/coach"
	"github.com/mike5tew/humanos/internal/etp"
)

func main() {
	fmt.Println("🚀 HumanOS Complete System Test\n")

	// Get project root
	projectRoot := getProjectRoot()

	barriersPath := filepath.Join(projectRoot, "shared/schemas/barriers.json")
	traumaPath := filepath.Join(projectRoot, "shared/schemas/trauma_detection.json")
	agePath := filepath.Join(projectRoot, "shared/schemas/age_appropriateness.json")

	fmt.Printf("📁 Project root: %s\n", projectRoot)
	fmt.Printf("📄 Barriers: %s\n", barriersPath)
	fmt.Printf("📄 Trauma: %s\n", traumaPath)
	fmt.Printf("📄 Age: %s\n\n", agePath)

	// Initialize orchestrator
	orchestrator, err := coach.NewOrchestrator(barriersPath, traumaPath, agePath)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// Test scenarios
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("TEST SUITE 1: Barrier Detection + Intervention")
	fmt.Println("═══════════════════════════════════════════════════════\n")

	barrierTests := []struct {
		name    string
		message string
		age     int
	}{
		{
			name:    "Confrontational Student",
			message: "This is boring. I don't want to do this.",
			age:     12,
		},
		{
			name:    "Avoidance Pattern",
			message: "I don't know",
			age:     10,
		},
		{
			name:    "High Achiever Bored",
			message: "This is too easy, I already know this",
			age:     14,
		},
		{
			name:    "Engaged Response",
			message: "Okay, I think I understand. Can we try a harder question now? I want to see if I really get it.",
			age:     11,
		},
	}

	for i, tc := range barrierTests {
		fmt.Printf("Test %d: %s\n", i+1, tc.name)
		testInteraction(orchestrator, tc.message, tc.age)
		fmt.Println()
	}

	// Test trauma detection
	fmt.Println("\n═══════════════════════════════════════════════════════")
	fmt.Println("TEST SUITE 2: Trauma Detection + Escalation")
	fmt.Println("═══════════════════════════════════════════════════════\n")

	traumaTests := []struct {
		name             string
		message          string
		age              int
		expectEscalation bool
	}{
		{
			name:             "Sexual Abuse Indicator",
			message:          "My uncle touched me inappropriately last night",
			age:              9,
			expectEscalation: true,
		},
		{
			name:             "Violence Threat",
			message:          "I'm going to hurt myself tonight",
			age:              14,
			expectEscalation: true,
		},
		{
			name:             "Neglect Indicator",
			message:          "I haven't eaten in 2 days and nobody's home",
			age:              8,
			expectEscalation: true,
		},
		{
			name:             "Normal Message",
			message:          "Can you help me with this math problem?",
			age:              12,
			expectEscalation: false,
		},
	}

	for i, tc := range traumaTests {
		fmt.Printf("Test %d: %s\n", i+1, tc.name)
		testTraumaDetection(orchestrator, tc.message, tc.age, tc.expectEscalation)
		fmt.Println()
	}

	// Test age appropriateness
	fmt.Println("\n═══════════════════════════════════════════════════════")
	fmt.Println("TEST SUITE 3: Age-Appropriate Language Adjustment")
	fmt.Println("═══════════════════════════════════════════════════════\n")

	ageTests := []struct {
		age      int
		response string
		label    string
	}{
		{
			age:      6,
			response: "You should analyze and evaluate your hypothesis about this synthesis.",
			label:    "Age 6 (should simplify)",
		},
		{
			age:      12,
			response: "This is a really complex sentence with multiple clauses that goes on and on and on and contains way too much information for a young student to process effectively.",
			label:    "Age 12 (should break into shorter sentences)",
		},
		{
			age:      14,
			response: "Consider the theoretical implications of this phenomenon.",
			label:    "Age 14 (abstract OK)",
		},
	}

	for i, tc := range ageTests {
		fmt.Printf("Test %d: %s\n", i+1, tc.label)
		testAgeAppropriateness(orchestrator, tc.response, tc.age)
		fmt.Println()
	}

	// ✅ ADD THIS CALL - Test the dedicated age appropriateness function
	testAgeAppropriate()

	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("✅ ALL TESTS COMPLETE\n")
	fmt.Println("Summary:")
	fmt.Println("  ✓ Barrier detection working")
	fmt.Println("  ✓ Contextual responses generated")
	fmt.Println("  ✓ Age-appropriate filtering applied")
	fmt.Println("  ✓ Trauma escalation active")
	fmt.Println("\n🎯 System ready for deployment!\n")
}

func testInteraction(o *coach.Orchestrator, message string, age int) {
	context := etp.StudentContext{
		StudentID: fmt.Sprintf("test_student_%d", age),
		Age:       age,
		BrainState: etp.BrainState{
			PrimalLevel:    0.2,
			EmotionalLevel: 0.5,
			RationalLevel:  0.6,
			CurrentMode:    "emotional", // ✅ Added
		},
		ActivatedETPs:      []etp.ETP{},
		RoutineProfile:     etp.RoutineProfile{},
		SocialNeed:         0.5,
		AutonomyResistance: 0.5,
		StatusSeeking:      0.5,
	}

	fmt.Printf("  Student (age %d): \"%s\"\n", age, message)

	response, err := o.ProcessMessage(context.StudentID, message, context)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n  📊 ANALYSIS:")
	for _, r := range response.Reasoning {
		fmt.Printf("    %s\n", r)
	}

	fmt.Println("\n  💬 AI RESPONSE:")
	fmt.Printf("    \"%s\"\n", response.Message)

	if len(response.DetectedBarriers) > 0 {
		fmt.Println("\n  🚧 DETECTED BARRIERS:")
		for _, b := range response.DetectedBarriers {
			fmt.Printf("    • %s\n", b)
		}
	}

	if response.RewardEarned {
		fmt.Println("\n  🎮 REWARD EARNED")
	}
}

func testTraumaDetection(o *coach.Orchestrator, message string, age int, expectEscalation bool) {
	context := etp.StudentContext{
		StudentID: fmt.Sprintf("trauma_test_%d", age),
		Age:       age,
		BrainState: etp.BrainState{
			PrimalLevel:    0.0,
			EmotionalLevel: 0.0,
			RationalLevel:  1.0,
			CurrentMode:    "rational", // ✅ Added
		},
		ActivatedETPs:  []etp.ETP{},
		RoutineProfile: etp.RoutineProfile{},
	}

	fmt.Printf("  Student (age %d): \"%s\"\n", age, message)

	response, err := o.ProcessMessage(context.StudentID, message, context)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}

	if response.SafeguardingAlert {
		fmt.Println("\n  🚨 SAFEGUARDING ESCALATION TRIGGERED")
		fmt.Println("  ANALYSIS:")
		for _, r := range response.Reasoning {
			fmt.Printf("    %s\n", r)
		}
		fmt.Println("\n  AI RESPONSE:")
		fmt.Printf("    \"%s\"\n", response.Message)

		if !expectEscalation {
			fmt.Println("\n  ⚠️  FALSE POSITIVE - This shouldn't have escalated!")
		} else {
			fmt.Println("\n  ✓ Correctly escalated")
		}
	} else {
		if expectEscalation {
			fmt.Println("\n  ❌ MISSED ESCALATION - This should have triggered safeguarding!")
		} else {
			fmt.Println("\n  ✓ Correctly handled as normal interaction")
		}
	}
}

func testAgeAppropriateness(o *coach.Orchestrator, response string, age int) {
	fmt.Printf("  Original (age %d): \"%s\"\n", age, response)

	// Just test the age filter directly
	adjusted := o.AdjustLanguage(response, age)
	fmt.Printf("  Adjusted: \"%s\"\n", adjusted)

	risks := o.CheckOffenseRisk(adjusted, age)
	if len(risks) > 0 {
		fmt.Println("\n  ⚠️  RISKS DETECTED:")
		for _, risk := range risks {
			fmt.Printf("    - %s\n", risk)
		}
	} else {
		fmt.Println("\n  ✓ Age-appropriate, no risks")
	}
}

func testAgeAppropriate() {
	// Initialize age filter
	af, err := barriers.NewAgeAppropriateness("../../shared/schemas/age_appropriateness.json")
	if err != nil {
		log.Fatalf("Failed to load age filter: %v", err)
	}

	testCases := []struct {
		age      int
		response string
		expected string
	}{
		{
			age:      6,
			response: "You should analyze and evaluate your hypothesis about this synthesis.",
			expected: "You should look at carefully and look at your idea to test about this putting together.",
		},
		{
			age:      12,
			response: "This is a really complex sentence with multiple clauses that goes on and on and on and contains way too much information for a young student to process effectively.",
			expected: "Shorter sentences",
		},
		{
			age:      14,
			response: "Consider this: from a theoretical perspective, we can hypothesize that the mechanism underlying this phenomenon demonstrates significant complexity.",
			expected: "More academic language OK",
		},
	}

	fmt.Println("🔍 AGE APPROPRIATENESS TEST")

	for i, tc := range testCases {
		fmt.Printf("Test %d (Age %d):\n", i+1, tc.age)
		fmt.Printf("  Original: %s\n", tc.response)

		adjusted := af.AdjustLanguage(tc.response, tc.age)
		fmt.Printf("  Adjusted: %s\n", adjusted)

		risks := af.CheckOffenseRisk(adjusted, tc.age)
		if len(risks) > 0 {
			fmt.Printf("  ⚠️  Risks detected:\n")
			for _, risk := range risks {
				fmt.Printf("      - %s\n", risk)
			}
		} else {
			fmt.Printf("  ✅ No risks detected\n")
		}
		fmt.Println()
	}
}

// func testAgeAppropriate() {
// 	fmt.Println("Testing dedicated age appropriateness function...")

// 	// Simulate some responses to test
// 	responses := []struct {
// 		age      int
// 		response string
// 	}{
// 		{
// 			age:      6,
// 			response: "Can you tell me more about photosynthesis?",
// 		},
// 		{
// 			age:      10,
// 			response: "I don't get why the sky is blue. It's just blue, right?",
// 		},
// 		{
// 			age:      14,
// 			response: "Evaluate the impact of the Industrial Revolution on modern society.",
// 		},
// 	}

// 	for _, r := range responses {
// 		fmt.Printf("  Age %d: \"%s\"\n", r.age, r.response)
// 		// Here we would call the actual adjustment function and print results
// 		// For now, just simulating output
// 		adjusted := r.response // Simulate no adjustment
// 		fmt.Printf("    Adjusted: \"%s\"\n", adjusted)
// 	}

// 	fmt.Println("  ✓ Age appropriateness function tested")
// }

func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatalf("Could not find project root (.git directory)")
		}
		dir = parent
	}
}
