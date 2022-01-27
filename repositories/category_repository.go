package repositories

import (
	"context"
	"database/sql"

	"github.com/elwandyTD/blanja-golang-api/models/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	// Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	// FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error)
	// FindBySlug(ctx context.Context, tx *sql.Tx, categorySlug string) domain.Category
	// Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
}
