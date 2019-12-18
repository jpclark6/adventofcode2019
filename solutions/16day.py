import math
import time
from collections import deque
from itertools import cycle


def find_patterns(signal, base_pattern):
    patterns = []
    for i, point in enumerate(signal):
        modified_base_pattern = deque()
        for bp in base_pattern:
            for j in range(i + 1):
                modified_base_pattern.append(bp)
        modified_base_pattern.rotate(-1)
        patterns.append(modified_base_pattern)
    return patterns

def run_phases(signal, base_pattern, phases):
    patterns = find_patterns(signal, base_pattern)
    for ph in range(phases):
        if ph % 20 == 0:
            print(ph, "%")
        new_signal = []
        for i, point in enumerate(signal):
            next_digit = abs(sum([a*b for a, b in zip(signal, cycle(patterns[i]))])) % 10
            new_signal.append(next_digit)
        signal = new_signal

    print("Answer to part 1:", "".join([str(j) for j in signal[:8]]))

def get_data():
    data = open('./puzzledata/16day.txt').read()
    return data

def run_part_1():
    data = get_data()
    # data = '80871224585914546619083218645595'
    # data = '12345678'
    signal = [int(d) for d in data]
    base_pattern = [0, 1, 0, -1]

    run_phases(signal, base_pattern, 100)


def run_part_2(phases):
    phases = phases
    repeated = 10000
    data = get_data()
    signal = [int(d) for d in data]
    offset = int(data[:7])
    relavent_data_length = (len(data) * repeated - offset)
    relavent_data = []
    while len(relavent_data) < relavent_data_length:
        relavent_data.extend(signal)
    offset_from_start = len(relavent_data) - relavent_data_length
    run_part_2_phases(relavent_data, phases, offset_from_start)

def run_part_2_phases(data, phases, offset_from_start):
    for x in range(phases):
        if x % 10 == 0:
            print(x, "%")
        next_line = [0 for a in range(len(data))]
        for i, digit in enumerate(data):
            next_line[-1 - i] = (data[-1 - i] + next_line[-i]) % 10
        data = next_line
    print("Answer to part 2:", "".join([str(d) for d in data[offset_from_start: offset_from_start + 8]]))


d = time.time()
run_part_1()
m = time.time()
run_part_2(100)
e = time.time()

p1 = m - d
p2 = e - m
tot = e - d

print("Part 1 took", p1, "seconds")
print("Part 2 took", p2, "seconds")
print("Total took", tot, "seconds")
