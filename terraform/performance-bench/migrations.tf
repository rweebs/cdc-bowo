resource "sql_migrate" "blue" {
  provider = sql.blue
  migration {
    id = "blue"
    up = <<SQL
    CREATE TABLE IF NOT EXISTS public.transaction (
    id SERIAL PRIMARY KEY NOT NULL,
    text text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    rps integer,
    op text
);

    CREATE TABLE IF NOT EXISTS public.transaction_delete_primary (
    id SERIAL PRIMARY KEY NOT NULL,
    text text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    created_cdc timestamp without time zone DEFAULT now(),
    rps integer,
    op text default 'd'
);
ALTER TABLE ONLY public.transaction_delete_primary REPLICA IDENTITY FULL;
ALTER TABLE ONLY public.transaction REPLICA IDENTITY FULL;
CREATE OR REPLACE FUNCTION archive_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO transaction_delete_primary (id, text, created,rps) VALUES (OLD.id, OLD.text, NOW(), old.rps);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER  archive_transaction
AFTER DELETE ON transaction
FOR EACH ROW
EXECUTE FUNCTION archive_transaction_function();

    SQL

    down = <<SQL

DROP TABLE IF EXISTS public.transaction CASCADE;
    SQL
  }

}



resource "sql_migrate" "green" {
  depends_on = [sql_migrate.blue]
  provider   = sql.green
  migration {
    id = "blue"
    up = <<SQL
    CREATE TABLE IF NOT EXISTS public.transaction (
    id SERIAL PRIMARY KEY NOT NULL,
    text text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    created_cdc timestamp without time zone DEFAULT now(),
    rps integer,
    op text default 'c'
);

    CREATE TABLE IF NOT EXISTS public.transaction_delete (
    id SERIAL PRIMARY KEY NOT NULL,
    message text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    created_cdc timestamp without time zone DEFAULT now(),
    rps integer,
    op text default 'd'
);


    CREATE TABLE IF NOT EXISTS public.transaction_unified (
    id SERIAL PRIMARY KEY NOT NULL,
    message text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    created_cdc timestamp without time zone DEFAULT now(),
    rps integer,
    op text
);

ALTER TABLE ONLY public.transaction REPLICA IDENTITY FULL;

CREATE OR REPLACE FUNCTION insert_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO transaction_unified (id, message, created_cdc,rps, created, op) VALUES (NEW.id, NEW.message, NEW.created_cdc, NEW.rps, NEW.created, 'c');
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER insert_transaction
AFTER INSERT ON transaction
FOR EACH ROW
EXECUTE FUNCTION insert_transaction_function();

CREATE OR REPLACE FUNCTION update_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO public.transaction_unified (id, message, created,rps, op) VALUES (NEW.id, NEW.message, NEW.created, NEW.rps, 'u');
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_transaction
AFTER UPDATE ON transaction
FOR EACH ROW
EXECUTE FUNCTION update_transaction_function();

---- CREATE OR REPLACE TRIGGER after before delete to move the record to t3
CREATE OR REPLACE FUNCTION archive_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO transaction_delete (id, message, created_cdc,rps) VALUES (OLD.id, OLD.message, NOW(), OLD.rps);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER archive_transaction
AFTER DELETE ON transaction
FOR EACH ROW
EXECUTE FUNCTION archive_transaction_function();

    SQL

    down = <<SQL
DROP TRIGGER IF EXISTS insert_transaction ON transaction;
DROP TRIGGER IF EXISTS update_transaction ON transaction;
DROP TRIGGER IF EXISTS archive_transaction ON transaction;
DROP TABLE IF EXISTS public.transaction CASCADE;
DROP TABLE IF EXISTS public.transaction CASCADE;
DROP TABLE IF EXISTS public.transaction_delete CASCADE;
    SQL
  }
}
