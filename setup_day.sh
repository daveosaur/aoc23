#!/bin/bash
if [ $1 ] && [ ! -d $1 ]; then
  mkdir $1
  curl --cookie "session=$AOC_SESSION" https://adventofcode.com/2023/day/$1/input > $1/input.txt

  cd $1
  go mod init day_$1
  echo "package main

  import (
    \"fmt\"
    \"os\"
  )

  func main() {
  
  }" > main.go
elif [ $1 ]; then
  echo "don't overwrite days!"
else
  echo "enter a day to setup"
fi
