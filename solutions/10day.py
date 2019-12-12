from math import atan2, sqrt

class Asteroid:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.asteroids_in_sight = 0

    def __repr__(self):
        return "<Asteroid ({}, {}) visible:{}>".format(self.x, self.y, self.asteroids_in_sight)

    def __str__(self):
        return "<Asteroid ({}, {}) visible:{}>".format(self.x, self.y, self.asteroids_in_sight)

    def add_rock(self):
        self.asteroids_in_sight += 1

def get_file():
    return './puzzledata/10day.txt'

def read_file():
    lines = open(get_file()).read().splitlines()
    asteroids = []
    for y, line in enumerate(lines):
        for x, space in enumerate(line):
            if space == "#":
                asteroids.append(Asteroid(x, y))
    return asteroids

def space_width():
    lines = open(get_file()).read().splitlines()
    # print("Width", len(lines[0]))
    return len(lines[0])

def space_height():
    lines = open(get_file()).read().splitlines()
    # print("Height", len(lines))
    return len(lines)

def find_delta(asteroid_1, asteroid_2):
    x_delta = asteroid_2.x - asteroid_1.x
    y_delta = asteroid_2.y - asteroid_1.y
    return {"x": asteroid_2.x - asteroid_1.x, "y": asteroid_2.y - asteroid_1.y}

def find_theta(delta):
    return round(atan2(delta["y"], delta["x"]), 3)

def find_hypotenuse(delta):
    return sqrt(delta["x"]**2 + delta["y"]**2)

def line_of_sight(location, asteroid, asteroids):
    if location == asteroid:
        return False
    delta_1 = find_delta(location, asteroid)
    theta_1 = find_theta(delta_1)
    hypotenuse_1 = find_hypotenuse(delta_1)
    for rock in asteroids:
        if rock == location or rock == asteroid:
            continue
        delta_2 = find_delta(location, rock)
        hypotenuse_2 = find_hypotenuse(delta_2)
        if hypotenuse_2 < hypotenuse_1:
            theta_2 = find_theta(delta_2)
            if theta_1 == theta_2:
                return False
    return True

def find_max_asteroids(asteroids):
    current_max = 0
    for asteroid in asteroids:
        for num, rock in enumerate(asteroids):
            if line_of_sight(asteroid, rock, asteroids):
                asteroid.add_rock()
            if asteroid.asteroids_in_sight + (len(asteroids) - num)  < current_max:
                continue
        if asteroid.asteroids_in_sight > current_max:
            current_max = asteroid.asteroids_in_sight

    max = {
        "visible": 0, 
        "location": {"x": 0, "y": 0}
    }
    for asteroid in asteroids:
        if asteroid.asteroids_in_sight > max["visible"]:
            max["visible"] = asteroid.asteroids_in_sight
            max["location"]["x"] = asteroid.x
            max["location"]["y"] = asteroid.y
    return max

asteroids = read_file()
# asteroid = find_max_asteroids(asteroids)
# part 1 answer
# {'visible': 334, 'location': {'x': 23, 'y': 20}}
print(asteroid)

# import pdb; pdb.set_trace()
