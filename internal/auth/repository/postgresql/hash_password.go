package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// getPasswordHashSQL
	// Input: $1 - userId
	// Output: password hash
	getPasswordHashSQL = `
SELECT hash_password
FROM people
WHERE user_id = $1
`
)

func (a *AuthRepository) GetPasswordHashByID(ctx context.Context, userID uuid.UUID) (string, error) {
	op := "internal.auth.repository.getPasswordHashByID"

	rows, err := a.db.Query(ctx, getPasswordHashSQL, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		hash string
	)

	for rows.Next() {
		if err = rows.Scan(&hash); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(op, "GetPasswordHashByID found hash: %s", hash)
		return hash, nil
	}
	return "", errors.Wrap(global.ErrUserNotFound, op)
}
