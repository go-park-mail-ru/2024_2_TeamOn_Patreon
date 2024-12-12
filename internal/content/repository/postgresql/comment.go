package postgresql

import "context"

func (cr *ContentRepository) CreateComment(ctx context.Context, userID, commentID string, content string) error {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) UpdateComment(ctx context.Context, userID, commentID string, content string) error {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) DeleteComment(ctx context.Context, userID, commentID string) error {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) GetUserIDByCommentID(ctx context.Context, commentID string) (string, error) {
	//TODO implement me
	panic("implement me")
}
