def password_check_part_1(password):
    if only_increases(password) and two_adjacent_part_1(password) and six_digit(password):
        return True
    return False

def password_check_part_2(password):
    if only_increases(password) and two_adjacent_part_2(password) and six_digit(password):
        return True
    return False

def six_digit(password):
    password = str(password)
    if len(password) != 6:
        return False
    return True

def two_adjacent_part_1(password):
    password = str(password)
    for i in range(len(password) - 1):
        if password[i] == password[i + 1]:
            return True
    return False

def two_adjacent_part_2(password):
    password = str(password)
    current_loc = 0
    has_double = False
    while current_loc < len(password) - 1:
        if password[current_loc] == password[current_loc + 1]:
            i = 1
            while i < len(password) - current_loc and password[current_loc] == password[current_loc + i]:
                i += 1
            if i == 2:
                has_double = True
            current_loc += i
        else:
            current_loc += 1
    return has_double

def only_increases(password):
    password = str(password)
    for i in range(len(password) - 1):
        if int(password[i]) > int(password[i + 1]):
            return False
    return True

start = 372304
end = 847060
valid_1 = []

for password in range(start, end + 1):
    if password_check_part_1(password):
        valid_1.append(password)
print("Part 1 solution:", len(valid_1))

valid_2 = []

for password in range(start, end + 1):
    if password_check_part_2(password):
        valid_2.append(password)

print("Part 2 solution", len(valid_2))
