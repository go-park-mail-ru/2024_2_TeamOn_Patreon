CREATE TABLE event_type (
    -- исправить на ID!!!
    event_type uuid PRIMARY KEY,
    default_event_type_name text NOT NULL UNIQUE -- будут создавать миграций различные
);
