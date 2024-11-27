package models

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/pkg/validate"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

// AddCustomSubscription Данные для добавления кастомной подписки, которые отправляет фронт
type AddCustomSubscription struct {
	// Заголовок кастномной (платной) подписки
	Title string `json:"title"`
	// Описание кастомной (платной) подписки
	Description string `json:"description,omitempty"`
	// Стоимость  кастомной (платной) подписки в рублях
	Cost int `json:"cost"`
	// Уровень кастомной (платной) подписки. Уровень (0:3]
	Layer int `json:"layer"`
}

func (acp *AddCustomSubscription) Validate() error {
	op := "custom_subscription.controller.model_add_custom_subscription.Validate"
	err := errors.Wrap(global.ErrBadRequest, op)

	ok := acp.validateTitle()
	if !ok {
		return errors.Wrap(err, "title")
	}
	ok = acp.validateDescription()
	if !ok {
		return errors.Wrap(err, "description")
	}
	ok = acp.validateCost()
	if !ok {
		return errors.Wrap(err, "cost")
	}
	ok = acp.validateLayer()
	if !ok {
		return errors.Wrap(err, "layer")
	}
	return nil
}

func (acp *AddCustomSubscription) validateTitle() bool {
	title, ok := validate.Title(acp.Title)
	if !ok {
		return false
	}
	acp.Title = title
	return true
}

func (acp *AddCustomSubscription) validateDescription() bool {
	if acp.Description == "" {
		return true
	}
	description, ok := validate.Description(acp.Description)
	if !ok {
		return false
	}
	acp.Description = description
	return true
}

func (acp *AddCustomSubscription) validateCost() bool {
	if ok := validate.Cost(acp.Cost); !ok {
		return false
	}
	return true
}
func (acp *AddCustomSubscription) validateLayer() bool {
	if ok := validate.Layer(acp.Layer); !ok {
		return false
	}
	return true
}
