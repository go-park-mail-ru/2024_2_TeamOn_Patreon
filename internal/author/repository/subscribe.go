package repositories

import (
	"context"
	"database/sql"
	"sync"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	utils "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"

	"github.com/pkg/errors"
)

var subscriptionRequests = make(map[string]repModels.SubscriptionRequest)
var mu sync.Mutex

func (p *Postgres) CreateSubscribeRequest(ctx context.Context, subReq repModels.SubscriptionRequest) (string, error) {
	op := "internal.author.repository.CreateSubscribeRequest"

	// Пользователь, на которого подписываемся, является автором
	isAuthor, err := p.isAuthor(ctx, subReq.AuthorID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if !isAuthor {
		return "", global.ErrUserIsNotAuthor
	}

	// Выбранный уровень подписки существует
	customSubID, err := p.getCustomSubscriptionID(ctx, subReq.AuthorID, subReq.Layer)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if customSubID == "" {
		return "", global.ErrCustomSubDoesNotExist
	}

	subReqID := utils.GenerateUUID()

	// Сохраняем в map до момента оплаты
	mu.Lock()
	defer mu.Unlock()
	subscriptionRequests[subReqID] = subReq

	return subReqID, nil
}

func (p *Postgres) RealizeSubscribeRequest(ctx context.Context, subReqID string) error {
	op := "internal.author.repository.RealizeSubscribeRequest"

	mu.Lock()
	defer mu.Unlock()

	// Получаем данные запроса на подписку
	logger.StandardDebugF(ctx, op, "want to get subscription request by reqID=%v", subReqID)
	subReq, exists := subscriptionRequests[subReqID]
	if !exists {
		return global.ErrSubReqDoesNotExist
	}

	customSubID, _ := p.getCustomSubscriptionID(ctx, subReq.AuthorID, subReq.Layer)
	subID := utils.GenerateUUID()

	// Сохраняем запись о подписке
	logger.StandardDebugF(ctx, op, "want to save new record subID=%v, userID=%v, customSubID=%v", subID, subReq.UserID, customSubID)

	query := `
    	INSERT INTO public.subscription(
        subscription_id, user_id, custom_subscription_id, started_date, finished_date)
    	VALUES ($1, $2, $3, NOW(), NOW() + INTERVAL $4 MONTH)
	`
	if _, err := p.db.Exec(ctx, query, subID, subReq.UserID, customSubID, subReq.MonthCount); err != nil {
		return errors.Wrap(err, op)
	}
	// Удаляем запрос из map после реализации
	delete(subscriptionRequests, subReqID)

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
