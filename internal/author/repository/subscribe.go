package repositories

import (
	"context"
	"database/sql"
	"sync"
	"time"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/jackc/pgx/v5"

	"github.com/pkg/errors"
)

// Для хранения ID запросов на подписку
var subscriptionRequests = make(map[string]repModels.SubscriptionRequest)
var mu sync.Mutex

func (p *Postgres) SaveSubscribeRequest(ctx context.Context, subReq repModels.SubscriptionRequest) error {
	op := "internal.author.repository.SaveSubscribeRequest"

	// Пользователь, на которого подписываемся, является автором
	isAuthor, err := p.isAuthor(ctx, subReq.AuthorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if !isAuthor {
		return global.ErrUserIsNotAuthor
	}

	// Выбранный уровень подписки существует
	customSubID, err := p.getCustomSubscriptionID(ctx, subReq.AuthorID, subReq.Layer)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if customSubID == "" {
		return global.ErrCustomSubDoesNotExist
	}

	// Сохраняем в map до момента оплаты
	mu.Lock()
	defer mu.Unlock()
	subscriptionRequests[subReq.SubReqID] = subReq

	return nil
}

func (p *Postgres) RealizeSubscribeRequest(ctx context.Context, subReqID string) (string, error) {
	op := "internal.author.repository.RealizeSubscribeRequest"

	mu.Lock()
	defer mu.Unlock()

	// Получаем данные запроса на подписку
	logger.StandardDebugF(ctx, op, "want to get subscription request by reqID=%v", subReqID)
	subReq, exists := subscriptionRequests[subReqID]
	if !exists {
		return "", global.ErrSubReqDoesNotExist
	}

	customSubID, _ := p.getCustomSubscriptionID(ctx, subReq.AuthorID, subReq.Layer)
	subID := utils.GenerateUUID()

	// Сохраняем запись о подписке
	logger.StandardDebugF(ctx, op, "want to save new record subID=%v, userID=%v, customSubID=%v", subID, subReq.UserID, customSubID)

	currentTime := time.Now()
	finishedDate := currentTime.AddDate(0, subReq.MonthCount, 0)

	query := `
    	INSERT INTO public.subscription(
        subscription_id, user_id, custom_subscription_id, started_date, finished_date)
    	VALUES ($1, $2, $3, $4, $5)
	`
	if _, err := p.db.Exec(ctx, query, subID, subReq.UserID, customSubID, currentTime, finishedDate); err != nil {
		return "", errors.Wrap(err, op)
	}
	// Удаляем запрос из map после реализации
	delete(subscriptionRequests, subReqID)

	return customSubID, nil
}

func (p *Postgres) GetCustomSubscriptionInfo(ctx context.Context, customSubID string) (string, string, error) {
	op := "internal.author.repository.GetCustomSubscriptionInfo"

	query := `
        SELECT 
			author_id, custom_name
        FROM
			custom_subscription
        WHERE
			custom_subscription_id = $1;
	`
	var (
		authorID   string
		customName string
	)

	err := p.db.QueryRow(ctx, query, customSubID).Scan(&authorID, &customName)

	if err != nil {
		return "", "", errors.Wrap(err, op)
	}

	return authorID, customName, nil
}

func (p *Postgres) GetCostCustomSub(ctx context.Context, authorID string, layer int) (int, error) {
	op := "internal.author.repository.GetCostCustomSub"

	query := `
		SELECT 
			cs.cost
		FROM
			custom_subscription cs
		JOIN
			subscription_layer sl ON cs.subscription_layer_id = sl.subscription_layer_id
		WHERE 
			sl.layer = $1 AND cs.author_id = $2;
	`

	var cost int
	if err := p.db.QueryRow(ctx, query, layer, authorID).Scan(&cost); err != nil {
		if err == pgx.ErrNoRows {
			return 0, global.ErrCustomSubDoesNotExist
		}
		return 0, errors.Wrap(err, op)
	}

	return cost, nil
}

// GetUsername - получение имени пользователя по userID
func (p *Postgres) GetUsername(ctx context.Context, userID string) (string, error) {
	op := "internal.account.repository.GetUsername"

	query := `
		SELECT 
			username
		FROM
			people
		WHERE
			user_id = $1;
	`

	var username string
	err := p.db.QueryRow(ctx, query, userID).Scan(&username)

	if err != nil {
		return "", errors.Wrap(err, op)
	}

	return username, nil
}

func (p *Postgres) SendNotification(ctx context.Context, message, userID, authorID string) error {
	op := "internal.content.repository.postgresql.SendNotification"

	query := `
	INSERT INTO 
		notification (notification_id, user_id, sender_id, about)
	VALUES
		($1, $2, $3, $4);
	`
	notificationID := utils.GenerateUUID()

	// Не путать: userID ($2) = authorID (получатель), senderID ($3) = userID (отправитель)
	_, err := p.db.Exec(ctx, query, notificationID, authorID, userID, message)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (p *Postgres) getCustomSubscriptionID(ctx context.Context, authorID string, layer int) (string, error) {
	op := "internal.author.repository.checkSubscriptionExists"

	var customSubscriptionID string
	query := `
        SELECT cs.custom_subscription_id
        FROM custom_subscription cs
        JOIN subscription_layer sl ON cs.subscription_layer_id = sl.subscription_layer_id
        WHERE cs.author_id = $1
          AND sl.layer = $2;`

	err := p.db.QueryRow(ctx, query, authorID, layer).Scan(&customSubscriptionID)

	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", errors.Wrap(err, op)
	}

	return customSubscriptionID, nil
}

func (p *Postgres) isAuthor(ctx context.Context, authorID string) (bool, error) {
	op := "internal.author.repository.isAuthor"

	query := `
		SELECT 
			r.role_default_name
		FROM 
			Role r
		JOIN 
			People p ON r.role_id = p.role_id
		WHERE 
			p.user_id = $1
	`

	rows, err := p.db.Query(ctx, query, authorID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		role string
	)

	for rows.Next() {
		if err = rows.Scan(&role); err != nil {
			return false, errors.Wrap(err, op)
		}
		if role == string(models.Author) {
			return true, nil
		}
	}
	return false, nil
}
