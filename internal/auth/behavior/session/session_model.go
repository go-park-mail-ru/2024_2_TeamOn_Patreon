package session

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"time"
)

type SessionModel struct {
	SessionID string
	UserID    string

	CreatedAt  time.Time
	FinishedAt time.Time
}

func NewSessionModel(sessionID, userID string) *SessionModel {
	return &SessionModel{
		SessionID:  sessionID,
		UserID:     userID,
		CreatedAt:  time.Now(),
		FinishedAt: time.Now().Add(global.SessionTimeH * time.Hour),
	}
}
