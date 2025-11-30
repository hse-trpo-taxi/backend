package repositories

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/hse-trpo-taxi/backend/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SupportRepository struct {
	db      *pgxpool.Pool
	builder *squirrel.StatementBuilderType
}

func NewSupportRepository(db *pgxpool.Pool, builder *squirrel.StatementBuilderType) *SupportRepository {
	return &SupportRepository{db: db, builder: builder}
}

func (s SupportRepository) GetRecent() ([]*models.SupportModel, error) {
	query, args, err := s.builder.Select("*").
		From("support_requests").
		Limit(5).
		OrderBy("date DESC").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(context.Background(), query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]*models.SupportModel, 0)

	for rows.Next() {
		newItem := &models.SupportModel{}

		// support_requests table has columns: id, name, comment,
		// time_reacting_score, quality_answer_score, date, created_at
		// We only need name..date for the model, so scan id and created_at
		// into throwaway variables.
		var id int
		var createdAt time.Time

		err = rows.Scan(
			&id,
			&newItem.Name,
			&newItem.Comment,
			&newItem.TimeReactingScore,
			&newItem.QualityAnswerScore,
			&newItem.Date,
			&createdAt,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, newItem)
	}

	return items, nil
}
