package postgresql

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// getAuthorsByName - поиск по схожести
	// Input: $1 - authorName, $2 - limit, $3 - offset
	// Output: username, схожесть
	getAuthorsByName = `
SELECT 
    user_id, 
    similarity(username, $1) AS similarity_score
FROM people
WHERE role_id = (select role_id from Role where role_default_name = 'Author')
AND username % $1
ORDER BY similarity_score DESC
LIMIT $2
OFFSET $3;
`
)

// SearchAuthor - возвращает авторов по схожести
func (csr *CustomSubscriptionRepository) SearchAuthor(ctx context.Context, searchTerm string, limit, offset int) ([]string, error) {
	op := "internal.custom_subscription.repository.postgresql.SearchAuthor"

	rows, err := csr.db.Query(ctx, getAuthorsByName, searchTerm, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	authorIDs := make([]string, 0)

	var (
		authorID string
		sim      string
	)

	for rows.Next() {
		if err = rows.Scan(&authorID, &sim); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got authorName userID=%v name=%v for author=%v", authorID, sim)
		authorIDs = append(authorIDs, authorID)
	}

	return authorIDs, nil
}
