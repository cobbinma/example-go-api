CREATE TABLE IF NOT EXISTS pets (
    id              INTEGER        UNIQUE PRIMARY KEY,
    name            VARCHAR        NOT NULL,
    tag             VARCHAR        NULL
);
