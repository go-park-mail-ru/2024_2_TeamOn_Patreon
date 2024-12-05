CREATE ROLE backend;

--- Возможность подключения к нашей бд
GRANT CONNECT ON DATABASE pushart TO backend;

--- Выдаем круд на все таблицу в схеме кроме People
GRANT SELECT, INSERT, UPDATE, DELETE
      ON role, page, avatar, tip, subscription_layer,
      subscription, post_status, custom_subscription,
      content_type, post, content, comment, like_post,
      like_comment, event_type, event, notification
      TO backend;

--- Выдаем SELECT на все столбцы кроме hash_password
GRANT SELECT(user_id, username, email, role_id) ON people TO backend;

