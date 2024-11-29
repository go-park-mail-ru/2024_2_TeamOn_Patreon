CREATE TABLE status_post (
    post_status_id uuid PRIMARY KEY,
    status text NOT NULL,    --  нельзя создать статус без названия
);
