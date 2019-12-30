file = './puzzledata/22day.txt'
file = open(file)
data = file.read().splitlines()
instructions = []
shuffle = []

# Part 1
# num_cards = 10007
# card_number = 2019

# Part 2
num_cards = 119315717514047
card_position = 2020
deals = 101741582076661

for line in data:
    inst = line.split(" ")
    if inst[1] == "with":
        increment = int(inst[3])
        instructions.append(("increment", increment))
        shuffle.append((increment, 0))
    elif inst[0] == "cut":
        cut = int(inst[1])
        instructions.append(("cut", cut))
        shuffle.append((1, -cut))
    else:
        instructions.append(("stack", 0))
        shuffle.append((-1, -1))

current = [shuffle[0][0], shuffle[0][1]]
for i, tup in enumerate(shuffle):
    if i < len(shuffle) - 1:
        next = [0, 0]
        next[0] = shuffle[i + 1][0]
        next[1] = shuffle[i + 1][1]
        current = [(current[0] * next[0]) % num_cards, (current[1] * next[0] + next[1]) % num_cards]

def ex_by_sq(base, exp, mod):
    if exp == 0:
        return 1
    if exp == 1:
        return base % mod

    t = ex_by_sq(base, int(exp / 2), mod)
    t = (t * t) % mod

    if exp % 2 == 0:
        return t
    else:
        return ((base % mod) * t) % mod

def mod_div(num, den, mod):
    num = num % mod
    inv = pow(den, mod - 2, mod)
    return (num * inv) % mod

a = ex_by_sq(current[0], deals, num_cards)
num = current[1] * (1 - a)
den = 1 - current[0]
b = mod_div(num, den, num_cards)

num = card_position - b
den = a

ans_part_2 = mod_div(num, den, num_cards)

# print("Answer to part 1:", (current[0] * card_number + current[1]) % num_cards)
print("Answer to part 2:", ans_part_2)
