package service

import (
	"context"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

func (s *Service) GetNotifications(ctx context.Context, userID string, opt *pkgModels.NotificationsOpt) ([]*sModels.Notification, error) {
	return []*sModels.Notification{}, nil
}
