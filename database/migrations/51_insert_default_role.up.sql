-- Вставка ролей
INSERT INTO Role (role_id, role_default_name) VALUES
    (gen_random_uuid(), 'Reader'),
    (gen_random_uuid(), 'Author'),
    (gen_random_uuid(), 'Moderator');
