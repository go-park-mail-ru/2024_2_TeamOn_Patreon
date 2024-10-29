-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE role (
      role_id uuid PRIMARY KEY,
      role_default_name text NOT NULL
);

INSERT INTO role (role_id, role_default_name) VALUES
    (uuid_generate_v4(), 'reader'),
    (uuid_generate_v4(), 'author');