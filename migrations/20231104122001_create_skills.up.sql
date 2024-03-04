CREATE TABLE Skill_Categories
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483646 CACHE 1 ),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(1024),
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE cascade,

    CONSTRAINT Skills_Categories_pkey PRIMARY KEY (id)
);

CREATE TABLE Skills
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483646 CACHE 1 ),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(1024),
    category_id INTEGER NULL REFERENCES Skill_Categories(id) ON DELETE set null,
    characteristic_id INTEGER NULL REFERENCES users(id) ON DELETE set null,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE cascade,

    CONSTRAINT Skills_pkey PRIMARY KEY (id),
    CONSTRAINT Skills_ukey UNIQUE (name, category_id)
);

