#! /bin/bash

path="dayX"
if [ -z "$1" ]; then
  echo "no arg, using default path $path"
else
  path="$1"
fi

mkdir -p "$path"
cp -R ./utils/boilerplate/ "$path"
