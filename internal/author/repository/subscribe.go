package repositories

import (
	"context"

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
)

func (p *Postgres) Subscribe(ctx context.Context, userID string, authorID string) error {
	op := "internal.author.repository.Subscribe"
	isAuthor, err := p.isAuthor(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if !isAuthor {
		return errors.New("not author")
	}

	customSubID, err := p.getCustomSubscription(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if customSubID == "" {
		err = p.createCustomSubscription(ctx, authorID)
		if err != nil {
			return errors.Wrap(err, op)
		}
	}
	customSubID, err = p.getCustomSubscription(ctx, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	err = p.createSubscription(ctx, customSubID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil

}

func (p *Postgres) createCustomSubscription(ctx context.Context, authorId string) error {
	op := "internal.author.repository.CreateCustomSubscription"
	customSubscrID := p.GenerateID()
	_, err := p.db.Exec(ctx, createCustomSubscription, authorId, customSubscrID)
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

func (p *Postgres) getCustomSubscription(ctx context.Context, authorId string) (string, error) {
	op := "internal.custom_subscription.getCustomSubscription"
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
