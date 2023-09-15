CREATE TYPE STATUS AS ENUM ('New', 'Success', 'Failure', 'Error', 'Cancel');

CREATE TABLE IF NOT EXISTS transactions
(
    id         SERIAL PRIMARY KEY,
    user_id    INT       NOT NULL,
    user_email TEXT      NOT NULL,
    amount     DECIMAL   NOT NULL,
    currency   TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    status     STATUS    NOT NULL
);