alter table transaction alter column message type varchar(255);
alter table transaction rename column text to message;

PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/10.sql -c 1 -R 10 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/20.sql -c 1 -R 20 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/30.sql -c 1 -R 30 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/40.sql -c 1 -R 40 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/50.sql -c 1 -R 50 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/insert/100.sql -c 1 -R 100 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/10.sql -c 1 -R 10 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/20.sql -c 1 -R 20 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/30.sql -c 1 -R 30 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/40.sql -c 1 -R 40 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/50.sql -c 1 -R 50 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/update/100.sql -c 1 -R 100 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/10.sql -c 1 -R 10 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/20.sql -c 1 -R 20 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/30.sql -c 1 -R 30 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/40.sql -c 1 -R 40 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/50.sql -c 1 -R 50 -T 60
sleep 20
PGPASSWORD=$PASSWORD pgbench -h $HOST -U postgres -d postgres -f /scripts/delete/100.sql -c 1 -R 100 -T 60