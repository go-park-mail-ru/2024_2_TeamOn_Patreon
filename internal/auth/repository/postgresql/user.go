package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// getUserByUserIdSQL получает username и роль по ид пользователя
	// Input: $1 - userId
	// Output: username, user role
	getUserByUserIdSQL = `
SELECT username, role_default_name as role
FROM people
	JOIN Role USING(role_id)
WHERE user_id = $1
;
`

	// getUserByUsernameSQL получает username и роль по ид пользователя
	// Input: $1 - username
	// Output: userId, user role
	getUserByUsernameSQL = `
SELECT user_id as userId, role_default_name as role
FROM people
	JOIN Role USING(role_id)
WHERE username = $1
;
`
	// getUserExistUsernameSQL - существует ли пользователь с таким именем
	// Input: $1 - username
	// Output: 1 - да, 0 - нет
	getUserExistUsernameSQL = `
SELECT CASE 
           WHEN EXISTS (
               SELECT 1 
               FROM people 
               WHERE username = $1
           ) THEN 1
           ELSE 0
       END as userExist;
`
)

func (a *AuthRepository) GetUserByUserId(ctx context.Context, userId uuid.UUID) (string, string, error) {
	op := "internal.auth.repository.getUserByUserId"

	rows, err := a.db.Query(ctx, getUserByUserIdSQL, userId)
	if err != nil {
		return "", "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		username string
		role     string
	)

	for rows.Next() {
		if err = rows.Scan(&username, &role); err != nil {
			return "", "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got username='%s' role = %v of userId='%v'", userId, role, userId)
		return username, role, nil
	}
	return "", "", errors.Wrap(global.ErrUserNotFound, op)
}

func (a *AuthRepository) GetUserByUsername(ctx context.Context, username string) (uuid.UUID, string, error) {
	op := "internal.auth.repository.getUserByUsername"

	rows, err := a.db.Query(ctx, getUserByUsernameSQL, username)
	if err != nil {
		return uuid.Nil, "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		userId uuid.UUID
		role   string
	)

	for rows.Next() {
		if err = rows.Scan(&userId, &role); err != nil {
			return uuid.Nil, "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got username='%s' role = %v of userId='%v'", userId, role, userId)
		return userId, role, nil
	}
	return uuid.Nil, "", errors.Wrap(global.ErrUserNotFound, op)

}

func (a *AuthRepository) UserExists(ctx context.Context, username string) (bool, error) {
	op := "internal.auth.repository.userExists"

	rows, err := a.db.Query(ctx, getUserExistUsernameSQL, username)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		userExist int
	)

	for rows.Next() {
		if err = rows.Scan(&userExist); err != nil {
			return false, errors.Wrap(err, op)
		}
		exist := userExist == 1
		logger.StandardDebugF(ctx, op, "Got username='%s' userExist='%v'", username, exist)

		return exist, nil
	}
	return false, nil
}
