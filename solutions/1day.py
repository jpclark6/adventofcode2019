import math

def find_fuel(mass):
	return math.floor(mass / 3) - 2

f = open("../puzzledata/1day.txt", "r")
lines = f.readlines()
total = 0

for line in lines:
	total += find_fuel(int(line))

print("Total part 1:", total)

f.seek(0)
lines = f.readlines()
total = 0

for line in lines:
	mass = int(line)
	current_fuel = find_fuel(mass)
	while current_fuel > 0:
		total += current_fuel
		current_fuel = find_fuel(current_fuel)

print("Total part 2:", total)
