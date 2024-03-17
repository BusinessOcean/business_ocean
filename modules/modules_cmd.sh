#!/bin/bash

# Function to run a command in a directory if go.mod exists
run_command() {
  if [ -f "go.mod" ]; then
    echo -e "\e[1;32mRunning $1 in $(pwd)\e[0m"
    echo -e "\e[1;34mCommand Output:\e[0m"
    eval "echo -e '\e[1;35m'; $1; echo -e '\e[0m'"  # Command output in magenta color
  fi
}

# Function to recursively search for directories containing go.mod files
search_and_run_command() {
  for dir in "$1"/*; do
    if [ -f "$dir/go.mod" ]; then
      (cd "$dir" || exit
       run_command "$2")
    elif [ -d "$dir" ]; then
      search_and_run_command "$dir" "$2"
    fi
  done
}

# Ask user for the command to be run
read -p "Enter the command to run: " cmd

# Start searching from the current directory
search_and_run_command . "$cmd"

echo "Finished running $cmd in all folders and subfolders containing go.mod files."
