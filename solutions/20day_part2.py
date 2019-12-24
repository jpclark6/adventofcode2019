import re, time
from collections import OrderedDict

class Puzzle:
    def __init__(self, file_name):
        self.file_name = file_name
        self.tiles = {}
        self.create_level()
        self.max_depth = 25

    def create_level(self):
        file = open(self.file_name).read().splitlines()
        matrix = [[x for x in line] for line in file]
        for y in range(len(matrix)):
            for x in range(len(matrix[0])):
                if bool(re.search('[#.]', matrix[y][x])):
                    tile = Tile(x, y, matrix[y][x], 0)
                    self.tiles[str(x) + "," + str(y)] = tile
                    if bool(re.search('[A-Z]', matrix[y - 1][x])):
                        if (matrix[y - 2][x] + matrix[y - 1][x] == "AA" or matrix[y - 2][x] + matrix[y - 1][x] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(0) + "," + str(0)
                            tile.direction = 0
                        elif y > 3:
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(0) + "," + str(1)
                            tile.direction = 1
                        elif y <= 3:
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(1) + "," + str(0)
                            tile.direction = -1
                    elif bool(re.search('[A-Z]', matrix[y + 1][x])):
                        if (matrix[y + 1][x] + matrix[y + 2][x] == "AA" or matrix[y + 1][x] + matrix[y + 2][x] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(0) + "," + str(0)
                            tile.direction = 0
                        elif y < len(matrix) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(0) + "," + str(1)
                            tile.direction = 1
                        elif y >= len(matrix) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(1) + "," + str(0)
                            tile.direction = -1
                    elif bool(re.search('[A-Z]', matrix[y][x - 1])):
                        if (matrix[y][x - 2] + matrix[y][x - 1] == "AA" or matrix[y][x - 2] + matrix[y][x - 1] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(0) + "," + str(0)
                            tile.direction = 0
                        elif x > 3:
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(0) + "," + str(1)
                            tile.direction = 1
                        elif x <= 3:
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(1) + "," + str(0)
                            tile.direction = -1
                    elif bool(re.search('[A-Z]', matrix[y][x + 1])):
                        if (matrix[y][x + 1] + matrix[y][x + 2] == "AA" or matrix[y][x + 1] + matrix[y][x + 2] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(0) + "," + str(0)
                            tile.direction = 0
                        elif x < len(matrix[0]) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(0) + "," + str(1)
                            tile.direction = 1
                        elif x >= len(matrix[0]) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(1) + "," + str(0)
                            tile.direction = -1

    def find_start(self):
        for loc, tile in self.tiles.items():
            if tile.portal == True and tile.portal_key == "AA,0,0":
                return tile

    def find_end(self):
        start = self.find_start()
        i = 0
        visited = {}
        visited[start.make_key(0)] = True
        queue = self.check_surrounding_spaces(start, visited, i, 0)
        while True:
            i += 1
            new_queue = []
            for tile_layer in queue:
                tile = tile_layer['tile']
                layer = tile_layer['layer']
                if tile.portal_key == "ZZ,0,0" and layer == 0:
                    print("Found end:", i)
                    return
                visited[tile.make_key(layer)] = True
                new_queue += self.check_surrounding_spaces(tile, visited, i, layer)
            queue = new_queue

    def check_surrounding_spaces(self, tile, visited, i, layer):
        queue = []
        if tile.portal:
            portal_locs = self.find_portal(tile.portal_key, layer)
            for portal in portal_locs:
                if layer - portal.direction > self.max_depth or visited.get(portal.make_key(layer - portal.direction)):
                    continue
                else:
                    return [{'tile': portal, 'layer': layer - portal.direction}]
        for tile_key in tile.make_key_check():
            try:
                tile = self.tiles[tile_key]
            except KeyError:
                continue
            try:
                if tile.space == "." and not self.recently_visited(visited, tile, i, layer):
                    queue.append({'tile': tile, 'layer': layer})
            except KeyError:
                queue.append({'tile': tile, 'layer': layer})
        return queue

    def recently_visited(self, visited, tile, i, layer):
        if visited.get(tile.make_key(layer)) == True:
            return True
        return False
    
    def find_portal(self, portal_key, current_layer):
        tiles = []
        key = portal_key.split(",")[0]
        layer = portal_key.split(",")[1]
        to_layer = portal_key.split(",")[2]
        for loc, tile in self.tiles.items():
            if tile.portal == True:
                key_check = tile.portal_key.split(",")[0]
                layer_check = tile.portal_key.split(",")[1]
                to_layer_check = tile.portal_key.split(",")[2]
                if key == key_check and layer == to_layer_check and layer_check == to_layer:
                    if current_layer == 0 and layer == '1':
                        pass
                    else:
                        tiles.append(tile)
        return tiles

class Tile:
    def __init__(self, x, y, space, layer, portal=False, portal_key=None):
        self.x = x
        self.y = y
        self.space = space
        self.layer = layer
        self.portal = portal
        self.portal_key = portal_key
        self.portal_visited = 0
        self.direction = 0

    def __repr__(self):
        if self.portal:
            add_key = self.portal_key
        else:
            add_key = ""
        return "(" + str(self.x) + "," + str(self.y) + "):" + self.space + add_key

    def __str__(self):
        if self.portal:
            add_key = self.portal_key
        else:
            add_key = ""
        return "(" + str(self.x) + "," + str(self.y) + "):" + self.space + add_key

    def make_key(self, i):
        return str(self.x) + "," + str(self.y) + "," + str(i)

    def make_key_check(self):
        return [
            str(self.x + 1) + "," + str(self.y),
            str(self.x - 1) + "," + str(self.y),
            str(self.x) + "," + str(self.y + 1),
            str(self.x) + "," + str(self.y - 1),
        ]

s = time.time()
puzzle = Puzzle('./puzzledata/20day.txt')
puzzle.find_end()
e = time.time()
print("Time for part 2:", round(e - s, 3), "Seconds")