package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CarRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewCarRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *CarRepository {
	return &CarRepository{db: db, builder: builder}
}

func (repository *CarRepository) GetCars() ([]*models.Car, error) {
	query, args, err := repository.builder.Select("*").
		From("cars").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]*models.Car, 0)

	for rows.Next() {
		newItem := &models.Car{}

		err = rows.Scan(
			&newItem.ID,
			&newItem.DriverID,
			&newItem.Brand,
			&newItem.Model,
			&newItem.Year,
			&newItem.LicensePlate,
			&newItem.Color,
			&newItem.CreatedAt,
			&newItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
	}

	return items, nil
}

func (repository *CarRepository) GetCarById(id uint32) (*models.Car, error) {
	query, args, err := repository.builder.Select("*").From("cars").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	row := repository.db.QueryRow(context.Background(), query, args...)
	newItem := &models.Car{}
	err = row.Scan(
		&newItem.ID,
		&newItem.DriverID,
		&newItem.Brand,
		&newItem.Model,
		&newItem.Year,
		&newItem.LicensePlate,
		&newItem.Color,
		&newItem.CreatedAt,
		&newItem.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (repository *CarRepository) CreateCar(model *models.Car) (*models.Car, error) {
	query, args, err := repository.builder.Insert("cars").Columns().Values().ToSql()
	if err != nil {
		return nil, err
	}

	_, err = repository.db.Exec(context.Background(), query, args...)

	if err != nil {
		return nil, err
	}

	newItem := &models.Car{}
}

func (repository *CarRepository) UpdateCar(model *models.Car) (*models.Car, error) {

}

func (repository *CarRepository) DeleteCar(id uint32) error {
	query, args, err := repository.builder.Delete("cars").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(context.Background(), query, args...)

	return err
}
