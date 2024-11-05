CREATE TABLE post (
    post_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL, -- пользователь создает пост
    title text NOT NULL,    -- нельзя создавать пост без названия
    about text,             -- можно создать пост без описания
    subscription_layer_id uuid REFERENCES subscription_layer ON DELETE RESTRICT NOT NULL, -- нельзя создать пост без уровня
    -- по умолчанию будет задаваться уровень доступный вссем
    -- будет в следующих миграция
    created_date timestamp NOT NULL DEFAULT now(),
    updated_date timestamp -- может быть пустым - не обновляли
);
