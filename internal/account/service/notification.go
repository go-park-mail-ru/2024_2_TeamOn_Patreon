package service

import (
	"context"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/pkg/models"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) GetNotifications(ctx context.Context, userID string, opt *pkgModels.NotificationsOpt) ([]*pkgModels.Notification, error) {
	op := "internal.account.service.GetNotifications"

	var (
		notifications = []*pkgModels.Notification{}
		err           error
	)

	if opt.Status == pkgModels.NotReadStatus {
		logger.StandardDebugF(ctx, op, "want to get map NOT READ notifications for userID = %v", userID)
		notifications, err = s.rep.GetNotReadNotifications(ctx, userID, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
	} else {
		logger.StandardDebugF(ctx, op, "want to get map ALL notifications for userID = %v", userID)
		notifications, err = s.rep.GetAllNotifications(ctx, userID, opt.Offset, opt.Limit)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}
	}

	return notifications, nil
}

func (s *Service) GetNewNotifications(ctx context.Context, userID string, opt *pkgModels.NotificationsTimeOpt) ([]*pkgModels.Notification, error) {
	op := "internal.account.service.GetNewNotifications"

	logger.StandardDebugF(ctx, op, "want to get map NEW notifications for last time = %v seconds for userID = %v", opt.Time, userID)
	notifications, err := s.rep.GetNewNotificationsByTime(ctx, userID, opt.Time)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return notifications, nil
}

func (s *Service) ReadNotification(ctx context.Context, userID, notificationID string) error {
	op := "internal.account.service.ReadNotification"

	logger.StandardDebugF(ctx, op, "want to read notificationID= %v by userID = %v", notificationID, userID)
	if err := s.rep.ChangeNotificationStatus(ctx, userID, notificationID); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
