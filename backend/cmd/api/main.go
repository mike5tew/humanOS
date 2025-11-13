package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/mike5tew/humanos/internal/coach"
	"github.com/mike5tew/humanos/internal/etp"
)

type Server struct {
	orchestrator *coach.Orchestrator
	profiles     map[string]*StudentProfile // In-memory profile storage (for MVP)
}

// StudentProfile represents a student's current state
type StudentProfile struct {
	StudentID       string         `json:"student_id"`
	Age             int            `json:"age"`
	BrainState      etp.BrainState `json:"brain_state"`
	ActiveBarriers  []string       `json:"active_barriers"`
	RewardsEarned   int            `json:"rewards_earned"`
	PlayBreakStage  string         `json:"play_break_stage"`
	LastInteraction string         `json:"last_interaction"`
}

func main() {
	// Load environment
	godotenv.Load()

	// Get paths from environment
	barriersPath := getEnvOrDefault("BARRIERS_PATH", "../../shared/schemas/barriers.json")
	traumaPath := getEnvOrDefault("TRAUMA_PATH", "../../shared/schemas/trauma_detection.json")
	agePath := getEnvOrDefault("AGE_PATH", "../../shared/schemas/age_appropriateness.json")

	// Initialize orchestrator
	orchestrator, err := coach.NewOrchestrator(barriersPath, traumaPath, agePath)
	if err != nil {
		log.Fatalf("Failed to initialize orchestrator: %v", err)
	}

	server := &Server{
		orchestrator: orchestrator,
		profiles:     make(map[string]*StudentProfile),
	}

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Routes
	r.Post("/api/coach/message", server.handleCoachMessage)
	r.Get("/api/student/{studentId}/profile", server.handleGetStudentProfile)
	r.Get("/api/health", server.handleHealth)

	// Start server
	port := getEnvOrDefault("PORT", "8080")
	log.Printf("Starting HumanOS API server on port %s", port)
	log.Printf("Barrier profiles loaded")
	log.Printf("Age appropriateness filtering active")
	log.Printf("Trauma detection active")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

type CoachMessageRequest struct {
	StudentID string             `json:"student_id"`
	Message   string             `json:"message"`
	Context   etp.StudentContext `json:"context"`
}

func (s *Server) handleCoachMessage(w http.ResponseWriter, r *http.Request) {
	var req CoachMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Process through orchestrator
	response, err := s.orchestrator.ProcessMessage(req.StudentID, req.Message, req.Context)
	if err != nil {
		log.Printf("Error processing message: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Update student profile
	s.updateProfile(req.StudentID, req.Context, response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetStudentProfile(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "studentId")

	// Get or create profile
	profile, exists := s.profiles[studentID]
	if !exists {
		// Return default profile
		profile = &StudentProfile{
			StudentID:      studentID,
			Age:            12, // Default age for testing
			PlayBreakStage: "level_1",
			RewardsEarned:  0,
			ActiveBarriers: []string{},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "healthy",
		"service": "humanos-api",
		"version": "0.1.0",
		"features": []string{
			"barrier_detection",
			"age_appropriate_responses",
			"trauma_detection",
			"intervention_selection",
		},
	})
}

func (s *Server) updateProfile(
	studentID string,
	context etp.StudentContext,
	response *coach.CoachResponse,
) {
	profile, exists := s.profiles[studentID]
	if !exists {
		profile = &StudentProfile{
			StudentID:      studentID,
			Age:            context.Age,
			PlayBreakStage: "level_1",
			RewardsEarned:  0,
		}
	}

	// Update brain state
	profile.BrainState = context.BrainState

	// Update barriers
	profile.ActiveBarriers = response.DetectedBarriers

	// Update rewards if earned
	if response.RewardEarned {
		profile.RewardsEarned++
		// TODO: Implement play break stage progression
	}

	// Update last interaction
	profile.LastInteraction = response.Timestamp

	s.profiles[studentID] = profile
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
