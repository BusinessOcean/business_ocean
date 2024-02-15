#!/bin/bash

# Function to recursively upgrade Go modules
upgrade_go_modules() {
  # Prompt user for Go version
  read -p "Enter the Go version (e.g., 1.17): " go_version
  if [ -z "$go_version" ]; then
    echo "Go version not provided. Exiting."
    exit 1
  fi

  # Update go.work file in the root directory
  sed -i "s/go[[:space:]]\+[0-9]\+\(\.[0-9]\+\)*$/go $go_version/" go.work

  # Upgrade Go modules
  for mod_file in $(find . -name "go.mod"); do
    dir=$(dirname "$mod_file")
    echo "Updating Go version in $dir"
    (cd "$dir" && go mod edit -go="$go_version")
  done
}

# Main execution
upgrade_go_modules
