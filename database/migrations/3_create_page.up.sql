CREATE TABLE page (
    page_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE, -- если удалили пользователя, удаляем страницу
    info text,
    background_picture_url text
);
