create table Game_System_Skills
(
    system_id         integer not null references game_systems(id) on delete cascade,
    skill_category_id integer not null references skill_categories(id) on delete cascade,
    CONSTRAINT Game_System_Skills_pkey PRIMARY KEY (system_id, skill_category_id)
);

create table Game_System_Races
(
    system_id         integer not null references game_systems(id) on delete cascade,
    race_id integer not null references races(id) on delete cascade,
    CONSTRAINT Game_System_Races_pkey PRIMARY KEY (system_id, race_id)
);

create table Game_System_Actions
(
    system_id         integer not null references game_systems(id) on delete cascade,
    action_id integer not null references actions(id) on delete cascade,
    CONSTRAINT Game_System_Actions_pkey PRIMARY KEY (system_id, action_id)
);

create table Game_System_Classes
(
    system_id         integer not null references game_systems(id) on delete cascade,
    class_id integer not null references character_classes(id) on delete cascade,
    CONSTRAINT Game_System_Classes_pkey PRIMARY KEY (system_id, class_id)
);
