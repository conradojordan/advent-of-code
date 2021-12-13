import pytest
from ..day3 import (
    closest_cross_to_center,
    calculate_points,
    distance,
    fewest_combined_steps,
)


@pytest.mark.parametrize(
    "wire, expected_points",
    [
        ("R3,U2", {(1, 0), (2, 0), (3, 0), (3, 1), (3, 2)}),
        ("D3,L1", {(0, -1), (0, -2), (0, -3), (-1, -3)}),
    ],
)
def test_calculate_points(wire, expected_points):
    assert calculate_points(wire) == expected_points


@pytest.mark.parametrize(
    "point_a, point_b, expected_distance",
    [
        ((2, 3), (0, 0), 5),
        ((-1, -1), (-1, 5), 6),
        ((-1, 2), (2, -3), 8),
    ],
)
def test_distance(point_a, point_b, expected_distance):
    assert distance(point_a, point_b) == expected_distance


@pytest.mark.parametrize(
    "wires, closest_cross_distance",
    [
        ("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159),
        (
            "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
            135,
        ),
    ],
)
def test_closest_cross_to_center(wires, closest_cross_distance):
    assert closest_cross_to_center(wires)[1] == closest_cross_distance


@pytest.mark.parametrize(
    "wires, cross_points, expected_fewest_steps",
    [
        (
            "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
            closest_cross_to_center(
                "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
            )[2],
            610,
        ),
        (
            "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
            closest_cross_to_center(
                "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
            )[2],
            410,
        ),
    ],
)
def test_fewest_combined_steps(wires, cross_points, expected_fewest_steps):
    assert fewest_combined_steps(wires, cross_points) == expected_fewest_steps
