import os
from PIL import Image

def remove_blue_channel(img):
    """Remove the blue channel from the image."""
    width, height = img.size
    no_blue = Image.new("RGBA", img.size)

    for y in range(height):
        for x in range(width):
            r, g, b, a = img.getpixel((x, y))
            no_blue.putpixel((x, y), (0, g, b, a))

    return no_blue

def remove_green_channel(img):
    """Remove the green channel from the image."""
    width, height = img.size
    no_green = Image.new("RGBA", img.size)

    for y in range(height):
        for x in range(width):
            r, g, b, a = img.getpixel((x, y))
            no_green.putpixel((x, y), (r, 0, b, a))

    return no_green

def remove_blue_channel(img):
    """Remove the blue channel from the image."""
    width, height = img.size
    no_blue = Image.new("RGBA", img.size)

    for y in range(height):
        for x in range(width):
            r, g, b, a = img.getpixel((x, y))
            no_blue.putpixel((x, y), (r, g, 0, a))

    return no_blue

def save_image(img, file_path):
    """Save the image to the specified file path."""
    img.save(file_path)
    print(f"Image written to: {file_path}")
