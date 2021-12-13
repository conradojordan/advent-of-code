from sys import exit as sys_exit


def parse_orbits(orbits_map):
    orbits_map = orbits_map.strip().split()
    orbits = dict()
    for orbit_pair in orbits_map:
        planet, satellite = orbit_pair.split(")")
        orbits[satellite] = planet
    return orbits


def count_orbits(orbits_map):
    orbits = parse_orbits(orbits_map)

    total_orbits = 0
    for satellite in orbits.keys():
        current_satellite = satellite
        while current_satellite != "COM":
            total_orbits += 1
            current_satellite = orbits[current_satellite]

    return total_orbits


def list_orbits(orbits, satellite):
    orbits_list = list()
    current = satellite
    while current != "COM":
        current = orbits[current]
        orbits_list.append(current)
    return orbits_list


def distance_you_san(orbits_map):
    orbits = parse_orbits(orbits_map)
    you_orbits = list_orbits(orbits, "YOU")
    san_orbits = list_orbits(orbits, "SAN")
    for satellite in you_orbits:
        if satellite in san_orbits:
            first_common_satellite = satellite
            break
    common_steps = len(list_orbits(orbits, first_common_satellite)) + 1
    return (len(you_orbits) - common_steps) + (len(san_orbits) - common_steps)


if __name__ == "__main__":
    ## Read file
    try:
        with open("2019/inputs/day6_input.txt") as file:
            orbits_map = file.read()
    except FileNotFoundError:
        sys_exit(
            "File not found! Make sure you have 'day6_input.txt' in the 'inputs' directory."
        )

    ## Part 1
    print(f"Total orbits: {count_orbits(orbits_map)}")

    ## Part 2
    print(f"Total steps from you to santa: {distance_you_san(orbits_map)}")
