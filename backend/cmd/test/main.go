package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mike5tew/humanos/internal/coach"
	"github.com/mike5tew/humanos/internal/etp"
)

func main() {
	fmt.Println("🚀 HumanOS Demo - Testing Interaction Workflow\n")

	// Get the correct paths relative to the project root (where .git is)
	projectRoot := getProjectRoot()

	fmt.Printf("📁 Project root: %s\n", projectRoot)

	barriersPath := filepath.Join(projectRoot, "shared/schemas/barriers.json")
	traumaPath := filepath.Join(projectRoot, "shared/schemas/trauma_detection.json")
	agePath := filepath.Join(projectRoot, "shared/schemas/age_appropriateness.json")

	fmt.Printf("📄 Barriers path: %s\n", barriersPath)
	fmt.Printf("📄 Trauma path: %s\n", traumaPath)
	fmt.Printf("📄 Age path: %s\n\n", agePath)

	// Initialize orchestrator
	orchestrator, err := coach.NewOrchestrator(barriersPath, traumaPath, agePath)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// Test scenarios
	scenarios := []struct {
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

	for i, scenario := range scenarios {
		fmt.Printf("═══════════════════════════════════════════════════════\n")
		fmt.Printf("Test %d: %s\n", i+1, scenario.name)
		fmt.Printf("═══════════════════════════════════════════════════════\n")
		fmt.Printf("Student (age %d): \"%s\"\n\n", scenario.age, scenario.message)

		// Create student context
		context := etp.StudentContext{
			StudentID: fmt.Sprintf("student_%d", i+1),
			Age:       scenario.age,
			BrainState: etp.BrainState{
				PrimalLevel:    0.2,
				EmotionalLevel: 0.5,
				RationalLevel:  0.6,
			},
			ActivatedETPs:      []etp.ETP{},
			RoutineProfile:     etp.RoutineProfile{},
			SocialNeed:         0.5,
			AutonomyResistance: 0.5,
			StatusSeeking:      0.5,
		}

		// Process message
		response, err := orchestrator.ProcessMessage(
			context.StudentID,
			scenario.message,
			context,
		)
		if err != nil {
			fmt.Printf("❌ Error: %v\n\n", err)
			continue
		}

		// Display results
		fmt.Println("📊 ANALYSIS:")
		for _, r := range response.Reasoning {
			fmt.Printf("  %s\n", r)
		}

		fmt.Println("\n💬 AI COACH RESPONSE:")
		fmt.Printf("  \"%s\"\n", response.Message)

		if len(response.DetectedBarriers) > 0 {
			fmt.Println("\n🚧 DETECTED BARRIERS:")
			for _, b := range response.DetectedBarriers {
				fmt.Printf("  • %s\n", b)
			}
		}

		if response.RewardEarned {
			fmt.Println("\n🎮 REWARD EARNED: 5-minute play break unlocked!")
		}

		if response.SafeguardingAlert {
			fmt.Println("\n⚠️  SAFEGUARDING ALERT: Human intervention required")
		}

		fmt.Println()
	}

	fmt.Println("✅ Demo complete! All workflows tested.")
}

// getProjectRoot finds the root directory of the project (where .git is)
func getProjectRoot() string {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// Walk up directories looking for .git
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root without finding .git
			// Fallback: go up 2 levels from backend (backend -> humanOS)
			fallback := filepath.Join(dir, "../..")
			if _, err := os.Stat(filepath.Join(fallback, ".git")); err == nil {
				return fallback
			}
			log.Fatalf("Could not find project root (.git directory)")
		}
		dir = parent
	}
}
