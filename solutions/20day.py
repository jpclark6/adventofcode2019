import re

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

        import pdb; pdb.set_trace()

    def get_tile(self, x, y):
        key = str(x) + "," + str(y)
        return self.tiles[key]
        

class Tile:
    def __init__(self, x, y, space, portal=False, portal_key=None):
        self.x = x
        self.y = y
        self.space = space
        self.portal = portal

    def __repr__(self):
        return "(" + str(self.x) + "," + str(self.y) + ")" + self.space

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
    

x = Puzzle('./puzzledata/20day.txt')
