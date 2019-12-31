# A bug dies(becoming an empty space) unless there is exactly one bug adjacent to it.
# An empty space becomes infested with a bug if exactly one or two bugs are adjacent to it.

def read_data(file):
    file = open(file).read().splitlines()
    data = [[bug for bug in row] for row in file]
    return data
data = read_data('./puzzledata/24day.txt')

def check_space(data, x, y):
    surrounding_bugs = 0
    directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]
    for direction in directions:
        check_x = x + direction[0]        
        check_y = y + direction[1]    
        if check_x >= 0 and check_x < len(data[0]) and check_y >= 0 and check_y < len(data):
            if data[check_y][check_x] == "#":
                surrounding_bugs += 1
    return surrounding_bugs

def check_bug(num_bugs, data, x, y):
    if data[y][x] == "#":
        if num_bugs == 1:
            return True
        return False
    if num_bugs == 1 or num_bugs == 2:
        return True
    return False

def one_iteration(data):
    next_state = [[bug for bug in row] for row in data]
    for y, row in enumerate(data):
        for x, bug in enumerate(row):
            surrounding_bugs = check_space(data, x, y)
            save_as_bug = check_bug(surrounding_bugs, data, x, y)
            if save_as_bug:
                next_state[y][x] = "#"
            else:
                next_state[y][x] = "."
    return next_state

def find_bio_rating(data):
    total = 0
    rating = 1
    for row in data:
        for bug in row:
            if bug == "#":
                total += rating
            rating *= 2
    return total

def part_1_find_repeat(data):
    states = set()
    i = 0
    while True:
        i += 1
        current_length = len(states)
        key = ["".join(row) for row in data]
        key = "".join(key)
        states.add(key)
        if len(states) == current_length:
            bio_rating = find_bio_rating(data)
            print("Found repeat", bio_rating)
            return
        data = one_iteration(data)

part_1_find_repeat(data)
