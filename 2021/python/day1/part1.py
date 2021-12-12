import os
import math

file_path = os.path.dirname(__file__) + "/input.txt"

last_measurement = math.inf
increased_count = 0

with open(file_path) as f:
    for line in f:
        if int(line) > last_measurement:
            increased_count += 1
        last_measurement = int(line)

print(f"The number of times the measurement increased is {increased_count}")
