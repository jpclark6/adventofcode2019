from math import atan2, sqrt, pi, inf
import copy


class Asteroid:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.asteroids_in_sight = 0
        self.theta = 0
        self.distance = 0

    def __repr__(self):
        return "<Asteroid ({}, {}) visible:{} theta:{} hypotenuse:{}>" \
            .format(self.x, self.y, self.asteroids_in_sight, self.theta, self.distance)

    def __str__(self):
        return "<Asteroid ({}, {}) visible:{} theta:{} hypotenuse:{}>" \
            .format(self.x, self.y, self.asteroids_in_sight, self.theta, self.distance)

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

def find_delta(asteroid_1, asteroid_2):
    x_delta = asteroid_2.x - asteroid_1.x
    y_delta = -(asteroid_2.y - asteroid_1.y)
    return {"x": x_delta, "y": y_delta}

def find_theta(delta):
    return atan2(delta["y"], delta["x"])

def find_distance(delta):
    return sqrt(delta["x"]**2 + delta["y"]**2)

def find_max_asteroids(asteroids):
    current_max = 0
    a = copy.deepcopy(asteroids)
    for asteroid in asteroids:
        for rock in a:
            if asteroid.x == rock.x and asteroid.y == rock.y:
                rock.distance = inf
                continue
            delta = find_delta(asteroid, rock)
            rock.theta = find_theta(delta)
            rock.distance = find_distance(delta)
        a.sort(key=lambda x: (x.theta, x.distance))
        for i in range(len(a) - 1):
            if a[i].theta != a[i + 1].theta:
                asteroid.asteroids_in_sight += 1
        asteroid.asteroids_in_sight += 1
    asteroids.sort(key=lambda x: x.asteroids_in_sight)
    max_asteroid = asteroids[-1]
    return max_asteroid

def destroy_asteroid(location, theta, asteroids):
    possible_asteroids = []
    for num, asteroid in enumerate(asteroids):
        if theta >= asteroid.theta - .00005 and theta < asteroid.theta + .00005:
            possible_asteroids.append({"asteroid": asteroid, "num": num})

    if len(possible_asteroids) == 0:
        return asteroids, False, None

    destroy = possible_asteroids[0]["asteroid"]
    asteroid_number = possible_asteroids[0]["num"]

    # print("Destroying asteroid (", destroy.x, destroy.y, ") at", destroy.theta)
    del asteroids[asteroid_number]
    return asteroids, True, destroy

def destroy_asteroids(x_loc, y_loc, theta, quantity, asteroids):
    for i, asteroid in enumerate(asteroids):
        if asteroid.x == x_loc and asteroid.y == y_loc:
            location = copy.deepcopy(asteroid)
            asteroid_number = i
            break
    del asteroids[asteroid_number]

    for asteroid in asteroids:
        delta = find_delta(location, asteroid)
        asteroid.theta = find_theta(delta)
        asteroid.distance = find_distance(delta)
    asteroids.sort(key=lambda x: (x.theta, x.distance))
    
    destroyed = 0
    while destroyed < quantity:
        asteroids, bomb, dead = destroy_asteroid(location, theta, asteroids)
        if bomb:
            destroyed += 1
            if destroyed == 200:
                print("200th destroyed at", dead.x, dead.y)
                print("Part 2 answer:", dead.x * 100 + dead.y)

        theta -= .0001
        if theta - .0001 < - pi:
            theta += 2 * pi

asteroids = read_file()
max_asteroid= find_max_asteroids(asteroids)
print("Part 1 answer at location",  max_asteroid.x,  max_asteroid.y, "qty:",  max_asteroid.asteroids_in_sight)
destroy_asteroids(max_asteroid.x,  max_asteroid.y, pi / 2, 200, asteroids)
