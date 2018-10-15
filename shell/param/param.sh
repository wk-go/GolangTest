#!/usr/bin/env bash
echo "\$*:"
for arg in "$*"
do
    echo $arg
done

echo "\$@:"
for arg in $@
do
    echo $arg
done
echo $[ $@[1:*] ]
