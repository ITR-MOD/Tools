"""
Note:
    This code was generated from ChatGPT being told to rewrite main.go.
"""
import os
import sys
from PIL import Image



def invert_image_colors(file_path):
    try:
        # Open the image
        img = Image.open(file_path)
        img = img.convert("RGBA")  # Ensure the image has an alpha channel

        # Invert colors
        inverted_img = Image.new("RGBA", img.size)
        width, height = img.size

        for y in range(height):
            for x in range(width):
                r, g, b, a = img.getpixel((x, y))
                inverted_img.putpixel((x, y), (255 - r, 255 - g, 255 - b, a))

        # Save the inverted image
        output_path = f"{os.path.splitext(file_path)[0]}.invert.png"
        inverted_img.save(output_path)
        print(f"Image written to: {output_path}")

    except Exception as e:
        print(f"Failed to process image {file_path}: {e}")


def main():
    if len(sys.argv) <= 1:
        print("Usage: python invert_image.py path/to-image.png")
        return

    # Process each provided file path
    for file_path in sys.argv[1:]:
        if not os.path.isfile(file_path):
            print(f"Invalid file path: {file_path}")
            continue

        invert_image_colors(file_path)

if __name__ == "__main__":
    main()
