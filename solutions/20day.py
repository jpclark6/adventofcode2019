import re, time

class Puzzle:
    def __init__(self, file_name):
        self.tiles = {}
        file = open(file_name).read().splitlines()
        matrix = [[x for x in line] for line in file]
        for y in range(len(matrix)):
            for x in range(len(matrix[0])):
                if bool(re.search('[#.]', matrix[y][x])):
                    tile = Tile(x, y, matrix[y][x])
                    self.tiles[str(x) + "," + str(y)] = tile
                    if bool(re.search('[A-Z]', matrix[y - 1][x])):
                        tile.portal = True
                        tile.portal_key = matrix[y - 2][x] + matrix[y - 1][x]
                    elif bool(re.search('[A-Z]', matrix[y + 1][x])):
                        tile.portal = True
                        tile.portal_key = matrix[y + 1][x] + matrix[y + 2][x]
                    elif bool(re.search('[A-Z]', matrix[y][x - 1])):
                        tile.portal = True
                        tile.portal_key = matrix[y][x - 2] + matrix[y][x - 1]
                    elif bool(re.search('[A-Z]', matrix[y][x + 1])):
                        tile.portal = True
                        tile.portal_key = matrix[y][x + 1] + matrix[y][x + 2]

    def get_tile(self, x, y):
        key = str(x) + "," + str(y)
        return self.tiles[key]

    def find_end(self):
        start = self.find_portal("AA")[0]
        i = 0
        visited = {start.make_key(): [i]}
        queue = self.check_surrounding_spaces(start, visited, i)
        while True:
            i += 1
            new_queue = []
            for tile in queue:
                # import pdb; pdb.set_trace()
                if tile.portal_key == "ZZ":
                    print("Found end:", i)
                    return
                if visited.get(tile.make_key()):
                    visited[tile.make_key()].append(i)
                else:
                    visited[tile.make_key()] = [i]
                new_queue += self.check_surrounding_spaces(tile, visited, i)
            queue = new_queue

    def check_surrounding_spaces(self, tile, visited, i):
        queue = []
        if tile.portal:
            portal_locs = self.find_portal(tile.portal_key)
            for portal in portal_locs:
                try:
                    if self.recently_visited(visited[portal.make_key()], i):
                        continue
                except KeyError:
                    return [portal]
        for tile in tile.make_key_check():
            try:
                tile = self.tiles[tile]
            except KeyError:
                continue
            try:
                if tile.space == "." and not self.recently_visited(visited[tile.make_key()], i):
                    queue.append(tile)
            except KeyError:
                queue.append(tile)
        return queue

    def recently_visited(self, visited, i):
        if len(visited) > 4:
            return True
        for time in visited:
            if time >= i - 1:
                return True
        return False
    
    def find_portal(self, letters):
        tiles = []
        for loc, tile in self.tiles.items():
            if tile.portal == True and tile.portal_key == letters:
                tiles.append(tile)
        
        return tiles
        

class Tile:
    def __init__(self, x, y, space, portal=False, portal_key=None):
        self.x = x
        self.y = y
        self.space = space
        self.portal = portal
        self.portal_key = portal_key

    def __repr__(self):
        if self.portal:
            add_key = self.portal_key
        else:
            add_key = ""
        return "(" + str(self.x) + "," + str(self.y) + ")" + self.space + add_key

    def __str__(self):
        return self.__repr__
        
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
        return str(self.x) + "," + str(self.y)

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
print("Time for part 1:", round(e - s, 3), "Seconds")
