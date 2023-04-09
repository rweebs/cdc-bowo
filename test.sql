-- update record in table t2
-- insert into public.t2(text) values (md5(random()::text));
-- delete from public.t3 where id in (select id from public.t3 limit 1);
update t3 set created=now() where id in (select value from update_temp limit 1);
update update_temp set value = value + 1;
-- update public.t3 set text 


-- -- create trigger after before delete to move the record to t3
-- CREATE FUNCTION archive_t2_function()
-- RETURNS trigger AS $$
-- BEGIN
--     INSERT INTO t2_archieve (id, message, created) VALUES (OLD.id, OLD.message, NOW());
--     RETURN NULL;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER archive_t2
-- AFTER DELETE ON t2
-- FOR EACH ROW
-- EXECUTE FUNCTION archive_t2_function();

-- CREATE OR REPLACE FUNCTION archive_t3_function()
-- RETURNS trigger AS $$
-- BEGIN
--     INSERT INTO public.t3_archieve (id, text, created) VALUES (OLD.id, OLD.text, NOW());
--     RETURN NULL;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER archive_t3
-- AFTER DELETE ON t3
-- FOR EACH ROW
-- EXECUTE FUNCTION archive_t3_function();