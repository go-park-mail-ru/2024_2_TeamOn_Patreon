CREATE TABLE custom_subscription (
    custom_subscription_id uuid PRIMARY KEY,
    author_id uuid REFERENCES people (user_id) ON DELETE CASCADE NOT NULL, -- без автора нет подписки
    custom_name text NOT NULL,  -- имя кастомной подписки должно быть
    cost integer NOT NULL,      -- стоимость должна быть
    info text,                  -- описание кастомной подписки, может быть нулем
    subscription_layer_id uuid REFERENCES subscription_layer ON DELETE RESTRICT NOT NULL, -- уровень подписки должен быть
    created_date timestamp NOT NULL DEFAULT now(),
    UNIQUE (author_id, custom_name),
    UNIQUE (author_id, subscription_layer_id)
);
