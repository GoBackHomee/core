package port

import (
	"context"

	"github.com/gobackhomee/sdk/types"
)

// AIProvider defines the interface for AI operations
type AIProvider interface {
	// GenerateSchema generates a schema from natural language description
	GenerateSchema(ctx context.Context, description string) (string, error)

	// Embed generates vector embeddings for text
	Embed(ctx context.Context, text string) ([]float32, error)
}

// Web3Provider defines the interface for blockchain interactions
type Web3Provider interface {
	// VerifySignature verifies a SIWE signature
	VerifySignature(ctx context.Context, message, signature string) (string, error)

	// GetBalance gets the balance of a wallet
	GetBalance(ctx context.Context, chain, address string) (string, error)
}

// HostingProvider defines the interface for deployment operations
type HostingProvider interface {
	// Deploy deploys a version of a project
	Deploy(ctx context.Context, projectID string, content []byte) (*types.Deployment, error)

	// GetDeployment retrieves a deployment
	GetDeployment(ctx context.Context, deploymentID string) (*types.Deployment, error)
}
