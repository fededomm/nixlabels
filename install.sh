#!/bin/bash

user="fededomm"

# Nome del file Go
file="main.go"

# Nome del file eseguibile
output="nixlbl"

# Directory di destinazione
dest="/home/fededomm/.local/bin"

# Compila il file Go
go build -o $output $file

# Verifica se la compilazione è andata a buon fine
if [ $? -eq 0 ]; then
    echo "Compilation successful, moving executable to $dest"
    # Sposta l'eseguibile nella directory di destinazione
    mv $output $dest
else
    echo "Compilation failed"
fi
