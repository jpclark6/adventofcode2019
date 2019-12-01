import math
import time

start = time.time()

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

mid = time.time()

# Day 2
total = 0
for mass in input():
    current_fuel = find_fuel(mass)
    while current_fuel > 0:
        total += current_fuel
        current_fuel = find_fuel(current_fuel)

print("Total part 2:", total)

end = time.time()
print("Elapsed time part 1:\t", round(mid - start, 6), "seconds")
print("Elapsed time part 2:\t", round(end - mid, 6), "seconds")
print("Elapsed time total:\t", round(end - start, 6), "seconds")