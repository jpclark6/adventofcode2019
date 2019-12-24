import re
from collections import OrderedDict

class Puzzle:
    def __init__(self, file_name):
        self.file_name = file_name
        self.tiles = {}
        for i in range(30):
            self.create_level(i)
        self.max_layer = 29

    def create_level(self, layer):
        file = open(self.file_name).read().splitlines()
        matrix = [[x for x in line] for line in file]
        for y in range(len(matrix)):
            for x in range(len(matrix[0])):
                if bool(re.search('[#.]', matrix[y][x])):
                    tile = Tile(x, y, matrix[y][x], layer)
                    self.tiles[str(layer) + "," + str(x) + "," + str(y)] = tile
                    if bool(re.search('[A-Z]', matrix[y - 1][x])):
                        if y > 3:
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(layer) + "," + str(layer + 1)
                            tile.portal_layer = layer + 1
                            tile.portal_inward = True
                        elif layer > 0:
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(layer) + "," + str(layer - 1)
                            tile.portal_layer = layer - 1
                        elif layer == 0 and (matrix[y - 2][x] + matrix[y - 1][x] == "AA" or matrix[y - 2][x] + matrix[y - 1][x] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x] + "," + str(layer) + "," + str(layer)
                            tile.portal_layer = -1
                    elif bool(re.search('[A-Z]', matrix[y + 1][x])):
                        if y < len(matrix) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(layer) + "," + str(layer + 1)
                            tile.portal_layer = layer + 1
                            tile.portal_inward = True
                        elif layer > 0:
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(layer) + "," + str(layer - 1)
                            tile.portal_layer = layer - 1
                        elif layer == 0 and (matrix[y + 1][x] + matrix[y + 2][x] == "AA" or matrix[y + 1][x] + matrix[y + 2][x] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x] + "," + str(layer) + "," + str(layer)
                            tile.portal_layer = -1
                    elif bool(re.search('[A-Z]', matrix[y][x - 1])):
                        if x > 3:
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(layer) + "," + str(layer + 1)
                            tile.portal_layer = layer + 1
                            tile.portal_inward = True
                        elif layer > 0:
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(layer) + "," + str(layer - 1)
                            tile.portal_layer = layer - 1
                        elif layer == 0 and (matrix[y][x - 2] + matrix[y][x - 1] == "AA" or matrix[y][x - 2] + matrix[y][x - 1] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1] + "," + str(layer) + "," + str(layer)
                            tile.portal_layer = -1
                    elif bool(re.search('[A-Z]', matrix[y][x + 1])):
                        if x < len(matrix[0]) / 2:
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(layer) + "," + str(layer + 1)
                            tile.portal_layer = layer + 1
                            tile.portal_inward = True
                        elif layer > 0:
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(layer) + "," + str(layer - 1)
                            tile.portal_layer = layer - 1
                        elif layer == 0 and (matrix[y][x + 1] + matrix[y][x + 2] == "AA" or matrix[y][x + 1] + matrix[y][x + 2] == "ZZ"):
                            tile.portal = True
                            tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2] + "," + str(layer) + "," + str(layer)
                            tile.portal_layer = -1

    def get_tile(self, layer, x, y):
        key = str(layer) + "," + str(x) + "," + str(y)
        return self.tiles[key]

    def find_start(self):
        for loc, tile in self.tiles.items():
            if tile.portal == True and tile.portal_key == "AA,0,0":
                return tile

    def find_end(self):
        for loc, tile in self.tiles.items():
            if tile.portal == True and tile.portal_key == "ZZ,0,0":
                return tile

    def find_end(self):
        start = self.find_start()
        i = 0
        visited = {}
        visited[start.make_key()] = i
        queue = self.check_surrounding_spaces(start, visited, i)
        while True:
            i += 1
            if i % 100 == 0:
                print(i)
            new_queue = []
            for tile in queue:
                if tile.portal_key == "ZZ,0,0":
                    print("Found end:", i)
                    return
                visited[tile.make_key()] = i
                new_queue += self.check_surrounding_spaces(tile, visited, i)
            if new_queue == []:
                import pdb; pdb.set_trace()
            queue = new_queue

    def check_surrounding_spaces(self, tile, visited, i):
        queue = []
        layer = tile.layer
        if self.max_layer < layer + 3:
            for _ in range(10):
                self.create_level(self.max_layer + 1)
                self.max_layer += 1
        if tile.portal:
            portal_locs = self.find_portal(tile.portal_key)
            for portal in portal_locs:
                if self.recently_visited(visited, portal, i):
                    continue
                else:
                    return [portal]
        for tile_key in tile.make_key_check(layer):
            try:
                tile = self.tiles[tile_key]
            except KeyError:
                continue
            try:
                if tile.space == "." and not self.recently_visited(visited, tile, i):
                    queue.append(tile)
            except KeyError:
                queue.append(tile)
        return queue

    def recently_visited(self, visited, tile, i):
        try:
            visited_times = visited[tile.make_key()]
            if visited_times >= i - 1:
                return True
            return False
        except KeyError:
            return False
    
    def find_portal(self, portal_key):
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
                    tiles.append(tile)
        
        return tiles
        

class Tile:
    def __init__(self, x, y, space, layer, portal=False, portal_key=None, portal_layer=0):
        self.x = x
        self.y = y
        self.space = space
        self.layer = layer
        self.portal = portal
        self.portal_key = portal_key
        self.portal_layer = portal_layer
        self.portal_visited = 0

    def __repr__(self):
        if self.portal:
            add_key = self.portal_key
        else:
            add_key = ""
        return str(self.layer) + ": (" + str(self.x) + "," + str(self.y) + "):" + self.space + add_key

    def __str__(self):
        if self.portal:
            add_key = self.portal_key
        else:
            add_key = ""
        return str(self.layer) + ": (" + str(self.x) + "," + str(self.y) + "):" + self.space + add_key
        
    def wall(self):
        if self.space == "#":
            return True
        else:
            return False

    def passage(self):
        if self.space == ".":
            return True
        else:
            return False

    def make_key(self):
        return str(self.layer) + "," + str(self.x) + "," + str(self.y)

    def make_key_check(self, layer):
        return [
            str(layer) + "," + str(self.x + 1) + "," + str(self.y),
            str(layer) + "," + str(self.x - 1) + "," + str(self.y),
            str(layer) + "," + str(self.x) + "," + str(self.y + 1),
            str(layer) + "," + str(self.x) + "," + str(self.y - 1),
        ]
    

puzzle = Puzzle('./puzzledata/20day.txt')
puzzle.find_end()
