CREATE TABLE public.rename_column (
    id SERIAL PRIMARY KEY,
    text TEXT
);

CREATE TABLE public.add_column (
    id SERIAL PRIMARY KEY,
    add_column INT,
    add_column_non_null INT NOT NULL DEFAULT 'nonNull'
);

CREATE TABLE public.drop_column (
    id SERIAL PRIMARY KEY
);

CREATE TABLE public.drop_column_non_null (
    id SERIAL PRIMARY KEY
);

CREATE TABLE public.rename_table_new (
    id SERIAL PRIMARY KEY
);

CREATE TABLE public.test_temporal (
    id SERIAL PRIMARY KEY,
    time2 TIME,
    timestamp2 TIMESTAMP (3),
    timestamp_with_time_zone2 TIMESTAMPTZ
);
