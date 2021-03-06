CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);
create table todo_lists (
  id serial not NULL UNIQUE,
  title VARCHAR(255) not NULL,
  description VARCHAR(255)
);
create table users_lists (
  id serial not null UNIQUE,
  user_id int REFERENCES users (id) on delete CASCADE not NULL,
  list_id int REFERENCES todo_lists (id) on delete CASCADE not null
);
create table todo_items (
  id serial NOT NULL UNIQUE,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  done boolean not NULL DEFAULT false
);
create table lists_items (
  id serial not null UNIQUE,
  item_id int REFERENCES todo_items (id) on delete CASCADE not NULL,
  list_id int REFERENCES todo_lists (id) on delete CASCADE not null
);