import os

file_path = os.path.dirname(__file__) + "/input.txt"

increase_count = 0

with open(file_path) as f:
    measurements = [int(x) for x in f.readlines()]

for i in range(3, len(measurements)):
    a, b, c, d = (
        measurements[i - 3],
        measurements[i - 2],
        measurements[i - 1],
        measurements[i],
    )
    if b + c + d > a + b + c:
        increase_count += 1


print(f"The number of times the measurement increased is {increase_count}")
