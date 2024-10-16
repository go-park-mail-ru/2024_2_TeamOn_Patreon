CREATE TABLE content (
    content_id uuid PRIMARY KEY,
    content_type_id uuid REFERENCES content_type ON DELETE RESTRICT, -- нельзя так
    content_url text
);
