package respository

import (
	"bot-test/internal/features/bot"
	"bot-test/pkg/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) AddGood(ctx context.Context, ownerId int64, name string) error {
	_, err := r.db.ExecContext(ctx, AddGoodQuery, name, ownerId)
	return err
}

func (r *Repository) DeleteGood(ctx context.Context, goodId string) error {
	_, err := r.db.ExecContext(ctx, DeleteGoodQuery, goodId)
	return err
}

func (r *Repository) GetGoodsByOwnerId(ctx context.Context, ownerId int64) ([]*models.Good, error) {
	result := make([]*models.Good, 0)
	err := r.db.SelectContext(ctx, &result, GetGoodsByOwnerIdQuery, ownerId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) AddChat(ctx context.Context, ownerId int64, chatId int64) error {
	_, err := r.db.ExecContext(ctx, AddChatQuery, chatId, ownerId)
	return err
}

func (r *Repository) DeleteChat(ctx context.Context, ChatId int64) error {
	_, err := r.db.ExecContext(ctx, DeleteChatQuery, ChatId)
	return err
}

func (r *Repository) GetChatsByOwnerId(ctx context.Context, ownerId int64) ([]*models.Chat, error) {
	result := make([]*models.Chat, 0)
	err := r.db.SelectContext(ctx, &result, GetChatsByOwnerIdQuery, ownerId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewRepository(db *sqlx.DB) bot.IRepository {
	return &Repository{db: db}
}
