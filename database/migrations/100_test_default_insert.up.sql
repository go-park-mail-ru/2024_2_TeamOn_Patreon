-- -- Вставка ролей
-- INSERT INTO Role (role_id, role_default_name) VALUES
--     (gen_random_uuid(), 'Reader'),
--     (gen_random_uuid(), 'Author');

-- Вставка пользователей
INSERT INTO People (user_id, username, email, role_id, hash_password) VALUES
    (gen_random_uuid(), 'user1', 'user1@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),    -- password abo!D5fdffba
    (gen_random_uuid(), 'user2', 'user2@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Author'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user3', 'user3@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user4', 'user4@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Author'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user5', 'user5@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user6', 'user6@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Author'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user7', 'user7@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user8', 'user8@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Author'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user9', 'user9@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Reader'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa'),
    (gen_random_uuid(), 'user10', 'user10@example.com', (SELECT role_id FROM Role WHERE role_default_name = 'Author'), '$2a$10$H1xs9gC/rgIb6TRu23UmJOJHhepSEoJspCNziKbsVa5DguCG8vUHa');

-- Cоздание авторов
UPDATE public.people
    SET role_id=(select Role.role_id FROM Role WHERE role_default_name='Author')
    WHERE public.people.username in (('user1'), ('user2'), ('user3'));

-- Вставка страниц
INSERT INTO Page (page_id, user_id, info, background_picture_url) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), 'Info about User 1', 'http://example.com/bg1.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), 'Info about User 2', 'http://example.com/bg2.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), 'Info about User 3', 'http://example.com/bg3.jpg');


-- Вставка аватаров
INSERT INTO Avatar (avatar_id, user_id, avatar_url) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), 'http://example.com/avatar1.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), 'http://example.com/avatar2.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), 'http://example.com/avatar3.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user4'), 'http://example.com/avatar4.jpg'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user5'), 'http://example.com/avatar5.jpg');

-- -- Вставка уровней подписки
-- INSERT INTO Subscription_Layer (subscription_layer_id, layer, default_layer_name) VALUES
--     (gen_random_uuid(), 0, 'Default Layer'),
--     (gen_random_uuid(), 1, 'Basic Layer'),
--     (gen_random_uuid(), 2, 'Premium Layer'),
--     (gen_random_uuid(), 3, 'VIP Layer');

-- Вставка кастомных подписок
INSERT INTO Custom_Subscription (custom_subscription_id, author_id, custom_name, cost, info, subscription_layer_id, created_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), 'Custom Sub 1', 10, 'Custom Subscription Info 1', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Basic Layer'), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), 'Custom Sub 2', 20, 'Custom Subscription Info 2', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Premium Layer'), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), 'Custom Sub 3', 30, 'Custom Subscription Info 3', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'VIP Layer'), NOW());

-- Вставка подписок
INSERT INTO Subscription (subscription_id, user_id, custom_subscription_id, started_date, finished_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), (SELECT custom_subscription_id FROM Custom_Subscription WHERE custom_name = 'Custom Sub 1'), NOW(), NOW() + INTERVAL '30 days'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user4'), (SELECT custom_subscription_id FROM Custom_Subscription WHERE custom_name = 'Custom Sub 2'), NOW(), NOW() + INTERVAL '60 days'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user5'), (SELECT custom_subscription_id FROM Custom_Subscription WHERE custom_name = 'Custom Sub 3'), NOW(), NOW() + INTERVAL '90 days');

-- Вставка постов
INSERT INTO Post (post_id, user_id, title, about, subscription_layer_id, created_date, updated_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), 'First Post', 'This is the first post.', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Basic Layer'), NOW(), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), 'Second Post', 'This is the second post.', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Premium Layer'), NOW(), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), 'Third Post', 'This is the third post.', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'VIP Layer'), NOW(), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user4'), 'Fourth Post', 'This is the fourth post.', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Basic Layer'), NOW(), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user5'), 'Fifth Post', 'This is the fifth post.', (SELECT subscription_layer_id FROM Subscription_Layer WHERE default_layer_name = 'Premium Layer'), NOW(), NOW());

-- Вставка типов контента
INSERT INTO Content_Type (content_type_id, default_content_type_name) VALUES
    (gen_random_uuid(), 'Video'),
    (gen_random_uuid(), 'Image'),
    (gen_random_uuid(), 'Audio'),
    (gen_random_uuid(), 'Text'),
    (gen_random_uuid(), 'Document');

-- Вставка контента
INSERT INTO Content (content_id, post_id, content_type_id, content_url) VALUES
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT content_type_id FROM Content_Type WHERE default_content_type_name = 'Video'), 'http://example.com/content1.mp4'),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT content_type_id FROM Content_Type WHERE default_content_type_name = 'Image'), 'http://example.com/content2.mp4'),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT content_type_id FROM Content_Type WHERE default_content_type_name = 'Audio'), 'http://example.com/content3.mp4'),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT content_type_id FROM Content_Type WHERE default_content_type_name = 'Text'), 'http://example.com/content4.mp4'),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT content_type_id FROM Content_Type WHERE default_content_type_name = 'Document'), 'http://example.com/content5.mp4');

