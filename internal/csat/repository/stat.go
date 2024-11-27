package repository

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

const (
	// возвращает статистику за все время
	// Output: theme, rating
	getAllStat = `
SELECT
    qt.theme AS question_theme,
	TO_CHAR(AVG(aq.result), 'FM999999999.00') AS average_result 
FROM
    asked_question aq
JOIN
    question q ON aq.question_id = q.question_id
JOIN
    question_theme qt ON q.question_theme_id = qt.question_theme_id
WHERE
    aq.result IS NOT NULL -- Игнорируем записи с NULL в поле result
GROUP BY
    qt.theme
`
	// возвращает статистику по дням
	// Input: $1 - days число дней = "CURRENT_DATE - INTERVAL '$1 days'"
	// Output: theme, rating
	getDayStat = `
SELECT
    qt.theme AS question_theme,
	TO_CHAR(AVG(aq.result), 'FM999999999.00') AS average_result 
FROM
    asked_question aq
JOIN
    question q ON aq.question_id = q.question_id
JOIN
    question_theme qt ON q.question_theme_id = qt.question_theme_id
WHERE
    aq.asked_date >= $1  -- Фильтруем по последней неделе
    AND aq.result IS NOT NULL -- Игнорируем записи с NULL в поле result
GROUP BY
    qt.theme
`
)

func (r *CSATRepository) GetAllStat(ctx context.Context) ([]*models.Stat, error) {
	op := "csat.repository.GetAllStat"

	stats := make([]*models.Stat, 0)

	rows, err := r.db.Query(ctx, getAllStat)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		theme  string
		rating string
	)

	for rows.Next() {
		if err = rows.Scan(&theme, &rating); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  question: theme=%v rating=%v",
			theme, rating)
		stats = append(stats, &models.Stat{
			Theme:  theme,
			Rating: rating,
		})

	}

	return stats, nil

}

func (r *CSATRepository) GetStatByDays(ctx context.Context, days int) ([]*models.Stat, error) {
	op := "csat.repository.GetDayStat"

	stats := make([]*models.Stat, 0)

	startDate := time.Now().AddDate(0, 0, -days)
	startDateStr := startDate.Format("2006-01-02") // Преобразуем в строку
	logger.StandardDebugF(ctx, op, "interval=%v", startDateStr)

	rows, err := r.db.Query(ctx, getDayStat, startDateStr)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		theme  string
		rating string
	)

	for rows.Next() {
		if err = rows.Scan(&theme, &rating); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  question: theme=%v rating=%v",
			theme, rating)
		stats = append(stats, &models.Stat{
			Theme:  theme,
			Rating: rating,
		})

	}

	return stats, nil
}
