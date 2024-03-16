#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 destination_folder"
    exit 1
fi

destination_folder=$1

# Create a temporary directory
temp_dir=$(mktemp -d)

# Clone the repository into the temporary directory
git clone https://github.com/Business-Ocean/protobuf "$temp_dir"

# Copy the contents of the subfolder directories to the specified destination folder
cp -r "$temp_dir"/*/ "$destination_folder"

# Remove the temporary directory
rm -rf "$temp_dir"

echo "Repository cloned and subfolder contents copied to $destination_folder"
