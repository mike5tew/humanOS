package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/joho/godotenv"
    "github.com/michaelstewart/humanos/internal/coach"
    "github.com/michaelstewart/humanos/internal/etp"
)

type Server struct {
    coach *coach.CoachOrchestrator
}

func main() {
    // Load environment variables
    godotenv.Load()

    // Initialize coach orchestrator
    barriersPath := os.Getenv("BARRIERS_PATH")
    if barriersPath == "" {
        barriersPath = "../../shared/schemas/barriers.json"
    }

    traumaPath := os.Getenv("TRAUMA_PATTERNS_PATH")
    if traumaPath == "" {
        traumaPath = "../../shared/schemas/trauma_detection.json"
    }

    coachOrch, err := coach.NewCoachOrchestrator(barriersPath, traumaPath)
    if err != nil {
        log.Fatalf("Failed to initialize coach: %v", err)
    }

    server := &Server{coach: coachOrch}

    // Setup router
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.RequestID)

    // Routes
    r.Post("/api/coach/message", server.handleCoachMessage)
    r.Get("/api/health", server.handleHealth)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting server on port %s", port)
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

    response, err := s.coach.ProcessStudentMessage(req.StudentID, req.Message, req.Context)
    if err != nil {
        log.Printf("Error processing message: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
package api
