package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gobackhomee/core/internal/domain"
	"github.com/gobackhomee/core/internal/port"
	"github.com/gobackhomee/sdk/config"
)

type Server struct {
	config     config.ServerConfig
	projectSvc *domain.ProjectService
	aiProvider port.AIProvider
	server     *http.Server
}

func New(cfg config.ServerConfig, projSvc *domain.ProjectService, ai port.AIProvider) *Server {
	return &Server{
		config:     cfg,
		projectSvc: projSvc,
		aiProvider: ai,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/api/projects", s.handleProjects)
	mux.HandleFunc("/api/ai/schema", s.handleGenerateSchema)

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}

	log.Printf("Starting server on %s", addr)
	return s.server.ListenAndServe()
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) handleProjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Logic to create project
		// In real impl, parse body, call s.projectSvc.CreateProject
		w.WriteHeader(http.StatusNotImplemented)
	case http.MethodGet:
		// Logic to list projects
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGenerateSchema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	schema, err := s.aiProvider.GenerateSchema(r.Context(), req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"schema": schema}
	json.NewEncoder(w).Encode(resp)
}
