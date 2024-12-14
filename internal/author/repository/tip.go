package repositories

import (
	"context"
	"fmt"
	"time"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

var tipRequests = make(map[string]repModels.TipRequest)

func (p *Postgres) SaveTipRequest(ctx context.Context, tipReq repModels.TipRequest) error {
	op := "internal.author.repository.SaveTipRequest"

	// Пользователь, которому донатим, является автором
	isAuthor, err := p.isAuthor(ctx, tipReq.AuthorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if !isAuthor {
		return global.ErrUserIsNotAuthor
	}

	// Сохраняем в map до момента оплаты
	mu.Lock()
	defer mu.Unlock()
	tipRequests[tipReq.TipReqID] = tipReq

	return nil
}

func (p *Postgres) RealizeTipRequest(ctx context.Context, tipReqID string) (repModels.TipRequest, error) {
	op := "internal.author.repository.RealizeTipRequest"

	mu.Lock()
	defer mu.Unlock()

	// Получаем данные запроса на донат
	logger.StandardDebugF(ctx, op, "want to get tip request by reqID=%v", tipReqID)
	tipReq, exists := tipRequests[tipReqID]
	if !exists {
		return repModels.TipRequest{}, global.ErrTipReqDoesNotExist
	}

	// Запрос на добавление записи Tip
	query := `
		INSERT INTO 
			tip (tip_id, user_id, author_id, cost, message, payed_date)
        VALUES 
			($1, $2, $3, $4, $5, $6)
	`

	tipID := p.GenerateID()
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, tipID, tipReq.UserID, tipReq.AuthorID, tipReq.Cost, tipReq.Message, time.Now()); err != nil {
		return repModels.TipRequest{}, errors.Wrap(err, op)
	}

	// Удаляем запрос из map после реализации
	delete(subscriptionRequests, tipReqID)

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful create new record for authorID: %s", tipReq.AuthorID),
		op,
	)

	return tipReq, nil
}
