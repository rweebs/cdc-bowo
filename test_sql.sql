ALTER TABLE public.t RENAME COLUMN message to c;
ALTER TABLE public.t ALTER COLUMN c type varchar(11);
ALTER TABLE public.t3 RENAME TO t2;
ALTER TABLE public.t2 RENAME COLUMN text to message;
ALTER TABLE public.t2 DROP COLUMN created_cdc;
CREATE TABLE IF NOT EXISTS public.create_table (id Serial Primary Key ,message text ,created_at timestamp);
CREATE TABLE IF NOT EXISTS public.create_table_after_2022 (id Serial Primary Key ,message text ,created_at timestamp);
CREATE TABLE IF NOT EXISTS public.vertical_splitting (id Serial Primary Key ,message text ,message2 text);
CREATE TABLE IF NOT EXISTS public.vertical_splitting_derived_1 (id Serial Primary Key ,message text);
CREATE TABLE IF NOT EXISTS public.vertical_splitting_derived_2 (id Serial Primary Key ,message2 text);
CREATE TABLE IF NOT EXISTS public.vertical_splitting_2 (id Serial Primary Key ,message text ,message2 text);
CREATE TABLE IF NOT EXISTS public.vertical_splitting_2_derived_1 (id Serial Primary Key ,message text);
CREATE TABLE IF NOT EXISTS public.vertical_splitting_2_derived_2 (id Serial Primary Key ,message2 text);
-- Special case add column in the middle of the table from table add_column_1
CREATE TABLE IF NOT EXISTS public.add_column_middle (id Serial Primary Key ,add_column int, message text);
INSERT INTO public.add_column_middle (id, message) SELECT id, message FROM public.add_column_1;
DROP TABLE public.add_column_1;
ALTER TABLE public.add_column_middle RENAME TO add_column_1;