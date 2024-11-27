package repository

import (
	"context"
	"database/sql"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type CSATRepository struct {
	db postgres.DbPool
}

func NewCSATRepository(db postgres.DbPool) *CSATRepository {
	return &CSATRepository{db}
}

func (cs *CSATRepository) GetQuestion(ctx context.Context, userID string) (repModels.Question, error) {

	op := "internal.csat.repository.GetQuestion"

	var question repModels.Question

	query := `
		SELECT question_id, question, question_theme_id
		FROM question
		ORDER BY RANDOM()
		LIMIT 1
	`
	rows, err := cs.db.Query(ctx, query)
	if err != nil {
		return question, errors.Wrap(err, op)
	}

	for rows.Next() {
		if err = rows.Scan(&question.QuestionID, &question.Question, &question.QuestionThemeID); err != nil {
			return question, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "got question=(%v) with questionID=%v", question.Question, question.QuestionID)
	}

	return question, nil
}

func (cs *CSATRepository) CreateNewAskedQuestion(ctx context.Context, userID string, question repModels.Question) error {
	op := "internal.csat.repository.CreateNewAskedQuestion"

	query := `
		INSERT INTO asked_question (asked_question_id, user_id, question_id, asked_date, result)
		VALUES ($1, $2, $3, DEFAULT, NULL);
	`

	askedQuestionID := utils.GenerateUUID()

	_, err := cs.db.Exec(ctx, query, askedQuestionID, userID, question.QuestionID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cs *CSATRepository) SaveRating(ctx context.Context, userID, questionID string, rating int) error {
	op := "internal.csat.repository.GetQuestion"

	query := `
		UPDATE asked_question
		SET result = $1 
		WHERE question_id = $2
	`
	_, err := cs.db.Exec(ctx, query, rating, questionID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (cs *CSATRepository) GetTimeLastQuestion(ctx context.Context, userID string) (sql.NullTime, error) {
	op := "internal.csat.repository.GetTimeLastQuestion"

	// Запрос на получение массива
	query := `
		SELECT MAX(asked_date)
		FROM asked_question
		WHERE user_id = $1;
	`

	rows, err := cs.db.Query(ctx, query, userID)
	if err != nil {
		return sql.NullTime{}, errors.Wrap(err, op)
	}

	defer rows.Close()

	var lastTime sql.NullTime

	for rows.Next() {
		if err = rows.Scan(&lastTime.Time); err != nil {
			return sql.NullTime{}, errors.Wrap(err, op)
		}
		lastTime.Valid = true
		logger.StandardDebugF(ctx, op, "Got lastTime='%v' for userID='%v'", lastTime, userID)
	}
	return lastTime, nil
}

// func (cs *CSATRepository) getThemeQuestion(ctx context.Context, questionID string) (string, error) {
// 	op := "internal.csat.repository.getThemeQuestion"

// 	query := `
// 		SELECT qt.theme
// 		FROM question q
// 		JOIN question_theme qt ON q.question_theme_id = qt.question_theme_id
// 		WHERE q.question_id = $1
// 	`

// 	rows, err := cs.db.Query(ctx, query)
// 	if err != nil {
// 		return "", errors.Wrap(err, op)
// 	}

// 	defer rows.Close()

// 	var (
// 		theme string
// 	)

// 	for rows.Next() {
// 		if err = rows.Scan(&theme); err != nil {
// 			return "", errors.Wrap(err, op)
// 		}
// 		logger.StandardDebugF(ctx, op, "Got theme='%v' for questionID='%v'", theme, questionID)
// 	}
// 	return theme, nil
// }
