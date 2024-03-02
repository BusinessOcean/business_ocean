#!/bin/bash

# Function to recursively find directories containing go.mod files
find_go_mod_directories() {
    # Find directories containing go.mod files recursively
    find . -type f -name "go.mod" -exec dirname {} \;
}

# Loop through each directory containing a go.mod file
while IFS= read -r folder; do
    # Check if the folder exists
    if [ -d "$folder" ]; then
        echo "Running linter in $folder"
        # Run the Go linter command, replace 'your-linter-command' with the actual linter command you use
        # Example: golint "$folder"
        golangci-lint run "$folder"
    else
        echo "Warning: $folder doesn't exist."
    fi
done < <(find_go_mod_directories)
