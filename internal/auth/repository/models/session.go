package models

import (
	bmodels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/session"
	"time"
)

// Session модель репозитория
type Session struct {
	SessionID SessionID
	UserID    UserID

	CreatedAt  time.Time
	FinishedAt time.Time
}

type SessionID string

func MapBehaviorSessionToRepositorySession(session bmodels.SessionModel) *Session {
	return &Session{
		SessionID:  SessionID(session.SessionID),
		UserID:     UserID(session.UserID),
		CreatedAt:  session.CreatedAt,
		FinishedAt: session.FinishedAt,
	}
}
