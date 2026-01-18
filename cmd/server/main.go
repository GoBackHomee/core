package main

import (
	"context"
	"log"
	"time"

	"github.com/gobackhomee/core/internal/adapter/ai"
	"github.com/gobackhomee/core/internal/domain"
	"github.com/gobackhomee/core/internal/server"
	"github.com/gobackhomee/sdk/config"
	"github.com/gobackhomee/sdk/types"
)

// InMemoryRepo is a temporary mock for bootstrapping
type InMemoryRepo struct{}

func (r *InMemoryRepo) Create(ctx context.Context, project *types.Project) error   { return nil }
func (r *InMemoryRepo) Get(ctx context.Context, id string) (*types.Project, error) { return nil, nil }
func (r *InMemoryRepo) ListByOwner(ctx context.Context, ownerID string) ([]types.Project, error) {
	return nil, nil
}
func (r *InMemoryRepo) Update(ctx context.Context, project *types.Project) error { return nil }

func main() {
	// 1. Load Config (Hardcoded for genesis)
	cfg := config.CoreConfig{
		Server: config.ServerConfig{
			Host:         "0.0.0.0",
			Port:         8080,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		AI: config.AIConfig{
			Provider:     "ollama",
			Endpoint:     "http://localhost:11434",
			DefaultModel: "llama3",
		},
	}

	// 2. Initialize Adapters
	aiProvider := ai.NewOllamaProvider(cfg.AI.Endpoint, cfg.AI.DefaultModel)
	repo := &InMemoryRepo{} // Replace with Postgres adapter later

	// 3. Initialize Domain Services (Inject Adapters)
	projectSvc := domain.NewProjectService(repo)

	// 4. Initialize Server (Inject Domain Services and Ports)
	srv := server.New(cfg.Server, projectSvc, aiProvider)

	// 5. Start Engine
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
