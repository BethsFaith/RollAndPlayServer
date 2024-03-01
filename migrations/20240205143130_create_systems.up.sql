CREATE TABLE Game_Systems
(
    id   integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
    name varchar not null unique,
    icon VARCHAR(1024),
    user_id integer not null references users(id) on delete cascade,

    CONSTRAINT Game_Systems_pkey PRIMARY KEY (id)
)