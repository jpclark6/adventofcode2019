import math

def find_fuel(mass):
    return math.floor(mass / 3) - 2

def input():
    f = open("puzzledata/1day.txt", "r")
    return [int(i) for i in f.readlines()]

# Day 1
total = 0
for mass in input():
    total += find_fuel(mass)

print("Total part 1:", total)

total = 0
for mass in input():
    current_fuel = find_fuel(mass)
    while current_fuel > 0:
        total += current_fuel
        current_fuel = find_fuel(current_fuel)

print("Total part 2:", total)
