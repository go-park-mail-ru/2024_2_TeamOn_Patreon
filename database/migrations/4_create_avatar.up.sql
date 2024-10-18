CREATE TABLE avatar (
    avatar_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL, -- если есть аватарка, значит, есть юзер
    avatar_url text             -- аватарка, может не быть
);
