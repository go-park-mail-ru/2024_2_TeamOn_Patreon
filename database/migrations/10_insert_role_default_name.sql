CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO role (role_id, role_default_name) VALUES
    (uuid_generate_v4(), 'reader'),
    (uuid_generate_v4(), 'author'),
    (uuid_generate_v4(), 'moderator');