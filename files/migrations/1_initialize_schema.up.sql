CREATE TABLE IF NOT EXISTS transactions (
    id              SERIAL         UNIQUE PRIMARY KEY,
    name            VARCHAR        NOT NULL,
    tag             VARCHAR        NULL
);
