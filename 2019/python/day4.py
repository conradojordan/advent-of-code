import time


def number_of_passwords_1(lower_limit, upper_limit):
    password_count = 0
    for d1 in range(1, 6):
        for d2 in range(d1, 10):
            for d3 in range(d2, 10):
                for d4 in range(d3, 10):
                    for d5 in range(d4, 10):
                        for d6 in range(d5, 10):
                            if d1 == d2 or d2 == d3 or d3 == d4 or d4 == d5 or d5 == d6:
                                int_number = int(
                                    str(d1)
                                    + str(d2)
                                    + str(d3)
                                    + str(d4)
                                    + str(d5)
                                    + str(d6)
                                )
                                if (
                                    int_number >= lower_limit
                                    and int_number <= upper_limit
                                ):
                                    password_count += 1
    return password_count


def number_of_passwords_2(lower_limit, upper_limit):
    password_count = 0
    for d1 in range(1, 6):
        for d2 in range(d1, 10):
            for d3 in range(d2, 10):
                for d4 in range(d3, 10):
                    for d5 in range(d4, 10):
                        for d6 in range(d5, 10):
                            if (
                                (d1 == d2 and d2 != d3)
                                or (d2 == d3 and d1 != d2 and d3 != d4)
                                or (d3 == d4 and d2 != d3 and d4 != d5)
                                or (d4 == d5 and d3 != d4 and d5 != d6)
                                or (d5 == d6 and d4 != d5)
                            ):
                                int_number = int(
                                    str(d1)
                                    + str(d2)
                                    + str(d3)
                                    + str(d4)
                                    + str(d5)
                                    + str(d6)
                                )
                                if (
                                    int_number >= lower_limit
                                    and int_number <= upper_limit
                                ):
                                    password_count += 1
    return password_count


if __name__ == "__main__":
    ## Part 1
    print("Part 1".center(40, "-"))
    print(f"Total number of passwords is: {number_of_passwords_1(136760, 595730)}")

    ## Part 2
    print("Part 2".center(40, "-"))
    print(f"Total number of passwords is: {number_of_passwords_2(136760, 595730)}")
