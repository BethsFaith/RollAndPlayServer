CREATE TABLE IF NOT EXISTS Race_Bonuses
(
    race_id INTEGER REFERENCES Races(id) ON DELETE CASCADE,
    skill_id INTEGER REFERENCES Skills(id) ON DELETE CASCADE,
    bonus INTEGER NOT NULL,

    CONSTRAINT Race_Bonuses_pkey PRIMARY KEY (race_id, skill_id)
);

CREATE TABLE IF NOT EXISTS public.Character_Class_Bonuses
(
    class_id INTEGER REFERENCES Character_Classes(id) ON DELETE CASCADE,
    skill_id INTEGER REFERENCES Skills(id) ON DELETE CASCADE,
    bonus INTEGER NOT NULL,

    CONSTRAINT Class_Bonuses_pkey PRIMARY KEY (class_id, skill_id)
);