import psycopg2
import time

conn = psycopg2.connect(
    host="localhost",
    database="log_replication",
    user="postgres"
)

cur = conn.cursor()

while True:
    cur.execute("SELECT sent_lsn,write_lsn,flush_lsn,replay_lsn FROM pg_stat_replication where application_name='mysub' ;")
    # Fetch all the rows from the SELECT statement
    row = cur.fetchone()
    # print(row)
    # Iterate over the rows and print each row
    print(row)
    if row[0]==row[1]==row[2]==row[3]:
        print("All LSNs are equal")
    else:
        print("LSNs are not equal")
    time.sleep(1)


    # Close the cursor and connection
cur.close()
conn.close()


