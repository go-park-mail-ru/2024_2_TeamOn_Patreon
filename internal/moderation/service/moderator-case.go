package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

// DecisionPost проставляет новый статус по решению модератора
func (s *Service) DecisionPost(ctx context.Context, postID string, userID string, status string) error {
	// TODO: Проверяем является ли юзер модератором
	// TODO: Проверяем у поста один из тех статусов, что должен валидировать модератор?
	// TODO: Проверяем валидный ли новый статус
	// TODO: Меняем статус на новый
	return nil
}

// GetPostsForModeration - возвращает посты в зависимости от фильтра для модератора
// В фильтре может быть одно из допустимых значений для модерации - BLOCKED | COMPLAINED
func (s *Service) GetPostsForModeration(ctx context.Context, userID string, filter string, limit, offset int) ([]*models.Post, error) {
	// TODO: Проверка является ли юзер модератором
	// TODO: Проверка валидный ли фильтр
	// TODO: Валидация лимит оффсет
	// TODO: Достаем посты по фильтру из репозитория
	return nil, nil
}
