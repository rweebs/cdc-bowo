import pandas as pd
import psycopg2
import os

dest_host = os.getenv('POSTGRES_HOST')
# Replace these values with your actual PostgreSQL connection details
db_config = {
    'host': dest_host,
    'port': '5432',
    'database': 'postgres',
    'user': 'postgres',
    'password': 'CuTGUoIA',
}
# Your SQL query
sql_query = """
SELECT
    op,
    rps,
    AVG(EXTRACT(EPOCH FROM replication_lag) * 1000) AS avg_replication_lag_ms,
    MIN(EXTRACT(EPOCH FROM replication_lag) * 1000) AS min_replication_lag_ms,
    MAX(EXTRACT(EPOCH FROM replication_lag) * 1000) AS max_replication_lag_ms,
    STDDEV(EXTRACT(EPOCH FROM replication_lag) * 1000) AS stddev_replication_lag_ms,
    PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY EXTRACT(EPOCH FROM replication_lag) * 1000) AS percentile_50_ms,
    PERCENTILE_CONT(0.75) WITHIN GROUP (ORDER BY EXTRACT(EPOCH FROM replication_lag) * 1000) AS percentile_75_ms,
    PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY EXTRACT(EPOCH FROM replication_lag) * 1000) AS percentile_90_ms,
    PERCENTILE_CONT(0.95) WITHIN GROUP (ORDER BY EXTRACT(EPOCH FROM replication_lag) * 1000) AS percentile_95_ms,
    PERCENTILE_CONT(0.99) WITHIN GROUP (ORDER BY EXTRACT(EPOCH FROM replication_lag) * 1000) AS percentile_99_ms
FROM
    transaction_unified
GROUP BY
    op,
    rps
ORDER BY
    op,
    rps;
"""

# Connect to the PostgreSQL database
conn = psycopg2.connect(**db_config)

# Execute the query and read the results into a pandas DataFrame
df = pd.read_sql_query(sql_query, conn)

# Close the database connection
conn.close()

# Save the DataFrame to an Excel file
excel_file = './result/output_data.xlsx'
df.to_excel(excel_file, index=False)

print(f"Data has been saved to '{excel_file}'.")