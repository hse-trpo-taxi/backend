package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DriverRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewDriverRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *DriverRepository {
	return &DriverRepository{db: db, builder: builder}
}

func (repository *DriverRepository) GetDrivers() ([]*models.Driver, error) {
	query, args, err := repository.builder.Select("*").
		From("drivers").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]*models.Driver, 0)

	for rows.Next() {
		newItem := &models.Driver{}

		err = rows.Scan(
			&newItem.ID,
			&newItem.Name,
			&newItem.Phone,
			&newItem.LicenseNumber,
			&newItem.Rating,
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

func (repository *DriverRepository) GetDriverById(id uint32) (*models.Driver, error) {
	query, args, err := repository.builder.Select("*").From("drivers").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	row := repository.db.QueryRow(context.Background(), query, args...)
	newItem := &models.Driver{}
	err = row.Scan(
		&newItem.ID,
		&newItem.Name,
		&newItem.Phone,
		&newItem.LicenseNumber,
		&newItem.Rating,
		&newItem.CreatedAt,
		&newItem.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (repository *DriverRepository) CreateDriver(model *models.CreateDriverModel) (*models.Driver, error) {
	query, args, err := repository.builder.Insert("drivers").
		Columns("name", "phone", "email", "license_number", "rating", "created_at", "updated_at").
		Values(model.Name, model.Phone, model.LicenseNumber, model.Rating, time.Now(), time.Now()).
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

	return &models.Driver{
		ID:            id,
		Name:          model.Name,
		Phone:         model.Phone,
		LicenseNumber: model.LicenseNumber,
		Rating:        model.Rating,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}

func (repository *DriverRepository) UpdateDriver(id uint32, model *models.UpdateDriverModel) (*models.Driver, error) {
	query, args, err := repository.builder.Update("drivers").
		Where(squirrel.Eq{"id": id}).
		SetMap(map[string]interface{}{
			"name":           model.Name,
			"phone":          model.Phone,
			"license_number": model.LicenseNumber,
			"rating":         model.Rating,
			"updated_at":     time.Now(),
		}).
		Suffix("RETURNING id, name, phone, license_number, rating, created_at, updated_at").
		ToSql()

	if err != nil {
		return nil, err
	}

	driver := &models.Driver{}

	err = repository.db.QueryRow(context.Background(), query, args...).Scan(
		&driver.ID,
		&driver.Name,
		&driver.Phone,
		&driver.LicenseNumber,
		&driver.Rating,
		&driver.CreatedAt,
		&driver.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil

}

func (repository *DriverRepository) DeleteDriver(id uint32) error {
	query, args, err := repository.builder.Delete("drivers").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(context.Background(), query, args...)

	return err
}
