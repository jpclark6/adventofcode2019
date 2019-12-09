
class Planet:
    def __init__(self, name="", parent=None):
        self.name = name
        self.parent = parent

class PlanetSystem:
    def __init__(self):
        self.planets = {}

    def add_planet(self, body_name, moon_name):
        self.planets[moon_name] = Planet(moon_name, body_name)

    def find_planet(self, name):
        return self.planets[name]

    def find_planet_parent(self, name):
        try:
            return self.find_planet(self.find_planet(name).parent)
        except KeyError:
            return None
        
def input():
    f = open("puzzledata/6day.txt", "r")
    return [i.split(")") for i in f.read().splitlines()]

system = PlanetSystem()
for planet in input():
    system.add_planet(planet[0], planet[1])

total = 0
for planet, parent in system.planets.items():
    current_planet = system.find_planet(planet)
    total += 1
    while system.find_planet_parent(current_planet.name):
        total += 1
        current_planet = system.find_planet_parent(current_planet.name)

print("Total for part 1:", total)

part_2 = {'YOU': [], 'SAN': []}
for planet, orbz in part_2.items():
    current_planet = system.find_planet(planet)
    while system.find_planet_parent(current_planet.name):
        current_planet = system.find_planet_parent(current_planet.name)
        orbz.append(current_planet.name)

my_planets = set(part_2['YOU'])
santa_planets = set(part_2['SAN'])

planet_path_count = len(my_planets - santa_planets) + len(santa_planets - my_planets)

print("Total for part 2:", planet_path_count)
