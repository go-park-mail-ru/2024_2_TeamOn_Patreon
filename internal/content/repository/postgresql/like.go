package postgresql

// IsPostLikePut - узнает проставлен ли лайк
func (cr *ContentRepository) IsPostLikePut() (bool, error) {
	return false, nil
}

// InsertPostLike - проставляет like
func (cr *ContentRepository) InsertPostLike() error {
	return nil
}

// DeletePostLike - удаляет лайк
func (cr *ContentRepository) DeletePostLike() error {
	return nil
}
