package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/elwandyTD/blanja-golang-api/helpers"
	"github.com/elwandyTD/blanja-golang-api/models/domain"
	"github.com/google/uuid"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	newUUID := uuid.New()
	parseUUID := newUUID.String()
	slug := helpers.ToSlug(category.Name)
	SQL := "INSERT INTO categories(id, name, image, slug) VALUES(?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, SQL, parseUUID, category.Name, category.Image, slug)
	helpers.PanicIfError(err)

	rows, err := result.RowsAffected()
	helpers.PanicIfError(err)

	if rows > 0 {
		category, err = repository.FindById(ctx, tx, parseUUID)
		helpers.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Failed create new category")
	}

}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId string) (domain.Category, error) {
	SQL := "SELECT id, name, image, slug, create_time, update_time, active FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helpers.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Image, &category.Slug, &category.CreateTime, &category.UpdateTime, &category.Active)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category is not found")
	}
}
