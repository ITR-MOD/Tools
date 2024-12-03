"""
Seperates an AORM/ORM image into its individual frames.

Args:
  image_path (str): The path to the image file to be resized.

Requires:
  pillow: This function requires the `pillow` library to be installed.
"""

import sys
import os
from PIL import Image


def unpack_orm_texture(file_path):
    """
    Unpack an ORM texture file into separate image files.

    Args:
        file_path (str): Path to the ORM texture file.
    """
    if not os.path.exists(file_path):
        print(f"File not found: {file_path}")
        return

    # Get the directory and base name of the file
    dir_name = os.path.dirname(file_path)
    base_name = os.path.splitext(os.path.basename(file_path))[0]

    try:
        # Load the ORM texture
        texture = Image.open(file_path).convert("RGB")

        # Extract channels
        occlusion, roughness, metallic = texture.split()

        # Save channels as individual images with the specified naming convention
        occlusion.save(os.path.join(dir_name, f"{base_name}.occlu.png"))
        roughness.save(os.path.join(dir_name, f"{base_name}.rough.png"))
        metallic.save(os.path.join(dir_name, f"{base_name}.metal.png"))

        print("Textures successfully unpacked:")
        print(f"- Occlusion: {base_name}.occlu.png")
        print(f"- Roughness: {base_name}.rough.png")
        print(f"- Metallic: {base_name}.metal.png")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python ./orm-extract.py ./path-to-orm.png")
        sys.exit(1)

    file_path = sys.argv[1]
    unpack_orm_texture(file_path)
