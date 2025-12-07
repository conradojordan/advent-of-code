from sys import exit as sys_exit


def listify(input_string):
    return [int(integer) for integer in input_string.split(",")]


def run_intcode(intcode):
    i = 0
    while intcode[i] != 99:
        if intcode[i] == 1:
            intcode[intcode[i + 3]] = intcode[intcode[i + 2]] + intcode[intcode[i + 1]]
        else:
            intcode[intcode[i + 3]] = intcode[intcode[i + 2]] * intcode[intcode[i + 1]]
        i += 4
    return intcode


def prepare_input(intcode, noun, verb):
    intcode[1] = noun
    intcode[2] = verb
    return intcode


def find_inputs_for(target):
    for noun in range(100):
        for verb in range(100):
            current_intcode = prepare_input([x for x in original_intcode], noun, verb)
            output = run_intcode(current_intcode)[0]
            if output > target:
                break
            elif output == target:
                return noun, verb


if __name__ == "__main__":
    ## Read file
    try:
        with open("2019/day02/input.txt") as file:
            original_intcode = file.read()
    except FileNotFoundError:
        sys_exit(
            "File not found! Make sure you have 'day2_input.txt' in the 'inputs' directory."
        )

    ## Part 1
    original_intcode = listify(original_intcode)
    prepared_intcode = prepare_input([x for x in original_intcode], 12, 2)
    output_intcode = run_intcode(prepared_intcode)
    print("Part 1".center(40, "-"))
    print(f"Program output after corrections: {output_intcode[0]}")

    ## Part 2
    target = 19690720
    target_noun, target_verb = find_inputs_for(target)
    print("Part 2".center(40, "-"))
    print(f"Target noun = {target_noun} and target verb = {target_verb}")
    print(f"{target_noun}*100 + {target_verb} = {target_noun * 100 + target_verb}")
