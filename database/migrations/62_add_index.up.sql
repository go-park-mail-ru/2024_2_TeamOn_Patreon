CREATE INDEX idx_post_subscription_layer_status ON post(subscription_layer_id, post_status_id);

CREATE INDEX idx_post_likes_created_date ON post(post_id, created_date);

CREATE INDEX idx_like_post_post_id ON Like_Post(post_id);

CREATE INDEX idx_comment_post_id ON Comment(post_id);

CREATE INDEX idx_subscription_layer_layer ON Subscription_Layer(layer);

CREATE INDEX idx_subscription_user_id_custom_id ON Subscription(user_id, custom_subscription_id);
CREATE INDEX idx_custom_subscription_author_layer ON Custom_Subscription(author_id, subscription_layer_id);

