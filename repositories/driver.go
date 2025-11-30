package repositories

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DriverRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewDriverRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *DriverRepository {
	return &DriverRepository{db: db, builder: builder}
}

func (repository *DriverRepository) GetDrivers() ([]*models.Driver, error) {
	query, args, err := repository.builder.Select("id, name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn, created_at, updated_at").
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
			&newItem.Status,
			&newItem.TimeWork,
			&newItem.X,
			&newItem.Y,
			&newItem.ScoreDriver,
			&newItem.Passport,
			&newItem.Snils,
			&newItem.Inn,
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
	query, args, err := repository.builder.Select("id, name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn, created_at, updated_at").From("drivers").Where(squirrel.Eq{"id": id}).ToSql()
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
		&newItem.Status,
		&newItem.TimeWork,
		&newItem.X,
		&newItem.Y,
		&newItem.ScoreDriver,
		&newItem.Passport,
		&newItem.Snils,
		&newItem.Inn,
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
		Columns("name", "phone", "license_number", "rating", "created_at", "updated_at").
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
		Status:        "atRequest",
		TimeWork:      4,
		X:             0,
		Y:             0,
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
		Suffix("RETURNING id, name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn, created_at, updated_at").
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
		&driver.Status,
		&driver.TimeWork,
		&driver.X,
		&driver.Y,
		&driver.ScoreDriver,
		&driver.Passport,
		&driver.Snils,
		&driver.Inn,
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

func (repository *DriverRepository) GetDriversCoords() ([]*models.DriverCoords, error) {
	query, args, err := repository.builder.Select("id, name, x, y").
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

	items := make([]*models.DriverCoords, 0)

	for rows.Next() {
		newItem := &models.DriverCoords{}
		err = rows.Scan(
			&newItem.ID,
			&newItem.Name,
			&newItem.X,
			&newItem.Y,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
	}

	return items, nil
}

func (repository *DriverRepository) GetDriversStat(model *models.DriverStatRequestModel) ([]*models.DriverStat, error) {
	query, args, err := repository.builder.Select("id, name, status, phone, time_work").
		From("drivers").
		Where(squirrel.Eq{"status": model.Status}).
		Offset(model.Skip).
		Limit(model.Limit).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]*models.DriverStat, 0)

	for rows.Next() {
		newItem := &models.DriverStat{}
		err = rows.Scan(
			&newItem.Id,
			&newItem.Name,
			&newItem.Status,
			&newItem.Phone,
			&newItem.TimeWork,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
	}

	return items, nil

}

func (repository *DriverRepository) GetDriverSchedule(id uint32, req *models.DriverScheduleRequestModel) ([]*models.DriverSchedule, error) {
	var start, end time.Time
	if req != nil {
		start = req.Start
		end = req.End
	}

	if start.IsZero() {
		now := time.Now()
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}
	if end.IsZero() {
		end = start.AddDate(0, 1, 0)
	}

	query, args, err := repository.builder.Select("date").
		From("driver_schedules").
		Where(squirrel.And{
			squirrel.Eq{"driver_id": id},
			squirrel.GtOrEq{"date": start},
			squirrel.LtOrEq{"date": end},
		}).
		OrderBy("date ASC").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*models.DriverSchedule, 0)
	var d time.Time
	for rows.Next() {
		if err := rows.Scan(&d); err != nil {
			return nil, err
		}
		items = append(items, &models.DriverSchedule{Date: d})
	}

	return items, nil
}
