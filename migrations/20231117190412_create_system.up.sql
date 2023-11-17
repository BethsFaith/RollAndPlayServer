CREATE TABLE IF NOT EXISTS Systems
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(1024),

    CONSTRAINT System_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS System_Races
(
    system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
    race_id INTEGER REFERENCES Races(id) ON DELETE CASCADE,

    CONSTRAINT System_Races_pkey PRIMARY KEY (system_id, race_id)
);

CREATE TABLE IF NOT EXISTS System_Classes
(
    system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
    class_id INTEGER REFERENCES Character_Classes(id) ON DELETE CASCADE,

    CONSTRAINT System_Classes_pkey PRIMARY KEY (system_id, class_id)
);

CREATE TABLE IF NOT EXISTS System_Skills
(
    system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
    skill_category_id INTEGER REFERENCES Skill_categories(id) ON DELETE CASCADE,

    CONSTRAINT System_Skills_pkey PRIMARY KEY (system_id, skill_category_id)
);


