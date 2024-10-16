CREATE TABLE people (
      user_id uuid PRIMARY KEY,
      username text,
      email   text,
      role_id uuid REFERENCES role ON DELETE RESTRICT, -- нельзя так
      hash_password text
);
