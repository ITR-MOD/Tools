import sys
from PIL import Image

def downscale_image(input_path, scale_factor=0.5):
    # Open the image
    img = Image.open(input_path)
    
    # Calculate the new dimensions
    new_width = int(img.width * scale_factor)
    new_height = int(img.height * scale_factor)
    
    # Downscale the image
    img_resized = img.resize((new_width, new_height), Image.LANCZOS)
    
    # Create the output path with .ds.png suffix
    output_path = f"{input_path.rsplit('.', 1)[0]}.hs.png"
    
    # Save the downscaled image
    img_resized.save(output_path)
    print(f"Image saved to {output_path}")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py path/to-image.png")
        sys.exit(1)

    # Get the input image path from command-line arguments
    input_path = sys.argv[1]
    
    # Downscale the image
    downscale_image(input_path)
