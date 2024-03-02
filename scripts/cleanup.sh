#!/bin/bash

cleanup_file=".cleanup"

# Check if .cleanup file exists
if [ ! -f "$cleanup_file" ]; then
    echo "Error: .cleanup file not found."
    exit 1
fi

# Read targets from .cleanup file and delete them
while IFS= read -r target || [[ -n "$target" ]]; do
    # Remove leading and trailing whitespace
    target="$(echo -e "${target}" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')"
    
    # Delete the target
    if [[ "$target" == **/* ]]; then
        # If the target has **/* pattern, use find command for deletion
        find . -path "./${target}" -exec rm -rf {} +
        echo "Deleted: $target"
    else
        # If the target is a directory with / prefix, remove the prefix
        if [[ "$target" == /* ]]; then
            target="${target:1}"
        fi
        
        # Check if target exists and delete it
        if [ -e "$target" ]; then
            rm -rf "$target"
            echo "Deleted: $target"
        fi
    fi
done < "$cleanup_file"
