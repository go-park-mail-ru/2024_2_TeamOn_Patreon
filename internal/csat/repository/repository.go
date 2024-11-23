package repository

import (
	"context"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
)

type CSATRepository struct {
	db postgres.DbPool
}

func NewCSATRepository(db postgres.DbPool) *CSATRepository {
	return &CSATRepository{db}
}

func (cs CSATRepository) GetQuestion(ctx context.Context, userID string) (repModels.Question, error) {
	// op := "internal.auth.behavior.GetQuestion"

	// query := '
	// SELECT asked_at
	// FROM asked_question
	// WHERE user_id = $1 AND
	// ORDER BY asked_at
	// '
	return repModels.Question{}, nil
}
