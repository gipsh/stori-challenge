

-- +migrate Up


CREATE TABLE transactions (
 	id INTEGER PRIMARY KEY,
 	month INTEGER,
 	day INTEGER,
 	amount REAL,
 	account TEXT
);

-- +migrate Down

DROP TABLE transactions;