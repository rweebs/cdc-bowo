import psycopg2
import time

conn = psycopg2.connect(
    host="localhost",
    database="log_replication",
    user="postgres"
)

cur = conn.cursor()

while True:
    cur.execute("INSERT INTO film_actor values (132,81,now())")
    conn.commit()
conn.close()
