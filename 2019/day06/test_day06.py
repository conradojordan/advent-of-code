from day06 import count_orbits


def test_count_orbits():
    assert count_orbits("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L") == 42
