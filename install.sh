#!/bin/bash

user="fededomm"

# file to compile
file="main.go"

# executable name
output="nixlbl"

# destination directory
dest="/home/$user/.local/bin"

# Compile go file
go build -o $output $file

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "Compilation successful, moving executable to $dest"
    # Move executable to destination directory
    mv $output $dest
else
    echo "Compilation failed"
fi
