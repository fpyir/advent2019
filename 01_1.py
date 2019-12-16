import math


def calculate_fuel_required(module_mass):
    return math.floor(module_mass / 3) - 2


if __name__ == "__main__":
    total_fuel_required = 0
    with open("01_1__input.txt", "r") as modules:
        for module in modules:
            total_fuel_required += calculate_fuel_required(int(module))
    print(f"total fuel required: {total_fuel_required}")
