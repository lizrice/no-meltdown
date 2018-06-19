#!/bin/bash
id=$(docker create $1)
echo $id
docker cp $id:/ms-out.html ms-out.html
docker rm -v $id
