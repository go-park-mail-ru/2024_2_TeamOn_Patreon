package repositories

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// создает подписку
	createSubscription = `
INSERT INTO public.subscription(
	subscription_id, user_id, custom_subscription_id, started_date, finished_date)
	VALUES ($3, $1, $2, NOW(), NOW() + INTERVAL '30 days');
`

	// создает кастомную подписку
	createCustomSubscription = `
INSERT INTO public.custom_subscription(
	custom_subscription_id, author_id, custom_name, cost, subscription_layer_id)
	VALUES ($2, $1, 'default', 10, (select  subscription_layer_id from subscription_layer where layer = 1));
`

	// проверяет есть ли кастомная подписка у автора
	getCustomSub = `
select 
	custom_subscription_id
from 
	Custom_Subscription
where author_id = $1
	and subscription_layer_id = (select subscription_layer_id from subscription_layer where layer = 1)
	`

	// getSubscription
	getSubscription = `
select 
custom_subscription_id
from
    Subscription
    JOIN Custom_Subscription USING (custom_subscription_id)
where Custom_Subscription.author_id = $1 and Subscription.user_id = $2
LIMIT 1;
`

	// автор ли юзер
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

	deleteSubscription = `
DELETE FROM subscription
USING custom_subscription
WHERE subscription.custom_subscription_id = custom_subscription.custom_subscription_id
  AND subscription.user_id = $1
  AND custom_subscription.author_id = $2;
`
)

func (p *Postgres) Subscribe(ctx context.Context, userID string, authorID string) (bool, error) {
	op := "internal.author.repository.Subscribe"
	logger.StandardDebugF(ctx, op, "Want to is author UserId %v AuthorId %v", userID, authorID)

	isAuthor, err := p.isAuthor(ctx, authorID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if !isAuthor {
		return false, errors.New("not author")
	}
	logger.StandardDebugF(ctx, op, "isAuthor %v", isAuthor)

	customSubID, err := p.getCustomSubscriptionLayerOne(ctx, authorID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if customSubID == "" {
		err = p.createCustomSubscription(ctx, authorID)
		if err != nil {
			return false, errors.Wrap(err, op)
		}
	}

	logger.StandardDebugF(ctx, op, "CustomSubId %v", customSubID)

	customSubID, err = p.getCustomSubscriptionLayerOne(ctx, authorID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	subId, err := p.getSubscription(ctx, userID, authorID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "SubId %v", subId)
	if subId == "" {
		err = p.createSubscription(ctx, customSubID, userID)
		if err != nil {
			return false, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "create sub custoSsub %v user %v", customSubID, userID)
		return true, nil
	}
	logger.StandardDebugF(ctx, op, "SubId %v", subId)

	err = p.deleteSubscription(ctx, authorID, userID)
	logger.StandardDebugF(ctx, op, "successful deleted subId %v", subId)

	return false, nil

}

func (p *Postgres) createCustomSubscription(ctx context.Context, authorId string) error {
	op := "internal.author.repository.CreateCustomSubscription"
	customSubID := p.GenerateID()
	_, err := p.db.Exec(ctx, createCustomSubscription, authorId, customSubID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (p *Postgres) isAuthor(ctx context.Context, authorId string) (bool, error) {
	op := "internal.author.repository.isAuthor"

	rows, err := p.db.Query(ctx, getUserRoleSQL, authorId)
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
		if role == "Author" {
			return true, nil
		}
	}
	return false, nil
}

func (p *Postgres) getCustomSubscriptionLayerOne(ctx context.Context, authorId string) (string, error) {
	op := "internal.custom_subscription.getCustomSubscriptionLayerOne"
	rows, err := p.db.Query(ctx, getCustomSub, authorId)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	defer rows.Close()
	var (
		customSubId string
	)
	for rows.Next() {
		if err = rows.Scan(&customSubId); err != nil {
			return "", errors.Wrap(err, op)
		}
		return customSubId, nil

	}
	return "", err
}

func (p *Postgres) createSubscription(ctx context.Context, customSubID, userID string) error {
	op := "internal.author.repository.CreateSubscription"
	subscriptionID := p.GenerateID()
	_, err := p.db.Exec(ctx, createSubscription, userID, customSubID, subscriptionID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil

}

//func (p *Postgres) getSubscription(ctx context.Context, userId)

func (p *Postgres) getSubscription(ctx context.Context, userID, authorID string) (string, error) {
	op := "internal.subscription.getSubscription"

	rows, err := p.db.Query(ctx, getSubscription, authorID, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()
	var (
		SubId string
	)
	for rows.Next() {
		if err = rows.Scan(&SubId); err != nil {
			return "", errors.Wrap(err, op)
		}
		return SubId, nil

	}
	return "", err
}

func (p *Postgres) deleteSubscription(ctx context.Context, authorID, userID string) error {
	op := "internal.subscription.deleteSubscription"
	_, err := p.db.Exec(ctx, deleteSubscription, userID, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
