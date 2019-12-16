#!/bin/bash

while read -r a b
do
  if [[ $a -gt "$b" ]]
  then
    echo "X is greater than Y"
  fi

  if [[ $a -lt "$b" ]]
  then
    echo "X is less than Y"
  fi

  if [[ $a -eq "$b" ]]
  then
    echo "X is equal to Y"
  fi
done <<< "$(< /dev/stdin tr '\n' ' ')"
