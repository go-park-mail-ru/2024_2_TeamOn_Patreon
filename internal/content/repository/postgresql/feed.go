package postgresql

const (
	// getPopularPostForUserSQL возвращает посты отсортированные по лайкам по убывающей, которые
	// может смотреть пользователь среди всех постов
	// возвращает postID, title, about, authorId, authorUsername, likes
	// likes - количество лайков
	// На вход получает authorId - uuid автора, {offset} и {limit}
	getPopularPostForUserSQL = `
		select post.post_id, post.Title, post.About, author.user_id as author_id, author.Username as author_username, COUNT(Like_Post.like_post_id) AS likes
		from post
			join People AS Author ON Author.user_id = Post.user_id 
			join Subscription_Layer on Subscription_Layer.subscription_layer_id = Post.subscription_layer_id
			left outer join Like_Post USING (post_id)
			where Subscription_Layer.layer >= (
					select Subscription_Layer.layer
					from Subscription
						join Custom_Subscription on Subscription.custom_subscription_id  = Custom_Subscription.custom_subscription_id
						join Subscription_Layer on Custom_Subscription.subscription_layer_id = Subscription_Layer.subscription_layer_id
					where Custom_Subscription.author_id = author.user_id and Subscription.user_id = {userId}
			)
			GROUP BY post.post_id,  post.About, post.Title, author_id, author_username
			ORDER BY likes DESC
			LIMIT {limit}
			OFFSET {offset}
			;
	`

	// getLikedPostsForUser - возвращает id поста и информацию о том, лайкнул ли этот пост пользователь
	// на вход получает userID - ид пользователя и posts - список id постов
	getIsLikedPostsForUser = `
		SELECT 
			p.post_id,
			CASE WHEN lp.user_id IS NOT NULL THEN true ELSE false END AS liked
		FROM 
			Post p
		LEFT JOIN 
			Like_Post lp ON p.post_id = lp.post_id AND lp.user_id = {userId}
		WHERE 
			p.post_id IN ({posts})
`
)

func (cr *ContentRepository) GetPopularPostForUser() error {
	return nil
}

func (cr *ContentRepository) GetSubscriptionsPostsForUser() error {
	return nil
}

func (cr *ContentRepository) GetPopularPostForAnon() error {
	return nil
}
