# For example, given an image 3 pixels wide and 2 pixels tall, 
# the image data 123456789012 corresponds to the following image layers:

# Layer 1: 123
# 456

# Layer 2: 789
# 012

# The image you received is 25 pixels wide and 6 pixels tall.

def input():
    return open("puzzledata/8day.txt", "r").read().rstrip('\n')

def slice_layers(wide, tall, data):
    if len(data) % wide * tall != 0:
        print("Data is not correct length")
        return
    image = []
    layer = []
    while data:
        row = list(data[0:wide])
        row = [int(n) for n in row]
        data = data[wide:]
        layer.append(row)
        if len(layer) == tall:
            image.append(layer)
            layer = []
    return image

wide = 25
tall = 6
data = input()


def find_fewest_0_layer_multi_1_by_2(image):
    wide = len(image[0][0])
    tall = len(image[0])
    fewest = wide * tall
    for i, layer in enumerate(image):
        pixels = [pixel for row in layer for pixel in row]
        num_zeros = pixels.count(0)
        if num_zeros < fewest:
            fewest = num_zeros
            fewest_layer = i

    pixels = [pixel for row in image[fewest_layer] for pixel in row]
    return pixels.count(1) * pixels.count(2)

def print_image(image):
    wide = len(image[0][0])
    tall = len(image[0])
    final_image = [[-1 for _ in range(wide)] for _ in range(tall)]

    for l, layer in enumerate(image):
        for r, row in enumerate(layer):
            for p, pixel in enumerate(row):
                if final_image[r][p] != -1 or pixel == 2:
                    pass
                else:
                    if pixel == 1:
                        final_image[r][p] = "X"
                    elif pixel == 0:
                        final_image[r][p] = " "
    
    print("\n")
    for row in final_image:
        print("".join(row))
    print("\n")
    # import pdb; pdb.set_trace()



image = slice_layers(wide, tall, data)
ans_part_1 = find_fewest_0_layer_multi_1_by_2(image)
print("Part 1 answer", ans_part_1)

print_image(image)
