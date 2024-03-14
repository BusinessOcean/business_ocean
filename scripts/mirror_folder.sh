#!/bin/bash

# Function to replicate folder structure
replicate_folder_structure() {
    source_dir=$1
    dest_dir=$2
    
    # Create destination directory if it doesn't exist
    mkdir -p "$dest_dir"
    
    # Use rsync to replicate folder structure
    rsync -av --include '*/' --exclude '*' "$source_dir/" "$dest_dir/"
}

# Main script
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 source_directory destination_directory"
    exit 1
fi

source_directory=$1
destination_directory=$2

replicate_folder_structure "$source_directory" "$destination_directory"
echo "Folder structure replication complete."
