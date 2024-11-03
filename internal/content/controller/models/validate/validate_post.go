package validate

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

func Title(title string) error {
	if len(title) == 0 {
		return global.ErrFieldTooShort
	}

	if len(title) > 64 {
		return global.ErrFieldTooLong
	}

	if utils.HasForbiddenChars(title) {
		return global.ErrFieldContainsSpecialChar
	}
	return nil
}

func Content(content string) error {
	if len(content) > 64 {
		return global.ErrFieldTooLong
	}

	if utils.HasForbiddenChars(content) {
		return global.ErrFieldContainsSpecialChar
	}
	return nil
}

func Layer(layer int) error {
	if layer < 0 || layer > 3 {
		return global.ErrInvalidJSON
	}

	return nil
}
