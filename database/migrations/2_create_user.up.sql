CREATE TABLE people (
      user_id uuid PRIMARY KEY,
      username text UNIQUE NOT NULL, -- имя пользователя должно быть
      email   text UNIQUE,           -- нельзя две почты на одном акке, почты может не быть
      role_id uuid REFERENCES role ON DELETE RESTRICT NOT NULL, -- у каждого пользователя роль, по дефолту будет читатель
      -- миграция с установлением дефолта для роли будет позже
      hash_password text NOT NULL -- нельзя без пароля
);
