CREATE TABLE content_type (
    content_type_id uuid PRIMARY KEY,
    default_content_type_name text NOT NULL UNIQUE      -- различные типы контента (будут задаваться в следующих миграциях
);
