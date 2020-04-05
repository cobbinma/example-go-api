CREATE TABLE IF NOT EXISTS transactions (
    identifier      SERIAL         UNIQUE PRIMARY KEY,
    name            VARCHAR        NOT NULL,
    animal          VARCHAR        NOT NULL
);
