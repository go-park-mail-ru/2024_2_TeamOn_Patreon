package service

import "context"

func (s *Service) ComplaintPost(ctx context.Context, postID string, userID string) error {
	// TODO: Проверяем может ли юзер видеть пост - в таком случае ошибка
	// TODO: Меняем статус поста на тот, на который пожаловались
	return nil
}
