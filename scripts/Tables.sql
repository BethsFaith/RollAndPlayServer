DROP TABLE IF EXISTS public.Race_Bonuses;
DROP TABLE IF EXISTS public.Character_Class_Bonuses;
DROP TABLE IF EXISTS public.System_Skills;
DROP TABLE IF EXISTS public.System_Classes;
DROP TABLE IF EXISTS public.System_Races;
DROP TABLE IF EXISTS public.System;
DROP TABLE IF EXISTS public.Actions;
DROP TABLE IF EXISTS public.Skills;
DROP TABLE IF EXISTS public.Skill_Categories;
DROP TABLE IF EXISTS public.Character_Classes;
DROP TABLE IF EXISTS public.Races;

CREATE TABLE IF NOT EXISTS public.Skill_Categories
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(1024),
    CONSTRAINT Skills_Categories_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.Skills
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(1024),
    category_id INTEGER REFERENCES Skill_Categories(id) ON DELETE SET NULL,

    CONSTRAINT Skills_pkey PRIMARY KEY (id),
    CONSTRAINT Skills_ukey UNIQUE (name, category_id)
);

CREATE TABLE IF NOT EXISTS public.Races 
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
	name VARCHAR(255) NOT NULL,
	model VARCHAR(1024),
	CONSTRAINT Races_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.Actions
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
	name VARCHAR(255) NOT NULL, 
	skill_id INTEGER,
	points INTEGER NOT NULL,
	
	CONSTRAINT Actions_pkey PRIMARY KEY (id),
	FOREIGN KEY (skill_id) REFERENCES Skills(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS public.Races
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
	name VARCHAR(255) NOT NULL,
	model VARCHAR(1024),
	
	CONSTRAINT Races_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.Character_Classes
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
	name VARCHAR(255) NOT NULL,
	icon VARCHAR(1024),
	
	CONSTRAINT Character_Classes_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.Race_Bonuses
(
	race_id INTEGER REFERENCES Races(id) ON DELETE CASCADE,
	skill_id INTEGER REFERENCES Skills(id) ON DELETE CASCADE,
	level_bonus INTEGER NOT NULL,
	
	CONSTRAINT Race_Bonuses_pkey PRIMARY KEY (race_id, skill_id)
);

CREATE TABLE IF NOT EXISTS public.Character_Class_Bonuses 
(
	class_id INTEGER REFERENCES Character_Classes(id) ON DELETE CASCADE,
	skill_id INTEGER REFERENCES Skills(id) ON DELETE CASCADE,
	level_bonus INTEGER NOT NULL,
	
	CONSTRAINT Class_Bonuses_pkey PRIMARY KEY (class_id, skill_id)
);

CREATE TABLE IF NOT EXISTS public.Systems
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
	name VARCHAR(255) NOT NULL,
	
	CONSTRAINT System_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.System_Races
(
	system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
	race_id INTEGER REFERENCES Races(id) ON DELETE CASCADE,
	
	CONSTRAINT System_Races_pkey PRIMARY KEY (system_id, race_id)
);

CREATE TABLE IF NOT EXISTS public.System_Classes
(
	system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
	class_id INTEGER REFERENCES Character_Classes(id) ON DELETE CASCADE,
	
	CONSTRAINT System_Classes_pkey PRIMARY KEY (system_id, class_id)
);

CREATE TABLE IF NOT EXISTS public.System_Skills
(
	system_id INTEGER REFERENCES Systems(id) ON DELETE CASCADE,
	skill_id INTEGER REFERENCES Skills(id) ON DELETE CASCADE,
	
	CONSTRAINT System_Skills_pkey PRIMARY KEY (system_id, skill_id)
);

ALTER TABLE IF EXISTS public.Skill_Categories
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.Skills
    OWNER to postgres;
ALTER TABLE IF EXISTS public.Races
    OWNER to postgres;
ALTER TABLE IF EXISTS public.Character_Classes
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.Actions
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.Race_Bonuses
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.Character_Class_Bonuses
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.Systems
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.System_Races
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.System_Classes
    OWNER to postgres;	
ALTER TABLE IF EXISTS public.System_Skills
    OWNER to postgres;	