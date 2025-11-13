package main
package main

import (
	"fmt"
	"log"

	"github.com/mike5tew/humanos/internal/barriers"
)

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

	fmt.Println("\n🔍 AGE APPROPRIATENESS TEST\n")

	for i, tc := range testCases {
		fmt.Printf("Test %d (Age %d):\n", i+1, tc.Age)
		fmt.Printf("  Original: %s\n", tc.response)

		adjusted := af.AdjustLanguage(tc.response, tc.Age)
		fmt.Printf("  Adjusted: %s\n", adjusted)

		risks := af.CheckOffenseRisk(adjusted, tc.Age)
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
