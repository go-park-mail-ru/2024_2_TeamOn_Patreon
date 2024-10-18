CREATE TABLE subscription_layer (
    subscription_layer_id uuid PRIMARY KEY,
    layer integer UNIQUE NOT NULL,   -- уровень подписки инт, нужно для сравнения что больше что меньше
    default_layer_name text NOT NULL -- имя уровня подписки
);
