package service

import (
	"context"
	"github.com/pkg/errors"
)

func (b *Behavior) SearchAuthor(ctx context.Context, searchTerm string, limit, offset int) ([]string, error) {
	op := "custom_subscription.service.SearchAuthor"

	authorIDs, err := b.rep.SearchAuthor(ctx, searchTerm, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return authorIDs, nil
}
