import os
import math

file_path = os.path.dirname(__file__) + "/input.txt"

last_measurement = math.inf
increase_count = 0

with open(file_path) as f:
    for line in f:
        if int(line) > last_measurement:
            increase_count += 1
        last_measurement = int(line)

print(f"The number of times the measurement increased is {increase_count}")
