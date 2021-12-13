import os
import math

FILE_PATH = os.path.join(
    os.path.dirname(os.path.realpath(__file__)), "..", "inputs", "day1.txt"
)

last_measurement = math.inf
increased_count = 0

with open(FILE_PATH) as f:
    measurements = [int(x) for x in f.readlines()]

for measurement in measurements:
    if measurement > last_measurement:
        increased_count += 1
    last_measurement = measurement

print(" Part 1 ".center(50, "-"))
print(f"The number of times the measurement increased is {increased_count}\n")

increased_count = 0

for i in range(3, len(measurements)):
    a, b, c, d = (
        measurements[i - 3],
        measurements[i - 2],
        measurements[i - 1],
        measurements[i],
    )
    if b + c + d > a + b + c:
        increased_count += 1

print(" Part 2 ".center(50, "-"))
print(f"The number of times the measurement increased is {increased_count}")
