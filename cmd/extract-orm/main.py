"""
Note:
    This code was generated from ChatGPT being told to rewrite main.go.
"""
import os
import sys
from PIL import Image

def extract_orm_channels(file_path):
    try:
        # Open the image file
        img = Image.open(file_path)
        img = img.convert("RGBA")  # Ensure image has an alpha channel

        # Get image dimensions
        width, height = img.size

        # Prepare new images for each channel
        occlusion = Image.new("L", (width, height))
        roughness = Image.new("L", (width, height))
        metallic = Image.new("L", (width, height))

        # Extract the channels
        for y in range(height):
            for x in range(width):
                r, g, b, _ = img.getpixel((x, y))
                occlusion.putpixel((x, y), r)
                roughness.putpixel((x, y), g)
                metallic.putpixel((x, y), b)

        # Get base name and directory
        dir_name = os.path.dirname(file_path)
        base_name, _ = os.path.splitext(os.path.basename(file_path))

        # Save each channel as a separate image
        occlusion.save(os.path.join(dir_name, f"{base_name}.occlu.png"))
        roughness.save(os.path.join(dir_name, f"{base_name}.rough.png"))
        metallic.save(os.path.join(dir_name, f"{base_name}.metal.png"))

        print("Textures successfully unpacked:")
        print(f"- Occlusion: {base_name}.occlu.png")
        print(f"- Roughness: {base_name}.rough.png")
        print(f"- Metallic: {base_name}.metal.png")

    except Exception as e:
        print(f"Failed to process image {file_path}: {e}")


def main():
    if len(sys.argv) != 2:
        print("Usage: python orm_extract.py ./path-to-orm.png")
        return

    file_path = sys.argv[1]

    if not os.path.isfile(file_path):
        print(f"File not found: {file_path}")
        return

    extract_orm_channels(file_path)

if __name__ == "__main__":
    main()
