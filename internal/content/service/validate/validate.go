package validate

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/validate"
	sanitize "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
	"github.com/pkg/errors"
)

func Post(ctx context.Context, title, content string, layer int) (string, string, int, error) {
	op := "service.validate.Post"

	err := validate.Title(title)
	if err != nil {
		return "", "", 0, errors.Wrap(err, op)
	}

	err = validate.Content(content)
	if err != nil {
		return "", "", 0, errors.Wrap(err, op)
	}

	err = validate.Layer(layer)
	if err != nil {
		return "", "", 0, errors.Wrap(err, op)
	}

	title = sanitize.Sanitize(title)
	content = sanitize.Sanitize(content)

	return title, content, layer, nil
}
