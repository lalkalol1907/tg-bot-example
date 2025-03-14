package respository

import (
	"bot-test/internal/features/bot"
	"bot-test/pkg/models"
	"context"
	"github.com/jmoiron/sqlx"
)

const prefix = "bot:message"

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) SaveGood(ctx context.Context, ownerId int64, name string) error {
	_, err := r.db.ExecContext(ctx, AddGoodQuery, name, ownerId)
	return err
}

func (r *Repository) DeleteGood(ctx context.Context, goodId string) error {
	_, err := r.db.ExecContext(ctx, DeleteGoodQuery, goodId)
	return err
}

func (r *Repository) GetGoods(ctx context.Context, ownerId int64) ([]*models.Good, error) {
	result := make([]*models.Good, 0)
	err := r.db.SelectContext(ctx, &result, GetGoodsByOwnerIdQuery, ownerId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewRepository(db *sqlx.DB) bot.IRepository {
	return &Repository{db: db}
}
