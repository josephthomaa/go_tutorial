import time
import concurrent.futures

def sleep_for_1_second():
    time.sleep(1)
    print("Task completed.")

start_time = time.time()

with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
    futures = [executor.submit(sleep_for_1_second) for _ in range(10)]

# Wait for all tasks to complete
concurrent.futures.wait(futures)

end_time = time.time()
print(f"Program completed in {end_time - start_time} seconds.")
