#!/bin/bash


j=$1
for ((i=100; i<=j; i++))
do
echo "num is $i"

git tag -d v1.0.$i 
git push origin :refs/tags/v1.0.$i
done
