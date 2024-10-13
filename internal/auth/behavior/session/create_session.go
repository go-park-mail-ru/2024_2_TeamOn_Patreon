package session

import "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/uuid"

func CreateSession(userID string) *SessionModel {
	sessionID := uuid.GenerateUUID()
	return NewSessionModel(sessionID, userID)
}
