package repositories

import (
	"context"
	"database/sql"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Поле структуры - pool соединений с БД
type Postgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) AuthorByID(ctx context.Context, authorID string) (*sModels.Author, error) {
	op := "internal.author.repository.AuthorByID"

	// SQL-запрос для получения username, info
	query := `
		SELECT 
			p.username,	pg.info
		FROM 
			page pg
		JOIN 
			people p ON pg.user_id = p.user_id
		WHERE 
			pg.user_id = $1;
	`

	var author sModels.Author

	if err := p.db.QueryRow(ctx, query, authorID).Scan(&author.Username, &author.Info); err != nil {
		if err == sql.ErrNoRows {
			// Если автор не найден, возвращаем nil без ошибки
			logger.StandardInfoF(
				"author with authorID='%v' not found", authorID,
				op)
			return nil, nil
		}
		logger.StandardDebugF(op, "get author error: {%v}", err)
		return nil, err
	}

	// SQL-запрос для получения followers
	queryFollowers := `
		SELECT 
			COUNT(DISTINCT s.user_id) AS followers
		FROM 
			subscription s
		JOIN 
			custom_subscription cs ON s.custom_subscription_id = cs.custom_subscription_id
		WHERE 
			cs.author_id = $1;
	`

	if err := p.db.QueryRow(ctx, queryFollowers, authorID).Scan(&author.Followers); err != nil {
		if err == sql.ErrNoRows {
			// Если подписчики не найдены, проставляем ноль
			author.Followers = 0
		}
		logger.StandardDebugF(op, "get followers error: {%v}", err)
		return nil, err
	}

	// Возвращаем данные автора
	return &author, nil
}
