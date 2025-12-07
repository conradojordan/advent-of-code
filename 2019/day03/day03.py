from sys import exit as sys_exit


def calculate_points(wire):
    x, y = 0, 0
    wire_instructions = wire.split(",")
    wire_points = set()
    for instruction in wire_instructions:
        if instruction[0] == "R":
            for _ in range(1, int(instruction[1:]) + 1):
                x += 1
                wire_points.add((x, y))
        elif instruction[0] == "L":
            for _ in range(1, int(instruction[1:]) + 1):
                x -= 1
                wire_points.add((x, y))
        elif instruction[0] == "U":
            for _ in range(1, int(instruction[1:]) + 1):
                y += 1
                wire_points.add((x, y))
        else:
            for _ in range(1, int(instruction[1:]) + 1):
                y -= 1
                wire_points.add((x, y))
    return wire_points


def distance(point_a, point_b):
    return abs(point_b[0] - point_a[0]) + abs(point_b[1] - point_a[1])


def closest_cross_to_center(wires):
    wire1, wire2 = wires.strip().split()
    origin = (0, 0)
    cross_points = []
    wire1_points, wire2_points = calculate_points(wire1), calculate_points(wire2)
    for point in wire1_points:
        if point in wire2_points:
            cross_points.append(point)
    min_distance = float("inf")
    for cross_point in cross_points:
        point_distance = distance(cross_point, origin)
        if point_distance < min_distance:
            min_distance = point_distance
            closest_cross = cross_point
    return closest_cross, min_distance, cross_points


def cross_points_steps(cross_points, wire):
    x, y = 0, 0
    steps = 0
    wire_instructions = wire.split(",")
    cross_points_steps = {}
    for instruction in wire_instructions:
        if instruction[0] == "R":
            for _ in range(1, int(instruction[1:]) + 1):
                x += 1
                steps += 1
                if (x, y) in cross_points:
                    cross_points_steps[(x, y)] = steps
        elif instruction[0] == "L":
            for _ in range(1, int(instruction[1:]) + 1):
                x -= 1
                steps += 1
                if (x, y) in cross_points:
                    cross_points_steps[(x, y)] = steps
        elif instruction[0] == "U":
            for _ in range(1, int(instruction[1:]) + 1):
                y += 1
                steps += 1
                if (x, y) in cross_points:
                    cross_points_steps[(x, y)] = steps
        else:
            for _ in range(1, int(instruction[1:]) + 1):
                y -= 1
                steps += 1
                if (x, y) in cross_points:
                    cross_points_steps[(x, y)] = steps
    return cross_points_steps


def fewest_combined_steps(wires, cross_points):
    wire1, wire2 = wires.strip().split()
    wire1_crosses_and_steps = cross_points_steps(cross_points, wire1)
    wire2_crosses_and_steps = cross_points_steps(cross_points, wire2)
    min_steps = float("inf")
    for cross_point in cross_points:
        current_steps = (
            wire1_crosses_and_steps[cross_point] + wire2_crosses_and_steps[cross_point]
        )
        if current_steps < min_steps:
            min_steps = current_steps
    return min_steps


if __name__ == "__main__":
    ## Read file
    try:
        with open("2019/day03/input.txt") as file:
            wires = file.read()
    except FileNotFoundError:
        sys_exit(
            "File not found! Make sure you have 'day3_input.txt' in the 'inputs' directory."
        )

    ## Part 1
    print("Part 1".center(40, "-"))
    closest_cross, min_distance, cross_points = closest_cross_to_center(wires)
    print(
        f"Closest cross point is at {closest_cross} with distance {min_distance} from the origin."
    )

    ## Part 2
    print("Part 2".center(40, "-"))
    fewest_steps = fewest_combined_steps(wires, cross_points)
    print(f"Cross point with fewest combined steps has {fewest_steps} steps total.")
