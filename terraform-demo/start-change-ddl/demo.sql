CREATE TABLE IF EXIST public.transaction_amounts (
    id uuid DEFAULT public.uuid_generate_v4(),
    status text,
    amount bigint
);
ALTER TABLE public.transaction_amounts REPLICA IDENTITY FULL;
--
INSERT INTO public.transaction_amounts SELECT id, status, amount FROM public.transactions;