CREATE TABLE t (
	id serial PRIMARY KEY,
	created_at TIMESTAMP NOT NULL Default now(),
    message TEXT
);

