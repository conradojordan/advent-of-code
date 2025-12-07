from day01 import solve


def test_part1_and_2():
    with open("example.txt") as f:
        entries = [int(entry) for entry in f.read().strip().split()]

    a, b = solve(entries, 2)
    assert a * b == 514579

    a, b, c = solve(entries, 3)
    assert a * b * c == 241861950
