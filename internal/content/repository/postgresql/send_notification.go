package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (cr *ContentRepository) SendNotification(ctx context.Context, message, userID, authorID string) error {
	op := "internal.content.repository.postgresql.SendNotification"

	query := `
	INSERT INTO 
		notification (notification_id, user_id, sender_id, about)
	VALUES
		($1, $2, $3, $4);
	`
	notificationID := utils.GenerateUUID()

	// Не путать: userID ($2) = authorID (получатель), senderID ($3) = userID (отправитель)
	_, err := cr.db.Exec(ctx, query, notificationID, authorID, userID, message)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
