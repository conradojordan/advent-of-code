FILE_PATH = "./2024/day01/input.txt"

with open(FILE_PATH) as f:
    locations = f.readlines()

locations_a = []
locations_b = []


for location_pair in locations:
    location_a, location_b = location_pair.strip().split("   ")
    locations_a.append(int(location_a))
    locations_b.append(int(location_b))

locations_a.sort()
locations_b.sort()

sum_of_distances = 0


for i in range(len(locations_a)):
    sum_of_distances += abs(locations_a[i] - locations_b[i])

print(f"Sum of distances is: {sum_of_distances}")
