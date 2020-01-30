import math


def calc_fuel(module_mass):
    return math.floor(module_mass / 3) - 2


if __name__ == "__main__":
    total = 0
    with open("01_1__input.txt", "r") as modules:
        for module in modules:
            fuel = calc_fuel(int(module))
            while fuel > 0:
                total += fuel
                fuel = calc_fuel(fuel)
    print(f"total fuel required: {total}")
