package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

const (
	getUserLayerSql = `
-- First result set.
select
	layer
from
	subscription_layer 
	join custom_subscription on subscription_layer.subscription_layer_id = custom_subscription.subscription_layer_id
	join subscription on subscription.custom_subscription_id = custom_subscription.custom_subscription_id
where
	subscription.user_id = ?
	and custom_subscription.author_id = ?
;
`

	// getUserRoleSQL - возвращает текстовое имя роли пользователя
	// Input: userId
	// Output: role_name ('Reader' or 'Author' or '')
	getUserRoleSQL = `
		select 
			role_default_name
		from 
		    Role
			join People USING (role_id)
		where People.user_id = $1
	`
)

func (cr *ContentRepository) GetUserLayerForAuthor(ctx context.Context, userId uuid.UUID, authorId uuid.UUID) (int, error) {
	op := "internal.content.repository.user.GetUserLayerForAuthor"

	rows, err := cr.db.Query(ctx, getUserLayerSql, userId, authorId)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		layer int
	)
	for rows.Next() {
		if err = rows.Scan(&layer); err != nil {
			return 0, errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got layer= %s user= %s author %s", layer, userId, authorId)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}
	logger.StandardDebugF(op, "Got layer= %s user= %s author %s", layer, userId, authorId)

	return 0, nil
}

func (cr *ContentRepository) GetUserRole(ctx context.Context, userId uuid.UUID) (string, error) {
	op := "internal.content.repository.user.GetUserRole"

	logger.StandardDebugF(op, "Want to get user role userID=%v, db = %v", userId, cr.db)

	rows, err := cr.db.Query(ctx, getUserRoleSQL, userId.String())
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		role string
	)
	for rows.Next() {
		if err = rows.Scan(&role); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "Got layer= %s user= %s", role, userId)
		return role, nil
	}

	return "", nil
}
