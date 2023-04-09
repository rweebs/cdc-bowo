truncate transactions CASCADE;

-- alter table transactions add column cdc_timestamp timestamp without time zone DEFAULT now();
-- alter table transactions add column rps int;
-- alter table menus add column cdc_timestamp timestamp without time zone DEFAULT now();
-- alter table menus add column rps int;
-- alter table day_menus add column cdc_timestamp timestamp without time zone DEFAULT now();
-- alter table day_menus add column rps int;
truncate menus CASCADE;
truncate day_menus CASCADE;
truncate transactions CASCADE;
drop table transaction_amounts CASCADE;
CREATE SUBSCRIPTION helper CONNECTION 'host=host.docker.internal port=5434 password=postgres user=postgres dbname=helper' PUBLICATION helper;