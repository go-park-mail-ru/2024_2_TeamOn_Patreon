-- Вставка возможных статусов поста
INSERT INTO Post_Status (post_status_id, status) VALUES
    (gen_random_uuid(), 'COMPLAINED'),
    (gen_random_uuid(), 'PUBLISHED'),
    (gen_random_uuid(), 'ALLOWED'),
    (gen_random_uuid(), 'BLOCKED');