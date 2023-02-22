#!/usr/bin/env bash

docker build --build-arg pgName=starter-example --build-arg httpPort=8800 -t clover-example .
docker build --build-arg pgName=starter-auth --build-arg httpPort=8810 -t clover-auth .

docker rmi `docker images | grep  '<none>' | awk '{print $3}'`

function pause(){
  read -n 1 -p "$*" INP
  if [ $INP != '' ] ; then
          echo -ne '\b \n'
  fi
}

pause 'Press any key to continue...'