package repositories

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type ClientRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewClientRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *ClientRepository {
	return &ClientRepository{db: db, builder: builder}
}

func (repository *ClientRepository) GetClients() ([]*models.Client, error) {
	query, args, err := repository.builder.Select("*").
		From("clients").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]*models.Client, 0)

	for rows.Next() {
		newItem := &models.Client{}

		err = rows.Scan(
			&newItem.ID,
			&newItem.Name,
			&newItem.Phone,
			&newItem.Email,
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

func (repository *ClientRepository) GetClientById(id uint32) (*models.Client, error) {
	query, args, err := repository.builder.Select("*").From("clients").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	row := repository.db.QueryRow(context.Background(), query, args...)
	newItem := &models.Client{}
	err = row.Scan(
		&newItem.ID,
		&newItem.Name,
		&newItem.Phone,
		&newItem.Email,
		&newItem.CreatedAt,
		&newItem.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (repository *ClientRepository) CreateClient(model *models.CreateClientModel) (*models.Client, error) {
	query, args, err := repository.builder.Insert("clients").
		Columns("name", "phone", "email", "created_at", "updated_at").
		Values(model.Name, model.Phone, model.Email, time.Now(), time.Now()).
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

	return &models.Client{
		ID:        id,
		Name:      model.Name,
		Phone:     model.Phone,
		Email:     model.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (repository *ClientRepository) UpdateClient(id uint32, model *models.UpdateClientModel) (*models.Client, error) {
	query, args, err := repository.builder.Update("cars").
		Where(squirrel.Eq{"id": id}).
		SetMap(map[string]interface{}{"name": model.Name, "phone": model.Phone, "email": model.Email, "updated_at": time.Now()}).
		Suffix("RETURNING id, name, phone, email, created_at, updated_at").
		ToSql()

	if err != nil {
		return nil, err
	}

	client := &models.Client{}

	err = repository.db.QueryRow(context.Background(), query, args...).Scan(
		&client.ID,
		&client.Name,
		&client.Phone,
		&client.Email,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return client, nil

}

func (repository *ClientRepository) DeleteClient(id uint32) error {
	query, args, err := repository.builder.Delete("clients").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(context.Background(), query, args...)

	return err
}
