#!/bin/bash

# Function to traverse directories recursively
traverse() {
    local folder="$1"
    for dir in "$folder"/*; do
        if [ -d "$dir" ]; then
            if [ -f "$dir/go.mod" ]; then
                go work use "$dir"
            else
                traverse "$dir"
            fi
        fi
    done
}

# Start traversal from the current directory
traverse "."
