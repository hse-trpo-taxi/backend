package repositories

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hse-trpo-taxi/backend/models"
)

type OrderRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewOrderRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *OrderRepository {
	return &OrderRepository{db: db, builder: builder}
}

// GetLastOrders returns last orders for given driver limited by `limit`.
func (r *OrderRepository) GetLastOrders(driverId int, limit int) ([]*models.OrderDriverList, error) {
	query, args, err := r.builder.Select("id, first_address, second_address, current").
		From("orders").
		Where(squirrel.Eq{"driver_id": driverId}).
		OrderBy("id DESC").
		Limit(uint64(limit)).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*models.OrderDriverList, 0)
	for rows.Next() {
		it := &models.OrderDriverList{}
		err = rows.Scan(&it.Id, &it.FirstAddress, &it.SecondAddress, &it.Current)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}

	return items, nil
}

// GetCurrentOrder returns current order info for driver (where current = true)
func (r *OrderRepository) GetCurrentOrder(driverId int) (*models.CurrentOrderInfo, error) {
	query, args, err := r.builder.Select("time, passengers, time_close, x, y").
		From("orders").
		Where(squirrel.Eq{"driver_id": driverId, "current": true}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(context.Background(), query, args...)

	var timeVal, passengers, timeClose int
	var x, y float64

	if err := row.Scan(&timeVal, &passengers, &timeClose, &x, &y); err != nil {
		return nil, err
	}

	return &models.CurrentOrderInfo{
		Info: &models.OrderDriverInfo{Time: timeVal, Passengers: passengers, TimeClose: timeClose},
		X:    x,
		Y:    y,
	}, nil
}

// GetOrderList returns client-facing order list with driver short info
func (r *OrderRepository) GetOrderList(skip uint64, limit uint64) ([]*models.OrderClientList, error) {
	// select order fields and driver short fields
	sb := r.builder.Select("o.id, o.first_address, o.second_address, o.score, d.name, d.phone, d.status, d.score_driver, d.passport, d.snils, d.inn").
		From("orders o").
		LeftJoin("drivers d ON o.driver_id = d.id")

	if limit > 0 {
		sb = sb.Limit(limit)
	}
	if skip > 0 {
		sb = sb.Offset(skip)
	}

	query, args, err := sb.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*models.OrderClientList, 0)

	for rows.Next() {
		it := &models.OrderClientList{Driver: &models.DriverShortModel{}}
		// scan into temp vars
		var name, phone, status string
		var scoreDriver float64
		var passport, snils, inn int

		err = rows.Scan(&it.Id, &it.FirstAddress, &it.SecondAddress, &it.Score, &name, &phone, &status, &scoreDriver, &passport, &snils, &inn)
		if err != nil {
			return nil, err
		}

		it.Driver.Name = name
		it.Driver.Phone = phone
		it.Driver.Status = status
		it.Driver.ScoreDriver = scoreDriver
		it.Driver.Passport = passport
		it.Driver.Snils = snils
		it.Driver.Inn = inn

		items = append(items, it)
	}

	return items, nil
}
