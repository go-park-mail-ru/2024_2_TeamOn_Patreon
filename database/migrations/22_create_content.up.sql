CREATE TABLE content (
    content_id uuid PRIMARY KEY,
    content_type_id uuid REFERENCES content_type ON DELETE RESTRICT NOT NULL, -- контент всегда с типом
    content_url text NOT NULL -- контент должен быть
);
