pgbench -c 1 -R 10 -T 60 -f /scripts/insert/10.sql
sleep 20
pgbench -c 1 -R 20 -T 60 -f /scripts/insert/20.sql
sleep 20
pgbench -c 1 -R 30 -T 60 -f /scripts/insert/30.sql
sleep 20
pgbench -c 1 -R 40 -T 60 -f /scripts/insert/40.sql
sleep 20
pgbench -c 1 -R 50 -T 60 -f /scripts/insert/50.sql
sleep 20
pgbench -c 1 -R 100 -T 60 -f /scripts/insert/100.sql
sleep 20
pgbench -c 1 -R 10 -T 60 -f /scripts/delete/10.sql
sleep 20
pgbench -c 1 -R 20 -T 60 -f /scripts/delete/20.sql
sleep 20
pgbench -c 1 -R 30 -T 60 -f /scripts/delete/30.sql
sleep 20
pgbench -c 1 -R 40 -T 60 -f /scripts/delete/40.sql
sleep 20
pgbench -c 1 -R 50 -T 60 -f /scripts/delete/50.sql
sleep 20
pgbench -c 1 -R 100 -T 60 -f /scripts/delete/100.sql
sleep 20