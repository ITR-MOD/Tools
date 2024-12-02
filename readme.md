# **ITR Tools**

A collection of utilities designed for extracting assets and modding game files, primarily for **ITR2**.

The tools here are primarily written in **Go**, but tools in other languages can be made and added to this repo. (or their own in this org)

> All tools can be adapted for other games and projects freely.

## Quick Terms

- **Language:** The programming language the tool is written in.
- **Type:** Defines the tool format:
  - **Compiled:** Executable files (e.g., .exe).
  - **Script:** Scripts, typically in Python or similar languages.
  - **Addon (for):** Extensions for editors or tools, such as:
    - **Addon (Unreal):** For Unreal Engine.
    - **Addon (Blender):** For Blender.

## **Tools**

- **Extract ORM**
  - **Language:** Go
  - **Type:** Compiled
  - **Description:** Extracts Occlusion, Roughness, and Metallic maps from channel-packed ORM files.
  - [Learn More](cmd/extract-orm/)
