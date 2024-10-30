package imagine

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"sync"
)

type ContentRepository struct {
	posts    map[uuid.UUID]*Post
	mu       sync.RWMutex
	authors  map[uuid.UUID]*Author
	keysPost []uuid.UUID
}

var rep *ContentRepository

func New() *ContentRepository {
	op := "content.repository.imagine.repository.ContentRepository"
	if rep == nil {
		logger.StandardDebugF(op, "New repository")
		rep = &ContentRepository{}
		rep.posts = make(map[uuid.UUID]*Post)
		rep.authors = make(map[uuid.UUID]*Author)
		rep.keysPost = make([]uuid.UUID, 0)
		return rep
	}
	return rep
}

func (pr *ContentRepository) InsertPost(userId uuid.UUID, postId uuid.UUID, title string, content string, layer int) error {
	op := "content.repository.imagine.repository.InsertPost"

	logger.StandardDebugF(op, "post.PostId=%v", postId)
	post := Post{postId: postId, title: title, content: content, layer: layer, authorID: userId}

	logger.StandardDebugF(op, "post.PostId=%v post=%v", post.postId, post)
	pr.mu.Lock()
	defer pr.mu.Unlock()
	pr.posts[postId] = &post
	pr.keysPost = append(pr.keysPost, postId)
	return nil
}

func (pr *ContentRepository) GetPopularPosts(offset int, limits int) ([]models.Post, error) {
	op := "content.repository.imagine.repository.GetPopularPostsForUser"
	logger.StandardDebugF(op, "offset=%v limits=%v", offset, limits)

	pr.mu.Lock()
	defer pr.mu.Unlock()

	if offset >= limits || offset >= len(pr.keysPost) {
		posts := make([]models.Post, 0)
		return posts, nil
	}

	top := len(pr.keysPost)
	if limits < top {
		top = limits
	}
	logger.StandardDebugF(op, "offset=%v, top=%v keys=%v", offset, top, pr.keysPost)

	var popularPosts []models.Post

	slicePosts := pr.keysPost[offset:top]
	logger.StandardDebugF(op, "slicePosts=%v", slicePosts)

	for _, postId := range slicePosts {
		logger.StandardDebugF(op, "post.PostId=%v", postId)
		repPost, ok := pr.posts[postId]
		if !ok {
			logger.StandardDebugF(op, "post.PostId=%v", postId)
			return nil, errors.Wrap(global.ErrServer, op)
		}
		logger.StandardDebugF(op, "repPost = %v", repPost)
		pkgPost := MapRepositoryPostToPkgPost(*repPost)

		_, ok = pr.authors[postId]
		if !ok {
			//return nil, errors.Wrap(global.ErrServer, op)
		}
		if ok {
			pkgPost.AuthorUsername = pr.authors[postId].AuthorUsername
		}

		popularPosts = append(popularPosts, pkgPost)

	}

	return popularPosts, nil
}

func (r *ContentRepository) GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error) {
	op := "content.repository.imagine.repository.GetAuthorByPost"

	r.mu.RLock()
	defer r.mu.RUnlock()
	post, ok := r.posts[postID]
	if ok != true {
		return postID, errors.Wrap(global.ErrPostDoesntExists, op)
	}

	return post.authorID, nil
}

func (r *ContentRepository) UpdatePost(authorId uuid.UUID, postId uuid.UUID, updatePost models.Post) error {
	op := "content.repository.imagine.repository.UpdatePost"

	logger.StandardDebugF(op, "post.PostId=%v", updatePost.PostId)
	r.mu.Lock()
	defer r.mu.Unlock()
	post, ok := r.posts[postId]
	if !ok {
		logger.StandardDebugF(op, "post.PostId=%v", updatePost.PostId)
		return global.ErrPostDoesntExists
	}

	if updatePost.Title != "" {
		post.title = updatePost.Title
	}
	if updatePost.Content != "" {
		post.content = updatePost.Content
	}
	if updatePost.Layer != 0 {
		post.layer = updatePost.Layer
	}
	return nil
}

// LikePost

func (cr *ContentRepository) IsLikePutPost(userId uuid.UUID, postID uuid.UUID) (bool, error) {
	op := "content.repository.imagine.repository.IsLikePutPost"
	logger.StandardDebugF(op, "post.PostId=%v", postID)
	return false, nil
}

func (cr *ContentRepository) InsertLikePost(userId uuid.UUID, postID uuid.UUID) error {
	op := "content.repository.imagine.repository.InsertLikePost"
	logger.StandardDebugF(op, "post.PostId=%v", postID)
	return nil
}

func (cr *ContentRepository) DeleteLikePost(userId uuid.UUID, postID uuid.UUID) error {
	op := "content.repository.imagine.repository.DeleteLikePost"
	logger.StandardDebugF(op, "post.PostId=%v", postID)
	return nil
}

func (cr *ContentRepository) GetPostLikes(postID uuid.UUID) (int, error) {
	op := "content.repository.imagine.repository.GetPostLikes"
	logger.StandardDebugF(op, "post.PostId=%v", postID)
	return 0, nil
}

func (cr *ContentRepository) GetPopularPostsForLayer(userId uuid.UUID, offset int, limit int) ([]models.Post, error) {
	/*
		потенциальный sql:
			// сначала выборка по уровням для авторов для этого пользователя


			select post.postId, post.Title, post.Content, author.userId, author.Username
			from post
				join people on people.userId == post.authorId as author
				join layer on layer.layerId == post.layerId
			where layer.layer >= (select layer.layer
								from layer
									join CustomSubscription on CustomSubscription.subscription_layer_id == layer.layer_id
									join Subscription on Subscription.custom_subscription_id  == CustomSubscription.subscription_id
								where Subscription.user_id == {userId} and CustomSubscription.author_id == author.userId
								)
			order by post.postId, post.Title, post.Content, author.userId, author.Username
			sort by likes (??????)
			limit {limits}
			top {offset}
	*/
	return nil, nil
}
