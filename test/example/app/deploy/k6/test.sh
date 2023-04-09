#!bin/bash
kubectl apply -k .
# k6 run /config/script.js --vus 1 --duration 10m --rps 10 --no-connection-reuse 



INSERT INTO public.transactions (remaining,address,user_id,created_at,updated_at,is_morning,is_afternoon,count,end_date,amount,lng,id,status,lat,upload,start_date,is_noon,menu_id) VALUES (0,'Auto Loan Account',1,'2023-04-09T02:53:09.474138Z','2023-04-09T02:53:09.474138Z',true,true,0,'',1,0,'a2f33558-aa09-4435-819b-c10fd71a0d93','pending',0,'','',true,'82f2e9b0-8e23-4170-8b16-8d00689a61c5');INSERT INTO public.transaction_amounts (id,status,amount) VALUES ('a2f33558-aa09-4435-819b-c10fd71a0d93','pending',1);