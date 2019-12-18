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

def run_part_1():
    data = open('./puzzledata/16day.txt').read()
    # data = '80871224585914546619083218645595'
    # data = '12345678'
    signal = [int(d) for d in data]
    base_pattern = [0, 1, 0, -1]

    d = time.time()
    run_phases(signal, base_pattern, 100)
    e = time.time() - d
    print(e, "seconds for part 1")

run_part_1()
