package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gobackhomee/sdk/types"
)

// ErrProjectNotFound is returned when a project is not found
var ErrProjectNotFound = errors.New("project not found")

// ErrUnauthorized is returned when an action is unauthorized
var ErrUnauthorized = errors.New("unauthorized")

// ProjectService handles core business logic for projects
type ProjectService struct {
	repo ProjectRepository
}

// ProjectRepository defines the interface for project data access
// This is a port, but defined here to keep domain self-contained
type ProjectRepository interface {
	Create(ctx context.Context, project *types.Project) error
	Get(ctx context.Context, id string) (*types.Project, error)
	ListByOwner(ctx context.Context, ownerID string) ([]types.Project, error)
	Update(ctx context.Context, project *types.Project) error
}

// NewProjectService creates a new project service
func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

// CreateProject initializes a new project
func (s *ProjectService) CreateProject(ctx context.Context, name, ownerID string) (*types.Project, error) {
	if name == "" {
		return nil, errors.New("project name is required")
	}

	project := &types.Project{
		ID:        generateID(), // In real impl, use UUID
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return project, nil
}

// GetProject retrieves a project by ID checking ownership
func (s *ProjectService) GetProject(ctx context.Context, id, ownerID string) (*types.Project, error) {
	project, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if project.OwnerID != ownerID {
		return nil, ErrUnauthorized
	}

	return project, nil
}

// Helper to generate IDs (placeholder)
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
