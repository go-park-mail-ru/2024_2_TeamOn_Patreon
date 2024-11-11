package validate

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func Title(title string) error {
	if len(title) < validate.MinLenTitle {
		return global.ErrFieldTooShort
	}

	if len(title) > validate.MaxLenTitle {
		return global.ErrFieldTooLong
	}

	return nil
}

func Content(content string) error {
	if len(content) > validate.MaxLenContent {
		return global.ErrFieldTooLong
	}

	return nil
}

func Layer(layer int) error {
	if layer < validate.MinLayer || layer > validate.MaxLayer {
		return global.ErrInvalidJSON
	}

	return nil
}
