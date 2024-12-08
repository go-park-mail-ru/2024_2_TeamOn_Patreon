package repositories

import (
	"context"
	"time"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	amountPostsByHour = `
		SELECT COUNT(*)
		FROM
			post
		WHERE
			created_date >= now() - interval '1 hour' * $2 AND created_date < now() - interval '1 hour' * $3 AND user_id = $1;
	`

	amountPostsByDay = `
		SELECT COUNT(*)
		FROM
			post
		WHERE
			created_date >= now() - interval '1 day' * $2 AND created_date < now() - interval '1 day' * $3 AND user_id = $1;
	`

	amountPostsByMonth = `
		SELECT COUNT(*)
		FROM
			post
		WHERE
			created_date >= now() - interval '1 month' * $2 AND created_date < now() - interval '1 month' * $3 AND user_id = $1;
	`

	amountPaymentsByHour = `
		WITH total_tips AS (
			SELECT 
				COALESCE(SUM(cost), 0) AS total_cost
			FROM 
				tip
			WHERE 
				payed_date >= now() - interval '1 hour' * $2 AND payed_date < now() - interval '1 hour' * $3 AND author_id = $1
			),
		total_subscriptions AS (
			SELECT 
				COALESCE(SUM(cs.cost), 0) AS total_cost
			FROM 
				subscription s
			LEFT JOIN
				custom_subscription cs ON cs.custom_subscription_id = s.custom_subscription_id
			WHERE 
				started_date >= now() - interval '1 hour' * $2 AND started_date < now() - interval '1 hour' * $3 AND cs.author_id = $1
		)

		SELECT 
			tt.total_cost + ts.total_cost AS total_payments
		FROM 
			total_tips tt,
			total_subscriptions ts;
	`
	amountPaymentsByDay = `
		WITH total_tips AS (
			SELECT 
				COALESCE(SUM(cost), 0) AS total_cost
			FROM 
				tip
			WHERE 
				payed_date >= now() - interval '1 day' * $2 AND payed_date < now() - interval '1 day' * $3 AND author_id = $1
			),
		total_subscriptions AS (
			SELECT 
				COALESCE(SUM(cs.cost), 0) AS total_cost
			FROM 
				subscription s
			LEFT JOIN
				custom_subscription cs ON cs.custom_subscription_id = s.custom_subscription_id
			WHERE 
				started_date >= now() - interval '1 day' * $2 AND started_date < now() - interval '1 day' * $3 AND cs.author_id = $1
		)

		SELECT 
			tt.total_cost + ts.total_cost AS total_payments
		FROM 
			total_tips tt,
			total_subscriptions ts;
	`
	amountPaymentsByMonth = `
		WITH total_tips AS (
			SELECT 
				COALESCE(SUM(cost), 0) AS total_cost
			FROM 
				tip
			WHERE 
				payed_date >= now() - interval '1 month' * $2 AND payed_date < now() - interval '1 month' * $3 AND author_id = $1
			),
		total_subscriptions AS (
			SELECT 
				COALESCE(SUM(cs.cost), 0) AS total_cost
			FROM 
				subscription s
			LEFT JOIN
				custom_subscription cs ON cs.custom_subscription_id = s.custom_subscription_id
			WHERE 
				started_date >= now() - interval '1 month' * $2 AND started_date < now() - interval '1 month' * $3 AND cs.author_id = $1
		)

		SELECT 
			tt.total_cost + ts.total_cost AS total_payments
		FROM 
			total_tips tt,
			total_subscriptions ts;
	`
)

func (p *Postgres) GetStatByDay(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error) {
	op := "internal.author.repository.GetStatByDay"

	var (
		graphic = &pkgModels.Graphic{}
	)

	// Текущий час
	currentTime := time.Now()             // Получаем текущее время
	currentHour := currentTime.Hour() + 3 // Извлекаем текущий час + смещение по часовому поясу
	logger.StandardDebugF(ctx, op, "Get currentHour=%v", currentHour)

	logger.StandardDebugF(ctx, op, "Want to get points for graphic")
	// Заполняем PointsX часами от текущего часа прошлого дня до текущего часа сегодняшнего дня
	// PointsY заполняется количеством постов соответствующим каждому часу
	for i := 1; i <= 24; i++ {
		hour := (i + currentHour) % 24
		graphic.PointsX = append(graphic.PointsX, hour)

		amountOfPosts, err := p.amountPostsByHour(ctx, userID, statParam, 24-i)
		if err != nil {
			return graphic, errors.Wrap(err, op)
		}
		graphic.PointsY = append(graphic.PointsY, amountOfPosts)
	}

	return graphic, nil
}

