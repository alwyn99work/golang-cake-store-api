package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CakeRepository struct {
}

func NewCakeRepository() CakeRepositoryInterface {
	return &CakeRepository{}
}

func (repository *CakeRepository) Save(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	sql := "INSERT INTO cakes(title,description,rating,image) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, cake.Title, cake.Description, cake.Rating, cake.Image)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cake.Id = int(id)
	return cake
}

func (repository *CakeRepository) Update(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	sql := "UPDATE cakes set title = ?, description = ?, rating = ?, image = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, cake.Title, cake.Description, cake.Rating, cake.Image, cake.Id)
	helper.PanicIfError(err)

	return cake
}

func (repository *CakeRepository) Delete(ctx context.Context, tx *sql.Tx, cakeId int) {
	sql := "DELETE FROM cakes WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, cakeId)
	helper.PanicIfError(err)
}

func (repository *CakeRepository) FindById(ctx context.Context, tx *sql.Tx, cakeId int) (domain.Cake, error) {
	sql := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, cakeId)
	helper.PanicIfError(err)
	defer rows.Close()

	cake := domain.Cake{}
	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.PanicIfError(err)

		return cake, nil
	} else {
		return cake, errors.New("cake not found")
	}
}

func (repository *CakeRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cake {
	sql := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes ORDER BY rating DESC, title ASC"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Cake
	for rows.Next() {
		cake := domain.Cake{}
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.PanicIfError(err)

		categories = append(categories, cake)
	}
	return categories
}
