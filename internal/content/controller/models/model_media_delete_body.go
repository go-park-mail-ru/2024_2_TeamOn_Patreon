package models

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

//go:generate easyjson -all

//easyjson:json
type MediaDeleteRequest struct {
	MediaIDs []string `json:"mediaID"`
}

func (mdr *MediaDeleteRequest) Validate() (bool, error) {
	if err := mdr.validateMediaID(); err != nil {
		return false, err
	}
	return true, nil
}

func (mdr *MediaDeleteRequest) validateMediaID() error {
	op := "content.controller.model_media_delete_body.validateMediaID"

	for _, mediaID := range mdr.MediaIDs {
		if !utils.IsValidUUIDv4(mediaID) {
			return errors.Wrap(global.ErrIsInvalidUUID, op)
		}
	}

	return nil
}
