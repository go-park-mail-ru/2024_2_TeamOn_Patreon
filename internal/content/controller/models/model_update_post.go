/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models/validate"
	validate2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
	"github.com/pkg/errors"
)

// Все, чего нет - не заменяется

type UpdatePost struct {
	// Идентификатор поста для обновления
	PostId string `json:"postId"`
	// Заголовок поста
	Title string `json:"title,omitempty"`
	// Содержимое поста (текст)
	Content string `json:"content,omitempty"`
	// Уровень на котором можно смотреть пост
	Layer int `json:"layer"`
}

func (ap *UpdatePost) Validate() (bool, error) {
	op := "content.controller.model_update_post.Validate"

	err := ap.validatePostId()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = ap.validateTitle()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = ap.validateContent()
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = ap.validateLayer()
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	return true, nil
}

func (ap *UpdatePost) validateTitle() error {
	op := "content.controller.model_update_post.validateTitle"

	if ap.Title == "" {
		return nil
	}

	ap.Title = validate2.Sanitize(ap.Title)

	if err := validate.Title(ap.Title); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (ap *UpdatePost) validateContent() error {
	op := "content.controller.model_update_post.validateContent"

	ap.Content = validate2.Sanitize(ap.Content)

	if err := validate.Content(ap.Content); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (ap *UpdatePost) validateLayer() error {
	op := "content.controller.model_update_post.validateLayer"

	if err := validate.Layer(ap.Layer); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (ap *UpdatePost) validatePostId() error {
	op := "content.controller.model_update_post.validatePostId"

	if err := validate.Uuid(ap.PostId); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
