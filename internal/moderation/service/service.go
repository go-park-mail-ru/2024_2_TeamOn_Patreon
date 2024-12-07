package service

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/service/interfaces"

type Service struct {
	rep interfaces.ModerationRepository
}

func New(rep interfaces.ModerationRepository) *Service {
	return &Service{rep}
}
