/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */

package models

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/validate"
	pkgValidate "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
	"github.com/pkg/errors"
)

type AddPost struct {
	// Заголовок поста
	Title string `json:"title"`
	// Содержимое поста
	Content string `json:"content,omitempty"`
	// Уровень подписки, на котором можно смотреть пост, по умолчанию - для всех
	Layer int `json:"layer"`
}

func (ap *AddPost) Validate() (bool, error) {
	op := "content.controller.model_add_post.Validate"

	err := ap.validateTitle()
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

func (ap *AddPost) validateTitle() error {
	op := "content.controller.model_add_post.validateTitle"

	ap.Title = pkgValidate.Sanitize(ap.Title)

	if err := validate.Title(ap.Title); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (ap *AddPost) validateContent() error {
	op := "content.controller.model_add_post.validateContent"

	ap.Content = pkgValidate.Sanitize(ap.Content)

	if err := validate.Content(ap.Content); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (ap *AddPost) validateLayer() error {
	op := "content.controller.model_add_post.validateLayer"

	if err := validate.Layer(ap.Layer); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
