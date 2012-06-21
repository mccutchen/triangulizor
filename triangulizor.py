import logging

import Image
import ImageDraw


def prep_image(image, tile_size):
    """Takes an image and a tile size and returns a possibly cropped version
    of the image that is evenly divisible in both dimensions by the tile size.
    """
    assert isinstance(image, Image.Image)
    assert isinstance(tile_size, int)
    w, h = image.size
    new_w, rem_w = divmod(w, tile_size)
    new_h, rem_h = divmod(h, tile_size)
    if new_w == w and new_h == h:
        return image
    else:
        dw = rem_w / 2
        dh = rem_h / 2
        crop_bounds = (dw, dh, w - dw, h - dh)
        return image.crop(crop_bounds)

def iter_tiles(image, tile_size):
    w, h = image.size
    for y in xrange(0, h, tile_size):
        for x in xrange(0, w, tile_size):
            yield x, y

def triangulize(original_image, tile_size):
    image = prep_image(original_image, tile_size)
    pix = image.load()
    draw = ImageDraw.Draw(image)
    for x, y in iter_tiles(image, tile_size):
        process_tile(x, y, tile_size, pix, draw, image)
    image.show()

def process_tile(tile_x, tile_y, tile_size, pix, draw, image):
    colors = []
    for y in xrange(tile_y, tile_y + tile_size):
        for x in xrange(tile_x, tile_x + tile_size):
            colors.append(pix[x,y])
    avg_color = calc_average_color(colors)
    color = 'rgb%r' % (avg_color,)
    box = [tile_x, tile_y, tile_x + tile_size, tile_y + tile_size]
    draw.rectangle(box, fill=color)

def calc_average_color(colors):
    tr, tg, tb = reduce(color_reducer, colors)
    total = len(colors)
    return tr/total, tg/total, tb/total

def color_reducer((r1, g1, b1), (r2, g2, b2)):
    return r1+r2, g1+g2, b1+b2


if __name__ == '__main__':
    logging.getLogger().setLevel(logging.INFO)
    image = Image.open('test.jpg')
    triangulize(image, 20)