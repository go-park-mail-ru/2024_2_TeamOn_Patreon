package repositories

import (
	"time"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"

	"context"
)

func (p *Postgres) GetAllNotifications(ctx context.Context, userID string, offset, limit int) ([]*pkgModels.Notification, error) {
	op := "internal.account.repository.GetAllNotifications"

	query := `
		SELECT 
			notification_id, sender_id, about, is_viewed, created_at
		FROM 
			notification
		WHERE
			user_id = $1
		ORDER BY
			created_at DESC
		LIMIT $3
		OFFSET $2;
	`

	notifications := make([]*pkgModels.Notification, 0)

	rows, err := p.db.Query(ctx, query, userID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		notificationID string
		message        string
		senderID       string
		isRead         bool
		createdAt      time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&notificationID, &senderID, &message, &isRead, &createdAt); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got notification: notificationID=%v message=%v senderID=%v isRead=%v createdAt=%v",
			notificationID, message, senderID, isRead, createdAt)
		notifications = append(notifications, &pkgModels.Notification{
			NotificationID: notificationID,
			Message:        message,
			SenderID:       senderID,
			IsRead:         isRead,
			CreatedAt:      createdAt,
		})
	}

	return notifications, nil
}

func (p *Postgres) GetNotReadNotifications(ctx context.Context, userID string, offset, limit int) ([]*pkgModels.Notification, error) {
	op := "internal.account.repository.GetAllNotifications"

	query := `
		SELECT 
			notification_id, sender_id, about, is_viewed, created_at
		FROM 
			notification
		WHERE
			user_id = $1 AND is_viewed = false
		ORDER BY
			created_at DESC
		LIMIT $3
		OFFSET $2;
	`

	notifications := make([]*pkgModels.Notification, 0)

	rows, err := p.db.Query(ctx, query, userID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		notificationID string
		message        string
		senderID       string
		isRead         bool
		createdAt      time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&notificationID, &senderID, &message, &isRead, &createdAt); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got notification: notificationID=%v message=%v senderID=%v isRead=%v createdAt=%v",
			notificationID, message, senderID, isRead, createdAt)
		notifications = append(notifications, &pkgModels.Notification{
			NotificationID: notificationID,
			Message:        message,
			SenderID:       senderID,
			IsRead:         isRead,
			CreatedAt:      createdAt,
		})
	}

	return notifications, nil
}

func (p *Postgres) GetNewNotificationsByTime(ctx context.Context, userID string, timeParam int) ([]*pkgModels.Notification, error) {
	op := "internal.account.repository.GetAllNotifications"

	query := `
		SELECT 
			notification_id, sender_id, about, is_viewed, created_at
		FROM 
			notification
		WHERE
			user_id = $1 AND created_at >= NOW() - INTERVAL '1 second' * $2
		ORDER BY
			created_at DESC
	`

	notifications := make([]*pkgModels.Notification, 0)

	rows, err := p.db.Query(ctx, query, userID, timeParam)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		notificationID string
		message        string
		senderID       string
		isRead         bool
		createdAt      time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&notificationID, &senderID, &message, &isRead, &createdAt); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got notification: notificationID=%v message=%v senderID=%v isRead=%v createdAt=%v",
			notificationID, message, senderID, isRead, createdAt)

		notifications = append(notifications, &pkgModels.Notification{
			NotificationID: notificationID,
			Message:        message,
			SenderID:       senderID,
			IsRead:         isRead,
			CreatedAt:      createdAt,
		})
	}

	return notifications, nil
}

func (p *Postgres) ChangeNotificationStatus(ctx context.Context, userID, notificationID string) error {
	op := "internal.account.repository.ChangeNotificationStatus"

	// SQL-запрос для обновления имени пользователя
	query := `
		UPDATE 
			notification 
		SET 
			is_viewed = true 
		WHERE 
			user_id = $1 AND notification_id = $2;
	`

	// Выполняем запрос
	_, err := p.db.Exec(ctx, query, userID, notificationID)
	if err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}
