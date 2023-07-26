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
ALTER TABLE ONLY public.transaction REPLICA IDENTITY FULL;
CREATE PUBLICATION my_publication FOR ALL TABLES WITH (publish = 'insert, update, delete, truncate');


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



    CREATE TABLE IF NOT EXISTS public.transaction_unified (
    id SERIAL PRIMARY KEY NOT NULL,
    text text NOT NULL,
    created timestamp without time zone DEFAULT now(),
    created_cdc timestamp without time zone DEFAULT now(),
    rps integer,
    op text
);

ALTER TABLE ONLY public.transaction REPLICA IDENTITY FULL;

CREATE OR REPLACE FUNCTION insert_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO transaction_unified (id, text, created_cdc,rps, created, op) VALUES (NEW.id, NEW.text, NEW.created_cdc, NEW.rps, NEW.created, 'c');
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER insert_transaction
AFTER INSERT ON transaction
FOR EACH ROW
EXECUTE FUNCTION insert_transaction_function();


---- CREATE OR REPLACE TRIGGER after before delete to move the record to t3
CREATE OR REPLACE FUNCTION archive_transaction_function()
RETURNS trigger AS $$
BEGIN
    INSERT INTO transaction_unified (id, text, created_cdc, created,rps) VALUES (OLD.id, OLD.text, OLD.created_cdc, OLD.created, OLD.rps);
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
        DROP TRIGGER IF EXISTS archive_transaction ON transaction;
        DROP TABLE IF EXISTS public.transaction CASCADE;
        DROP TABLE IF EXISTS public.transaction_unified CASCADE;
    SQL
  }

  migration {
    id   = "blue"
    up   = <<SQL
    CREATE SUBSCRIPTION my_subscription CONNECTION 'host=${local.primary_endpoint_address} port=5432 password=CuTGUoIA user=postgres dbname=postgres' PUBLICATION my_publication;
    SQL
    down = <<SQL
    DROP SUBSCRIPTION IF EXISTS my_subscription;
    SQL
  }
}
