package validate

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func ValidationAuthorName(authorName string) (string, error) {
	authorName = validate.Sanitize(authorName)
	if len(authorName) > validate.MaxLenAuthorNameForSearch {
		return "", global.ErrAuthorNameTooLong
	}
	return authorName, nil
}
