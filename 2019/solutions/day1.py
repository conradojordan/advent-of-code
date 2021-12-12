from math import floor
from sys import exit as sys_exit


def calculate_fuel(module_weight):
    if module_weight < 9:
        return 0
    else:
        return module_weight // 3 - 2


def calculate_fuel_recursive(module_weight):
    if module_weight < 9:
        return 0
    else:
        fuel = module_weight // 3 - 2
        return fuel + calculate_fuel_recursive(fuel)


if __name__ == "__main__":
    ## Read file
    try:
        with open("../inputs/day1_input.txt") as file:
            lines = list(file)
    except FileNotFoundError:
        sys_exit(
            "File not found! Make sure you have 'day1_input.txt' in the 'inputs' directory."
        )

    ## Part 1
    total_fuel = 0
    for line in lines:
        total_fuel += calculate_fuel(int(line))
    print(total_fuel)

    ## Part 2
    total_recursive_fuel = 0
    for line in lines:
        total_recursive_fuel += calculate_fuel_recursive(int(line))
    print(total_recursive_fuel)

