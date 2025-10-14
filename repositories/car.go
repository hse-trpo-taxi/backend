package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
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

func (repository *CarRepository) CreateCar(model *models.CreateCarModel) (*models.Car, error) {
	query, args, err := repository.builder.Insert("cars").
		Columns("driver_id", "brand", "model", "year", "license_plate", "color", "created_at", "updated_at").
		Values(model.DriverID, model.Brand, model.Model, model.Year, model.LicensePlate, model.Color, time.Now(), time.Now()).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()

	if err != nil {
		return nil, err
	}

	var id int
	var createdAt, updatedAt time.Time

	err = repository.db.QueryRow(context.Background(), query, args...).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, err
	}

	return &models.Car{
		ID:           id,
		DriverID:     model.DriverID,
		Brand:        model.Brand,
		Model:        model.Model,
		Year:         model.Year,
		LicensePlate: model.LicensePlate,
		Color:        model.Color,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func (repository *CarRepository) UpdateCar(id uint32, model *models.UpdateCarModel) (*models.Car, error) {
	query, args, err := repository.builder.Update("cars").
		Where(squirrel.Eq{"id": id}).
		SetMap(map[string]interface{}{"driver_id": model.DriverID, "color": model.Color, "updated_at": time.Now()}).
		Suffix("RETURNING id, driver_id, brand, model, year, license_plate, color, created_at, updated_at").
		ToSql()

	if err != nil {
		return nil, err
	}

	car := &models.Car{}

	err = repository.db.QueryRow(context.Background(), query, args...).Scan(
		&car.ID,
		&car.DriverID,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.LicensePlate,
		&car.Color,
		&car.CreatedAt,
		&car.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return car, nil

}

func (repository *CarRepository) DeleteCar(id uint32) error {
	query, args, err := repository.builder.Delete("cars").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(context.Background(), query, args...)

	return err
}
