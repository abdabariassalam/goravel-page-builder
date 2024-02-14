CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  name varchar NOT NULL,
  email varchar NOT NULL,
  password varchar NOT NULL,
  deleted_at timestamp NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
