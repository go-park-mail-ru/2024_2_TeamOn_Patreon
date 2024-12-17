/*
 * PushART - Posts | API
 *
 * API для управления постами и лентой
 */
package models

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

//go:generate easyjson -all

//easyjson:json
type PostLikeBody struct {
	// ID поста, который нужно лайкнуть или убрать лайк
	PostId string `json:"postId"`
}

func (plb *PostLikeBody) Validate() (bool, error) {
	if err := plb.validatePostId(); err != nil {
		return false, err
	}
	return true, nil
}

func (plb *PostLikeBody) validatePostId() error {
	op := "content.controller.model_post_like_body.validatePostId"

	if !utils.IsValidUUIDv4(plb.PostId) {
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	return nil
}