-- Вставка комментариев
INSERT INTO Comment (comment_id, post_id, user_id, about, created_date, updated_date) VALUES
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT user_id FROM People WHERE username = 'user1'), 'Great post!', NOW(), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT user_id FROM People WHERE username = 'user2'), 'Interesting insights.', NOW(), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT user_id FROM People WHERE username = 'user3'), 'I learned something new!', NOW(), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT user_id FROM People WHERE username = 'user4'), 'Thank you for sharing!', NOW(), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Third Post'), (SELECT user_id FROM People WHERE username = 'user5'), 'Looking forward to more posts.', NOW(), NOW());

-- Вставка лайков к постам
INSERT INTO Like_Post (like_post_id, post_id, user_id, posted_date) VALUES
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT user_id FROM People WHERE username = 'user1'), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'First Post'), (SELECT user_id FROM People WHERE username = 'user2'), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Second Post'), (SELECT user_id FROM People WHERE username = 'user3'), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Third Post'), (SELECT user_id FROM People WHERE username = 'user4'), NOW()),
    (gen_random_uuid(), (SELECT post_id FROM Post WHERE title = 'Third Post'), (SELECT user_id FROM People WHERE username = 'user5'), NOW());

INSERT INTO Like_Comment (like_comment_id, comment_id, user_id, posted_date) VALUES
    (gen_random_uuid(), (SELECT comment_id FROM Comment WHERE about = 'Great post!'), (SELECT user_id FROM People WHERE username = 'user1'), NOW()),
    (gen_random_uuid(), (SELECT comment_id FROM Comment WHERE about = 'Interesting insights.'), (SELECT user_id FROM People WHERE username = 'user2'), NOW()),
    (gen_random_uuid(), (SELECT comment_id FROM Comment WHERE about = 'I learned something new!'), (SELECT user_id FROM People WHERE username = 'user3'), NOW()),
    (gen_random_uuid(), (SELECT comment_id FROM Comment WHERE about = 'Thank you for sharing!'), (SELECT user_id FROM People WHERE username = 'user4'), NOW()),
    (gen_random_uuid(), (SELECT comment_id FROM Comment WHERE about = 'Looking forward to more posts.'), (SELECT user_id FROM People WHERE username = 'user5'), NOW());

-- Вставка типов событий
INSERT INTO Event_Type (event_type, default_event_type_name) VALUES
    (gen_random_uuid(), 'Post Created'),
    (gen_random_uuid(), 'Comment Added'),
    (gen_random_uuid(), 'Post Liked'),
    (gen_random_uuid(), 'Subscription Created'),
    (gen_random_uuid(), 'Tip Sent');

-- Вставка событий
INSERT INTO Event (event_id, event_type_id, about, happened_date) VALUES
    (gen_random_uuid(), (SELECT event_type FROM Event_Type WHERE default_event_type_name = 'Post Created'), 'User 1 made a post.', NOW()),
    (gen_random_uuid(), (SELECT event_type FROM Event_Type WHERE default_event_type_name = 'Post Created'), 'User 2 commented on a post.', NOW()),
    (gen_random_uuid(), (SELECT event_type FROM Event_Type WHERE default_event_type_name = 'Post Liked'), 'User 3 liked a post.', NOW()),
    (gen_random_uuid(), (SELECT event_type FROM Event_Type WHERE default_event_type_name = 'Post Liked'), 'User 4 created a subscription.', NOW()),
    (gen_random_uuid(), (SELECT event_type FROM Event_Type WHERE default_event_type_name = 'Post Liked'), 'User 5 sent a tip.', NOW());

-- Вставка уведомлений
INSERT INTO Notification (notification_id, user_id, event_id, sent_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), (SELECT event_id FROM Event LIMIT 1 ), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), (SELECT event_id FROM Event LIMIT 1 ), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), (SELECT event_id FROM Event LIMIT 1 ), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user4'), (SELECT event_id FROM Event LIMIT 1 ), NOW()),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user5'), (SELECT event_id FROM Event LIMIT 1 ), NOW());

-- Вставка сессий
INSERT INTO Session (session_id, user_id, created_date, finished_date) VALUES
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user1'), NOW(), NOW() + INTERVAL '2 hours'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user2'), NOW(), NOW() + INTERVAL '3 hours'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user3'), NOW(), NOW() + INTERVAL '1 hour'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user4'), NOW(), NOW() + INTERVAL '4 hours'),
    (gen_random_uuid(), (SELECT user_id FROM People WHERE username = 'user5'), NOW(), NOW() + INTERVAL '5 hours');
