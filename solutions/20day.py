class Puzzle:
    def __init__(self, file_name):
        self.tiles = []
        file = open(file_name).read().splitlines()
        matrix = [[x for x in line] for line in file]
        for y in range(2, len(matrix) - 2):
            for x in range(2, len(matrix[0]) - 2):
                self.tiles.append(Tile(x, y, matrix[y][x]))
        import pdb; pdb.set_trace()

    def get_tile(self, x, y):
        
        

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