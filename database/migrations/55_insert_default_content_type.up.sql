-- Вставка типов контента
INSERT INTO Content_Type (content_type_id, default_content_type_name) VALUES
    (gen_random_uuid(), 'jpeg'),
    (gen_random_uuid(), 'png'),
    (gen_random_uuid(), 'mp3'),
    (gen_random_uuid(), 'mp4'),
    (gen_random_uuid(), 'pdf');