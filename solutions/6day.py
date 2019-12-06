# import math
# import time

def input():
    f = open("puzzledata/6day_example.txt", "r")
    return [i.split(")") for i in f.read().splitlines()]

orbits = {}

for orbit in input():
    if orbit[0] in orbits:
        orbits[orbit[0]].append(orbit[1])
    else:
        orbits[orbit[0]] = [orbit[1]]

print(count)


# start = time.time()
# mid = time.time()
# end = time.time()
# print("Elapsed time part 1:\t", round(mid - start, 6), "seconds")
# print("Elapsed time part 2:\t", round(end - mid, 6), "seconds")
# print("Elapsed time total:\t", round(end - start, 6), "seconds")
