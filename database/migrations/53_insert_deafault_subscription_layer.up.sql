-- Вставка уровней подписки
INSERT INTO Subscription_Layer (subscription_layer_id, layer, default_layer_name) VALUES
    (gen_random_uuid(), 0, 'Default Layer'),
    (gen_random_uuid(), 1, 'Basic Layer'),
    (gen_random_uuid(), 2, 'Premium Layer'),
    (gen_random_uuid(), 3, 'VIP Layer');
