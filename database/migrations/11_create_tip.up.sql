CREATE TABLE tip (
    tip_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,      -- не может быть пожертвования без пользователя
    author_id uuid REFERENCES people (user_id) ON DELETE CASCADE NOT NULL , -- не может быть пожертвования без автора
    cost integer NOT NULL CHECK ( cost > 0 ),  -- нельзя пожертвовать ноль рублей
    message text,           --  можно без описания
    payed_date timestamp NOT NULL DEFAULT NOW() -- время должно быть
);
