#!/usr/bin/env bash

protoc --gofast_out=plugins=grpc:. --micro_out=. *.proto

function pause(){
  read -n 1 -p "$*" INP
  if [ $INP != '' ] ; then
          echo -ne '\b \n'
  fi
}

pause 'Press any key to continue...'
