package repository

import (
	"context"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db postgres.DbPool
}

func NewCSATRepository(db postgres.DbPool) *CSATRepository {
	return &CSATRepository{db}
}

func (cs CSATRepository) GetQuestion(ctx context.Context, userID string) (repModels.Question, error) {
	op := "internal.auth.behavior.GetQuestion"

	var question repModels.Question

	query := `
		SELECT question_id, question
		FROM question
		ORDER BY RANDOM()
		LIMIT 1
	`
	rows, err := cs.db.Query(ctx, query)
	if err != nil {
		return question, errors.Wrap(err, op)
	}

	for rows.Next() {
		if err = rows.Scan(&question.QuestionID, &question.Question); err != nil {
			return question, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "got question=(%v) with questionID=%v", question.Question, question.QuestionID)
	}

	return question, nil
}
