from sys import exit as sys_exit


def run_intcode_part1(intcode):
    intcode = [int(integer) for integer in intcode.split(",")]
    i = 0

    while True:
        instruction = str(intcode[i]).zfill(5)
        param_modes = instruction[:3]
        opmode = instruction[3:]

        if opmode == "99":
            break
        elif opmode == "01":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            intcode[intcode[i + 3]] = param1 + param2
            increment = 4
        elif opmode == "02":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            intcode[intcode[i + 3]] = param1 * param2
            increment = 4
        elif opmode == "03":
            intcode[intcode[i + 1]] = int(input("Enter input: "))
            increment = 2
        else:
            if param_modes[2] == "0":
                print(intcode[intcode[i + 1]])
            else:
                print(intcode[i + 1])
            increment = 2
        i += increment

    return intcode


def run_intcode_part2(intcode):
    intcode = [int(integer) for integer in intcode.split(",")]
    i = 0

    while True:
        instruction = str(intcode[i]).zfill(5)
        param_modes = instruction[:3]
        opmode = instruction[3:]

        if opmode == "99":
            break
        elif opmode == "01":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            intcode[intcode[i + 3]] = param1 + param2
            increment = 4
        elif opmode == "02":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            intcode[intcode[i + 3]] = param1 * param2
            increment = 4
        elif opmode == "03":
            intcode[intcode[i + 1]] = int(input("Enter input: "))
            increment = 2
        elif opmode == "04":
            if param_modes[2] == "0":
                print(intcode[intcode[i + 1]])
            else:
                print(intcode[i + 1])
            increment = 2
        elif opmode == "05":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            if param1 != 0:
                i = param2
                increment = 0
            else:
                increment = 3
        elif opmode == "06":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            if param1 == 0:
                i = param2
                increment = 0
            else:
                increment = 3
        elif opmode == "07":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            if param1 < param2:
                intcode[intcode[i + 3]] = 1
            else:
                intcode[intcode[i + 3]] = 0
            increment = 4
        elif opmode == "08":
            param1 = (
                intcode[intcode[i + 1]] if param_modes[2] == "0" else intcode[i + 1]
            )
            param2 = (
                intcode[intcode[i + 2]] if param_modes[1] == "0" else intcode[i + 2]
            )
            if param1 == param2:
                intcode[intcode[i + 3]] = 1
            else:
                intcode[intcode[i + 3]] = 0
            increment = 4
        i += increment

    return intcode


if __name__ == "__main__":
    ## Read file
    try:
        with open("2019/inputs/day5_input.txt") as file:
            intcode_original = file.read()
    except FileNotFoundError:
        sys_exit(
            "File not found! Make sure you have 'day5_input.txt' in the 'inputs' directory."
        )

    ## Part 1
    print("Part 1".center(40, "="))
    output_intcode_part1 = run_intcode_part1(intcode_original)

    ## Part 2
    print("Part 2".center(40, "="))
    output_intcode_part2 = run_intcode_part2(intcode_original)
