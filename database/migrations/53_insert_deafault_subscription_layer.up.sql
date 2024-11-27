-- Вставка уровней подписки
INSERT INTO Subscription_Layer (subscription_layer_id, layer, default_layer_name) VALUES
    (gen_random_uuid(), 0, 'Нулевой уровень'),
    (gen_random_uuid(), 1, 'Базовая подписка'),
    (gen_random_uuid(), 2, 'Премиум подписка'),
    (gen_random_uuid(), 3, 'VIP подписка');
