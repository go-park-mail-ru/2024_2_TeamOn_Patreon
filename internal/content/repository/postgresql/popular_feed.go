package postgresql

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	//"github.com/satori/go.uuid"
)

const (
	// getPopularPostForUserSQL возвращает посты отсортированные по лайкам по убывающей, которые
	// может смотреть пользователь среди всех постов
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков, comments
	// Input: получает $1 userId - uuid пользователя, {$2 offset} и { $3 limit}
	getPopularPostForUserSQL = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    COUNT(Like_Post.like_post_id) AS likes,
    post.created_date,
   	(SELECT COUNT(*) FROM Comment where post_id = post.post_id) as comments
FROM 
    post
JOIN 
    People AS author ON author.user_id = post.user_id 
RIGHT OUTER JOIN 
    Subscription_Layer ON Subscription_Layer.subscription_layer_id = post.subscription_layer_id
LEFT OUTER JOIN 
    Like_Post USING (post_id)
WHERE 
    (Subscription_Layer.layer <= (
        SELECT COALESCE(Subscription_Layer.layer, 0)
        FROM Subscription
        JOIN Custom_Subscription ON Subscription.custom_subscription_id = Custom_Subscription.custom_subscription_id
        JOIN Subscription_Layer ON Custom_Subscription.subscription_layer_id = Subscription_Layer.subscription_layer_id
        WHERE Custom_Subscription.author_id = author.user_id AND Subscription.user_id = $1
    )
    OR post.subscription_layer_id = (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = 0)
    OR post.user_id = $1)
	and post.post_status_id IN (select post_status_id FROM Post_Status WHERE status = 'PUBLISHED' or status = 'ALLOWED' or status = 'COMPLAINED')
GROUP BY 
    post.post_id,  
    post.About, 
    post.Title, 
    author_id, 
    author_username
ORDER BY 
    likes DESC
LIMIT $3
OFFSET $2;
	`

	// getLikedPostsForUser - возвращает id поста и информацию о том, лайкнул ли этот пост пользователь
	// на вход получает $1 userID - ид пользователя и $2 posts - список id постов
	getIsLikedPostsForUser = `
		SELECT 
			p.post_id,
			CASE WHEN lp.user_id IS NOT NULL THEN true ELSE false END AS liked
		FROM 
			Post p
		LEFT JOIN 
			Like_Post lp ON p.post_id = lp.post_id AND lp.user_id = $1
		WHERE 
			p.post_id = ANY($2)
`

	// getPopularPostsForAnonSQL возвращает посты отсортированные по лайкам по убывающей, которые
	// может смотреть пользователь среди всех постов
	// Output: postID, title, about, authorId, authorUsername, likes, created_date
	// likes - количество лайков, numCOMM
	// Input: {$1 offset} и { $2 limit}
	getPopularPostsForAnonSQL = `
SELECT 
    post.post_id, 
    post.Title, 
    post.About, 
    author.user_id AS author_id, 
    author.Username AS author_username, 
    COUNT(Like_Post.like_post_id) AS likes,
    post.created_date,
   	(SELECT COUNT(*) FROM Comment where post_id = post.post_id) as comments
FROM 
    post
JOIN 
    People AS author ON author.user_id = post.user_id 
RIGHT OUTER JOIN 
    Subscription_Layer ON Subscription_Layer.subscription_layer_id = post.subscription_layer_id
LEFT OUTER JOIN 
    Like_Post USING (post_id)
WHERE 
    post.subscription_layer_id = (SELECT subscription_layer_id FROM Subscription_Layer WHERE layer = 0)
	and post.post_status_id IN (select post_status_id FROM Post_Status WHERE status = 'PUBLISHED' or status = 'ALLOWED' or status = 'COMPLAINED')
GROUP BY 
    post.post_id,  
    post.About, 
    post.Title, 
    author_id, 
    author_username
ORDER BY 
    likes DESC
LIMIT $2
OFFSET $1;
`
)

func (cr *ContentRepository) GetPopularPostsForUser(ctx context.Context, userID string, offset int, limits int) ([]*models.Post, error) {
	op := "internal.content.repository.postgresql.GetPopularPostsForUser"

	rows, err := cr.db.Query(ctx, getPopularPostForUserSQL, userID, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         string
		title          string
		content        string
		authorId       string
		authorUsername string
		likes          int
		createdDate    time.Time
		numComments    int
	)

	posts := make([]*models.Post, 0)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &authorId, &authorUsername, &likes, &createdDate, &numComments); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorId, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID,
			Title:          title,
			Content:        content,
			AuthorID:       authorId,
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
			NumComments:    numComments,
		})

	}

	return posts, nil
}

// GetPopularPosts Выдает ленту популярных для анонимиа
func (cr *ContentRepository) GetPopularPosts(ctx context.Context, offset int, limits int) ([]*models.Post, error) {
	op := "internal.content.repository.postgresql.GetPopularPosts"

	posts := make([]*models.Post, 0)

	rows, err := cr.db.Query(ctx, getPopularPostsForAnonSQL, offset, limits)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		postID         string
		title          string
		content        string
		authorId       string
		authorUsername string
		likes          int
		createdDate    time.Time
		numComments    int
	)

	for rows.Next() {
		if err = rows.Scan(&postID, &title, &content, &authorId, &authorUsername, &likes, &createdDate, &numComments); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  post: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			postID, title, authorId, authorUsername, likes, createdDate)
		posts = append(posts, &models.Post{
			PostID:         postID,
			Title:          title,
			Content:        content,
			AuthorID:       authorId,
			AuthorUsername: authorUsername,
			Likes:          likes,
			CreatedDate:    createdDate,
			NumComments:    numComments,
		})

	}

	return posts, nil
}

func (cr *ContentRepository) GetIsLikedForPosts(ctx context.Context, UserID string, posts []*models.Post) error {
	op := "internal.content.repository.postgresql.GetIsLikedForPosts"

	postIDs := make([]string, 0, len(posts))
	postsIDMap := make(map[string]*models.Post)

	for _, post := range posts {
		postIDs = append(postIDs, post.PostID)
		postsIDMap[post.PostID] = post
	}

	rows, err := cr.db.Query(ctx, getIsLikedPostsForUser, UserID, postIDs)
	if err != nil {
		return errors.Wrap(err, op)
	}

	defer rows.Close()
	for rows.Next() {
		var (
			postId      string
			isLikedPost bool
		)
		if err = rows.Scan(&postId, &isLikedPost); err != nil {
			return errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "Got  post: post_id=%v isLiked=%v for user=%v",
			postId, isLikedPost, UserID)
		post, ok := postsIDMap[postId]
		if !ok {
			return errors.Wrap(global.ErrServer, op)
		}
		post.IsLiked = isLikedPost
	}
	return nil
}
