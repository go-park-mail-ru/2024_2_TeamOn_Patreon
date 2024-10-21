CREATE TABLE page (
    page_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE, -- если удалили пользователя, удаляем страницу
    info text,                      -- описание, может отсутствовать
    background_picture_url text     -- картинка на заднем фоне, может не быть
);
