insert into t(message) values('hello world');

CREATE TABLE test_numeric (
    id SERIAL PRIMARY KEY,
    numeric NUMERIC(5,2),
    decimal decimal,
    real real,
    double_precision double precision
);

CREATE TABLE test_temporal (
    id SERIAL PRIMARY KEY,
    time Time without time zone Default now(),
    time_mili Time(3) without time zone Default now(),
    time_with_time_zone Time with time zone Default now(),
    timestamp Timestamp without time zone Default now(),
    timestamp_mili Timestamp(3) without time zone Default now(),
    timestamp_with_time_zone Timestamp with time zone Default now(),
    date Date Default now()

);

insert into test_numeric (numeric, decimal, real, double_precision) values (1.1, 1.1, 1.1, 1.1);

insert into test_temporal(