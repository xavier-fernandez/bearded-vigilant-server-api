#!/usr/bin/env bash
function printSlashes {
   echo "---------------------------------------------------"
}
function printDone {
   echo "DONE"
}
printSlashes
echo "Downloading Environ data Gatherer dependencies..."
printSlashes

# Downloads the GORM persistence Framework: https://github.com/jinzhu/gorm
echo "Downloading GORM persistance framework..."
go get -u -v "github.com/jinzhu/gorm"
printDone
printSlashes

# Downloads the HWIO I/O library. https://github.com/mrmorphic/hwio
echo "Downloading the HWIO I/O library..."
go get -u -v "github.com/mrmorphic/hwio"
printDone
printSlashes
echo "All the libraries have been downloaded successfully."