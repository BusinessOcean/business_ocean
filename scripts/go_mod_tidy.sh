#!/bin/bash

# Function to run go mod tidy in a directory if go.mod exists
run_go_mod_tidy() {
  if [ -f "go.mod" ]; then
    echo "Running go mod tidy in $(pwd)"
    go mod tidy
  fi
}

# Function to recursively search for directories containing go.mod files
search_and_run_go_mod_tidy() {
  for dir in "$1"/*; do
    if [ -d "$dir" ]; then
      (cd "$dir" || exit
       run_go_mod_tidy
       search_and_run_go_mod_tidy "$dir")
    fi
  done
}

# Start searching from the current directory
search_and_run_go_mod_tidy .

echo "Finished running go mod tidy in all folders containing go.mod files."
