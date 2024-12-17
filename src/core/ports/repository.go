package ports

import (
	"context"

	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/gofrs/uuid"
)

type Repository interface {
	AddData(ctx context.Context, id uuid.UUID, data domain.Data) error
	GetData(ctx context.Context, id uuid.UUID) (*domain.Data, error)
	RemoveData(ctx context.Context, id uuid.UUID) error
}
