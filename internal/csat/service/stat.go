package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
	"strconv"
)

func (s *Behavior) GetSTATByTime(ctx context.Context, time string) ([]*models.Stat, error) {
	op := "csat.service.GetStatByTime"
	if time == Infinity {
		stats, err := s.rep.GetAllStat(ctx)
		if err != nil {
			err = errors.Wrap(err, op)
			return nil, err
		}
		return stats, nil
	}

	nDays, err := strconv.Atoi(time)
	if err != nil {
		err = global.ErrDaysIsNotDigital
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "not int")
	}

	if nDays <= 0 || nDays > 60 {
		err = global.ErrNotValidDays
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "invalid param")
	}

	stats, err := s.rep.GetStatByDays(ctx, nDays)
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "not stats")
	}
	return stats, nil
}
