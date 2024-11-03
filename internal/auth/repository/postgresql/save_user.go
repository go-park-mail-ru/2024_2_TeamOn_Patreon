package postgresql

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

const (
	// Input: $1 - userId, $2 - username, $3 - role,  $4 - hash password
	saveUserSQl = `
INSERT INTO People (user_id, username, role_id, hash_password) VALUES
    ($1, $2, (SELECT role_id FROM Role WHERE role_default_name = $3), $4)
;
`
)

func (ar *AuthRepository) SaveUserWithRole(ctx context.Context, userId uuid.UUID, username string, role string, passwordHash string) error {
	op := "internal.auth.repository.SaveUserWithRole"

	_, err := ar.db.Exec(ctx, saveUserSQl, userId, username, role, passwordHash)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
