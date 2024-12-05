-- роли на каждый сервис
-- наследуются от бэкенда

-- будет доступ к колонке хеш_пассворд
CREATE ROLE auth_service WITH LOGIN PASSWORD 'example_password';
CREATE ROLE author_service WITH LOGIN PASSWORD 'example_password';
CREATE ROLE account_service WITH LOGIN PASSWORD 'example_password';
CREATE ROLE content_service WITH LOGIN PASSWORD 'example_password';
CREATE ROLE custom_subscription_service WITH LOGIN PASSWORD 'example_password';

GRANT backend TO auth_service;
GRANT backend TO author_service;
GRANT backend TO account_service;
GRANT backend TO content_service;
GRANT backend TO moderation_service;
GRANT backend TO custom_subscription_service;


--- Раздаем специфичные права
-- Ауф сервис КРУД для Юзеров
GRANT SELECT, INSERT, UPDATE, DELETE ON people TO auth_service;

--- Сервисы аккаунта и автора круд для всех колонок кроме хэш пассворда
GRANT SELECT(user_id, username, email, role_id),
    INSERT(user_id, username, email, role_id),
    UPDATE(user_id, username, email, role_id),
    DELETE
    ON people TO account_service, author_service;
