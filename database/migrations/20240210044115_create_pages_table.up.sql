CREATE TABLE pages (
  id uuid PRIMARY KEY NOT NULL,
  title varchar(100) NOT NULL,
  short_description varchar NULL,
  template_type varchar(10) NULL,
  content text NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL
);