func (p *Postgres) GetStatByMonth(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error) {
	op := "internal.author.repository.GetStatByMonth"

	var (
		graphic = &pkgModels.Graphic{}
	)

	// Текущий день
	currentTime := time.Now()       // Получаем текущее время
	currentDay := currentTime.Day() // Извлекаем текущий день
	logger.StandardDebugF(ctx, op, "Get currentDay=%v", currentDay)

	logger.StandardDebugF(ctx, op, "Want to get points for graphic")
	// Заполняем PointsX днями от текущей даты прошлого месяца до текущей даты настоящего месяца
	// PointsY заполняется количеством постов соответствующим каждому дню
	for i := 1; i <= 30; i++ {
		day := (i+currentDay)%31 + 1
		graphic.PointsX = append(graphic.PointsX, day)

		amountOfPosts, err := p.amountPostsByDay(ctx, userID, statParam, 30-i)
		if err != nil {
			return graphic, errors.Wrap(err, op)
		}
		graphic.PointsY = append(graphic.PointsY, amountOfPosts)
	}

	return graphic, nil
}

func (p *Postgres) GetStatByYear(ctx context.Context, userID, statParam string) (*pkgModels.Graphic, error) {
	op := "internal.author.repository.GetStatByYear"

	var (
		graphic = &pkgModels.Graphic{}
	)

	// Текущий месяц
	currentTime := time.Now()             // Получаем текущее время
	currentMonth := currentTime.Day() + 4 // Извлекаем текущий месяц
	logger.StandardDebugF(ctx, op, "Get currentMonth=%v", currentMonth)

	logger.StandardDebugF(ctx, op, "Want to get points for graphic")
	// Заполняем PointsX днями от текущей даты прошлого месяца до текущей даты настоящего месяца
	// PointsY заполняется количеством постов соответствующим каждому дню
	for i := 1; i <= 12; i++ {
		month := (i+currentMonth)%13 + 1
		graphic.PointsX = append(graphic.PointsX, month)

		amountOfPosts, err := p.amountPostsByMonth(ctx, userID, statParam, 12-i)
		if err != nil {
			return graphic, errors.Wrap(err, op)
		}
		graphic.PointsY = append(graphic.PointsY, amountOfPosts)
	}

	return graphic, nil
}

func (p *Postgres) amountPostsByHour(ctx context.Context, userID, statParam string, hour int) (int, error) {
	op := "internal.author.repository.amountPostsByHour"

	var (
		amount int
		query  string
	)

	// Используем запрос в зависимости от требуемого параметра выборки для статистики
	if statParam == global.StatPosts {
		query = amountPostsByHour
	} else if statParam == global.StatPayments {
		query = amountPaymentsByHour
	}

	err := p.db.QueryRow(ctx, query, userID, hour, hour-1).Scan(&amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrap(err, op)
	}
	return amount, nil
}

func (p *Postgres) amountPostsByDay(ctx context.Context, userID, statParam string, day int) (int, error) {
	op := "internal.author.repository.amountPostsByDay"

	var (
		amount int
		query  string
	)

	// Используем запрос в зависимости от требуемого параметра выборки для статистики
	if statParam == global.StatPosts {
		query = amountPostsByDay
	} else if statParam == global.StatPayments {
		query = amountPaymentsByDay
	}

	err := p.db.QueryRow(ctx, query, userID, day, day-1).Scan(&amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrap(err, op)
	}
	return amount, nil
}

func (p *Postgres) amountPostsByMonth(ctx context.Context, userID, statParam string, month int) (int, error) {
	op := "internal.author.repository.amountPostsByMonth"

	var (
		amount int
		query  string
	)

	// Используем запрос в зависимости от требуемого параметра выборки для статистики
	if statParam == global.StatPosts {
		query = amountPostsByMonth
	} else if statParam == global.StatPayments {
		query = amountPaymentsByMonth
	}

	err := p.db.QueryRow(ctx, query, userID, month, month-1).Scan(&amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrap(err, op)
	}
	return amount, nil
}
