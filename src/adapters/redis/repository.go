package redis

import (
	"context"
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


func NewRepository(client *redis.Client) ports.Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddData(ctx context.Context, id uuid.UUID, data domain.Data) error {

	if err := r.client.HSet(
		ctx, id.String(),
		data).Err(); err != nil {
		log.Error(err.Error())
		return ErrSaveData
	}

	ttl := time.Duration(constants.Int(constants.Env.RedisTTL)) * time.Second

	if err := r.client.Expire(ctx, id.String(), ttl).Err(); err != nil {
		log.Error(err.Error())
		return ErrExpire
	}

	return nil
}

func (r *repository) GetData(ctx context.Context, id uuid.UUID) (*domain.Data, error) {

	var data domain.Data
	err := r.client.HGetAll(ctx, id.String()).Scan(&data)
	if err != nil {
		log.Error(err.Error())
		if err == redis.Nil {
			return nil, ErrValueDoesntExist
		}
		return nil, ErrFetchValue
	}

	return &data, nil
}

func (r *repository) RemoveData(ctx context.Context, id uuid.UUID) error {
	err := r.client.Del(ctx,id.String()).Err()
	if err != nil {
		log.Error(err.Error())
		return ErrRemove
	}
	return nil
}

func (r *repository) IsContentPresent(ctx context.Context, id uuid.UUID) bool {
	val := r.client.Exists(ctx,id.String())
	if val.Val() == 1 {
		return true
	}
	log.Error("content doesn't exists")
	return false
}