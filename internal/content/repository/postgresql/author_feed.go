package postgresql

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

const (
	// getAuthorPostsForAnon возвращает посты одного автора отсортированные по дате по убывающей, которые
	// может смотреть пользователь среди всех постов
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков
	// Input: $1 - authorId ид автора, {$2 offset} и { $3 limit}
	getAuthorPostsForAnon = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    COUNT(Like_Post.like_post_id) AS likes,
    post.created_date
FROM 
    post
JOIN 
	    People AS author ON author.user_id = post.user_id 
	RIGHT OUTER JOIN 
	    Subscription_Layer ON Subscription_Layer.subscription_layer_id = post.subscription_layer_id
	LEFT OUTER JOIN 
	    Like_Post USING (post_id)
WHERE 
    post.user_id = $1
	and post.subscription_layer_id = (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = 0)
GROUP BY 
    post.post_id,  
    post.About, 
    post.Title, 
    author_id, 
    author_username
ORDER BY 
    created_date DESC
LIMIT $3
OFFSET $2;
`

	// getAuthorPostsForMe возвращает автору его посты, отсортированные по дате по убывающей
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков
	// Input: $1 - authorId {$2 offset} и { $3 limit}
	getAuthorPostsForMe = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    COUNT(Like_Post.like_post_id) AS likes,
    post.created_date
FROM 
    post
JOIN 
	People AS author ON author.user_id = post.user_id 
LEFT OUTER JOIN 
	Like_Post USING (post_id)
WHERE
    post.user_id = $1
GROUP BY 
    post.post_id,  
    post.About, 
    post.Title, 
    author_id, 
    author_username
ORDER BY 
	created_date DESC
LIMIT $3
OFFSET $2;
`

	// getAuthorPostsForLayerSQL возвращает посты автора на определенном уровне, отсортированные по дате по убывающей
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков
	// Input: $1 - layer (int) $2 - authorId {$3 offset} и { $4 limit}
	getAuthorPostsForLayerSQL = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    COUNT(Like_Post.like_post_id) AS likes,
    post.created_date
FROM 
    post
JOIN 
    People AS author ON author.user_id = post.user_id 
RIGHT OUTER JOIN 
    Subscription_Layer ON Subscription_Layer.subscription_layer_id = post.subscription_layer_id
LEFT OUTER JOIN 
    Like_Post USING (post_id)
WHERE 
    (Subscription_Layer.layer <= $1
     OR post.subscription_layer_id = (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = 0) )
    AND post.user_id = $2
GROUP BY 
    post.post_id,  
    post.About, 
    post.Title, 
    author_id, 
    author_username
ORDER BY 
    created_date DESC
LIMIT $4
OFFSET $3;
`
)

// GetAuthorPostsForMe - возвращает посты автора для самого автора с offset по (offset + limit)
func (cr *ContentRepository) GetAuthorPostsForMe(ctx context.Context, authorID string, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPostsForMe"

	posts := make([]*models.Post, 0)

	rows, err := cr.db.Query(ctx, getAuthorPostsForMe, authorID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         uuid.UUID
		title          string
		content        string
		_authorId      uuid.UUID // Не хочется стандартный код менять...
		authorUsername string
		likes          int
		createdDate    time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &_authorId, &authorUsername, &likes, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorID, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID.String(),
			Title:          title,
			Content:        content,
			AuthorID:       _authorId.String(),
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
		})

	}

	return posts, nil
}

// GetAuthorPostsForLayer - подписки автора, которые может смотреть пользователь
func (cr *ContentRepository) GetAuthorPostsForLayer(ctx context.Context, layer int, authorID string, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.GetAuthorPostsForLayer"
	posts := make([]*models.Post, 0)

	rows, err := cr.db.Query(ctx, getAuthorPostsForLayerSQL, layer, authorID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         string
		title          string
		content        string
		_authorId      string // не хочется стандартный код менять...
		authorUsername string
		likes          int
		createdDate    time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &_authorId, &authorUsername, &likes, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorID, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID,
			Title:          title,
			Content:        content,
			AuthorID:       _authorId,
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
		})

	}

	return posts, nil
}

func (cr *ContentRepository) GetAuthorPostsForAnon(ctx context.Context, authorID string, offset, limit int) ([]*models.Post, error) {
	op := "internal.content.repository.author_feed.GetAuthorPostsForAnon"

	posts := make([]*models.Post, 0)

	rows, err := cr.db.Query(ctx, getAuthorPostsForAnon, authorID, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         uuid.UUID
		title          string
		content        string
		_authorId      uuid.UUID // не хочется стандартный код менять...
		authorUsername string
		likes          int
		createdDate    time.Time
	)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &_authorId, &authorUsername, &likes, &createdDate); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorID, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID.String(),
			Title:          title,
			Content:        content,
			AuthorID:       _authorId.String(),
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
		})

	}

	return posts, nil
}
