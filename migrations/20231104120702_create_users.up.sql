CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    nickname varchar null,
    encrypted_password varchar not null
)