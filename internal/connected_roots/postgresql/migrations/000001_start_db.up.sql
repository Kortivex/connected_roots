CREATE TABLE users
(
    "id"         varchar(26) primary key,
    "name"       varchar(100),
    "surname"    varchar(100),
    "email"      varchar(255) unique not null,
    "password"   varchar(255),
    "telephone"  varchar(30),
    "language"   varchar(3),
    "role_id"    varchar(26),
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone
);

CREATE TABLE roles
(
    "id"          varchar(26) primary key,
    "name"        varchar(50),
    "description" varchar(255),
    "protected"   boolean default false,
    "created_at"  timestamp with time zone,
    "updated_at"  timestamp with time zone,
    "deleted_at"  timestamp with time zone
);

CREATE TABLE crop_types
(
    "id"              varchar(26) primary key,
    "name"            varchar(100),
    "scientific_name" varchar(100),
    "life_cycle"      varchar(100),
    "planting_season" varchar(100),
    "harvest_season"  varchar(100),
    "irrigation"      varchar(100),
    "image_url"       text,
    "description"     text,
    "created_at"      timestamp with time zone,
    "updated_at"      timestamp with time zone,
    "deleted_at"      timestamp with time zone
);

CREATE TABLE orchards
(
    "id"           varchar(26) primary key,
    "name"         varchar(100),
    "location"     varchar(255),
    "size"         decimal,
    "soil"         varchar(255),
    "fertilizer"   varchar(255),
    "composting"   varchar(255),
    "image_url"    text,
    "user_id"      varchar(100),
    "crop_type_id" varchar(26),
    "created_at"   timestamp with time zone,
    "updated_at"   timestamp with time zone,
    "deleted_at"   timestamp with time zone
);

CREATE TABLE sensors
(
    "id"                     varchar(26) primary key,
    "name"                   varchar(255),
    "type"                   varchar(255),
    "location"               varchar(255),
    "model_number"           varchar(255),
    "manufacturer"           varchar(255),
    "measurement_range"      varchar(255),
    "calibration_date"       timestamp with time zone,
    "battery_life"           decimal,
    "communication_protocol" varchar(255),
    "status"                 int,
    "firmware_version"       decimal,
    "high_threshold"         decimal,
    "low_threshold"          decimal,
    "orchard_id"             varchar(26),
    "created_at"             timestamp with time zone,
    "updated_at"             timestamp with time zone,
    "deleted_at"             timestamp with time zone
);

CREATE TABLE sensor_data
(
    "id"         varchar(26) primary key,
    "sensor_id"  varchar(26),
    "value"      decimal(10, 2),
    "timestamp"  timestamp with time zone,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone
);

CREATE TABLE agricultural_activities
(
    "id"          varchar(26) primary key,
    "name"        varchar(100),
    "description" text,
    "orchard_id"  varchar(26),
    "created_at"  timestamp with time zone,
    "updated_at"  timestamp with time zone,
    "deleted_at"  timestamp with time zone
);

CREATE TABLE reports
(
    "id"           varchar(26) primary key,
    "name"         varchar(100),
    "description"  text,
    "content"      text,
    "generated_by" varchar(26),
    "generated_at" timestamp with time zone,
    "created_at"   timestamp with time zone,
    "updated_at"   timestamp with time zone,
    "deleted_at"   timestamp with time zone
);

CREATE TABLE sessions
(
    "id"         text not null primary key,
    "data"       text,
    "created_at" timestamp with time zone,
    "expires_at" timestamp with time zone
);

ALTER TABLE users
    ADD FOREIGN KEY ("role_id") REFERENCES roles ("id");

ALTER TABLE orchards
    ADD FOREIGN KEY ("user_id") REFERENCES users ("id");

ALTER TABLE orchards
    ADD FOREIGN KEY ("crop_type_id") REFERENCES crop_types ("id");

ALTER TABLE sensors
    ADD FOREIGN KEY ("orchard_id") REFERENCES orchards ("id");

ALTER TABLE sensor_data
    ADD FOREIGN KEY ("sensor_id") REFERENCES sensors ("id");

ALTER TABLE agricultural_activities
    ADD FOREIGN KEY ("orchard_id") REFERENCES orchards ("id");

ALTER TABLE reports
    ADD FOREIGN KEY ("generated_by") REFERENCES users ("id");
