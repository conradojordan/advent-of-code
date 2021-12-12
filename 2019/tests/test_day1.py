import pytest
from solutions.day1 import calculate_fuel, calculate_fuel_recursive


@pytest.mark.parametrize(
    "input, expected_fuel", [(12, 2), (14, 2), (1969, 654), (100756, 33583)]
)
def test_calculate_fuel(input, expected_fuel):
    assert calculate_fuel(input) == expected_fuel


@pytest.mark.parametrize(
    "input, expected_fuel", [(14, 2), (1969, 966), (100756, 50346)]
)
def test_calculate_fuel_recursive(input, expected_fuel):
    assert calculate_fuel_recursive(input) == expected_fuel

