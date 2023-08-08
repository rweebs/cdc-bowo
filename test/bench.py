import time
import statistics
import psycopg2
import pandas as pd

def run_benchmark(rps, test_period, sql_query, connection_string):
    conn = psycopg2.connect(connection_string)
    cursor = conn.cursor()

    # Calculate the interval between requests based on desired RPS
    interval = 1 / rps

    # Initialize lists to hold response times
    response_times = []

    # Run the benchmark for the specified test period
    end_time = time.time() + test_period
    while time.time() < end_time:
        start_time = time.time()
        cursor.execute(sql_query)
        end_time_query = time.time()
        response_times.append(end_time_query - start_time)

        # Ensure RPS is maintained by sleeping for the required interval
        time.sleep(max(0, interval - (end_time_query - start_time)))

    conn.close()

    # Calculate statistics
    response_times.sort()
    minimum = response_times[0]
    maximum = response_times[-1]
    average = statistics.mean(response_times)
    p50 = response_times[int(len(response_times) * 0.5)]
    p75 = response_times[int(len(response_times) * 0.75)]
    p90 = response_times[int(len(response_times) * 0.90)]
    p95 = response_times[int(len(response_times) * 0.95)]
    p99 = response_times[int(len(response_times) * 0.99)]

    return {
        "Minimum": minimum,
        "Maximum": maximum,
        "Average": average,
        "50th Percentile": p50,
        "75th Percentile": p75,
        "90th Percentile": p90,
        "95th Percentile": p95,
        "99th Percentile": p99
    }

if __name__ == "__main__":
    # Replace the following values with your PostgreSQL connection details
    db_host = "your_db_host"
    db_port = "your_db_port"
    db_name = "your_db_name"
    db_user = "your_db_user"
    db_password = "your_db_password"

    # Connection string
    connection_string = f"host={db_host} port={db_port} dbname={db_name} user={db_user} password={db_password}"

    # Sample SQL query
    sql_query = "SELECT * FROM your_table"

    # Benchmark settings
    requests_per_second = 10
    test_duration_seconds = 30

    result = run_benchmark(requests_per_second, test_duration_seconds, sql_query, connection_string)

    # Create a DataFrame from the benchmark result
    df = pd.DataFrame.from_dict(result, orient='index', columns=['Response Time'])

    # Save the DataFrame to an Excel file
    output_file = "benchmark_results.xlsx"
    df.to_excel(output_file)

    print(f"Benchmark results saved to {output_file}")
