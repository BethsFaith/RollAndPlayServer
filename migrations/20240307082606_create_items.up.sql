CREATE TABLE ItemTypes
(
    id   integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483646 CACHE 1 ),
    name varchar not null,

    user_id integer not null references users(id) on delete cascade,

    CONSTRAINT ItemTypes_pkey PRIMARY KEY (id),
    CONSTRAINT ItemTypes_ukey UNIQUE (name, user_id)
);

CREATE TABLE Items
(
    id   integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 0 MAXVALUE 2147483646 CACHE 1 ),
    name varchar not null,
    description varchar,
    icon VARCHAR(1024),
    count integer not null,

    item_type_id integer not null references ItemTypes(id) on delete cascade,
    user_id integer not null references users(id) on delete cascade,

    CONSTRAINT Items_pkey PRIMARY KEY (id),
    CONSTRAINT Items_ukey UNIQUE (name, item_type_id)
);

CREATE TABLE ItemDescriptors
(
    id   integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483646 CACHE 1 ),
    name varchar not null,

    item_type_id integer not null references ItemTypes(id) on delete cascade,

    CONSTRAINT ItemDescriptors_pkey PRIMARY KEY (id),
    CONSTRAINT ItemDescriptors_ukey UNIQUE (name, item_type_id)
);

CREATE TABLE ItemDescriptorLines
(
    id   integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483646 CACHE 1 ),

    descriptor_id integer not null references ItemDescriptors(id) on delete cascade,
    item_id integer not null references ItemTypes(id) on delete cascade,

    CONSTRAINT ItemDescriptorLines_pkey PRIMARY KEY (id),
    CONSTRAINT ItemDescriptorLines_ukey UNIQUE (descriptor_id, item_id)
);