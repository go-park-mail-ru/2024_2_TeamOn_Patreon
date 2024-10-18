CREATE TABLE event_type (
    event_type uuid PRIMARY KEY,
    default_event_type_name text NOT NULL UNIQUE -- будут создавать миграций различные
);
