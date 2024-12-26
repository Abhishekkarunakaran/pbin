package ports

import (
	"context"

	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/gofrs/uuid"
)

type Service interface {
	SaveContent(ctx context.Context, payload *domain.Payload) (uuid.UUID, error)
	GetContent(ctx context.Context, dataRequest *domain.DataRequest) (*domain.Content, error)
	IsContentPresent(ctx context.Context, id uuid.UUID) bool
}
