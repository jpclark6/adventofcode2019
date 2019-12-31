import time

# Global vars. Yuck
filename = './puzzledata/24day.txt'
minutes_to_iterate = 200

def read_data(file):
    file = open(file).read().splitlines()
    data = [[bug for bug in row] for row in file]
    return data

def check_space(data, x, y):
    surrounding_bugs = 0
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]
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

def check_side(dir, layers, depth):
    if dir == "left":
        side = []
        for y in range(5):
            side.append(layers[depth][y][0])
        return side.count("#")
    elif dir == "right":
        side = []
        for y in range(5):
            side.append(layers[depth][y][-1])
        return side.count("#")
    elif dir == "top":
        return layers[depth][0].count("#")
    elif dir == "bottom":
        return layers[depth][-1].count("#")

def check_3d_space(layers, depth, x, y):
    surrounding_bugs = 0
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    data = layers[depth]
    for direction in directions:
        check_x = x + direction[0]
        check_y = y + direction[1]
        if check_x >= 0 and check_x < len(data[0]) and check_y >= 0 and check_y < len(data):
            if data[check_y][check_x] == "#":
                surrounding_bugs += 1
            elif data[check_y][check_x] == "?":
                if x == 1:
                    surrounding_bugs += check_side("left", layers, depth - 1)
                elif x == 3:
                    surrounding_bugs += check_side("right", layers, depth - 1)
                elif y == 1:
                    surrounding_bugs += check_side("top", layers, depth - 1)
                elif y == 3:
                    surrounding_bugs += check_side("bottom", layers, depth - 1)
        else:
            if check_x == -1:
                if layers[depth + 1][2][1] == "#":
                    surrounding_bugs += 1
            elif check_x == 5:
                if layers[depth + 1][2][3] == "#":
                    surrounding_bugs += 1
            elif check_y == -1:
                if layers[depth + 1][1][2] == "#":
                    surrounding_bugs += 1
            elif check_y == 5:
                if layers[depth + 1][3][2] == "#":
                    surrounding_bugs += 1
    return surrounding_bugs

def empty_layer():
    layer = [["." for _ in range(5)] for _ in range(5)]
    layer[2][2] = "?"
    return layer

def one_3d_iteration(layers):
    next_state = {}
    for k, v in layers.items():
        data = [[bug for bug in row] for row in v]
        next_state[k] = data
    depth = len(next_state) // 2 + 1

    if layers[depth - 1] != empty_layer() or layers[1 - depth] != empty_layer():
        next_state[depth], next_state[-depth] = empty_layer(), empty_layer()
    layers[depth], layers[-depth] = empty_layer(), empty_layer()

    for depth, data in layers.items():
        if depth == len(layers) // 2 or depth == -len(layers) // 2 + 1:
            pass
        else:
            for y, row in enumerate(data):
                for x, bug in enumerate(row):
                    if bug != "?":
                        surrounding_bugs = check_3d_space(layers, depth, x, y)
                        save_as_bug = check_bug(surrounding_bugs, data, x, y)
                        if save_as_bug:
                            next_state[depth][y][x] = "#"
                        else:
                            next_state[depth][y][x] = "."
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

def part_2():
    layers = {0: read_data(filename)}
    layers[0][2][2] = "?"
    layers[1], layers[-1] = empty_layer(), empty_layer()
    for i in range(minutes_to_iterate):
        layers = one_3d_iteration(layers)
    bugs = 0
    for depth, layer in layers.items():
        for y in layer:
            for bug in y:
                if bug == "#":
                    bugs += 1
    print("Part 2 bugs:", bugs)

def print_layer(layers):
    for k, v in layers.items():
        print()
        print("Depth", -k)
        for row in v:
            print("".join(row))


s = time.time()
data = read_data(filename)
part_1_find_repeat(data)
part_2()
e = time.time()
print("Elapsed time:", round(e - s, 2))