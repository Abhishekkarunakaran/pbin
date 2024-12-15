package redis

import (
	"context"
	"encoding/json"
	"time"

	log "log/slog"

	"github.com/Abhishekkarunakaran/pbin/src/core/constants"
	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/Abhishekkarunakaran/pbin/src/core/ports"
	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	client *redis.Client
}

// GetData implements ports.Repository.
func (r *repository) GetData(ctx context.Context, id uuid.UUID) (domain.Data, error) {
	panic("unimplemented")
}

// GetPasswordHash implements ports.Repository.
func (r *repository) GetPasswordHash(ctx context.Context, id uuid.UUID) ([]byte, error) {
	panic("unimplemented")
}

func NewRepository(client *redis.Client) ports.Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddData(ctx context.Context, id uuid.UUID, data domain.Data) error {

	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return ErrSerializeData
	}
	if err := r.client.HSet(
		ctx, id.String(),
		serializedData,nil).Err(); err != nil {
		log.Error(err.Error())
		return ErrSaveData
	}

	ttl := time.Duration(constants.Int(constants.Env.RedisTTL))*time.Second

	if err := r.client.Expire(ctx,id.String(),ttl).Err(); err != nil {
		log.Error(err.Error())
		return ErrExpire
	}

	return nil
}
