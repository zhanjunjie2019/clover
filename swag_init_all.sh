#!/usr/bin/env bash

cd starter-example
swag init --pd
cd ..

cd starter-auth
swag init --pd
cd ..

function pause(){
  read -n 1 -p "$*" INP
  if [ $INP != '' ] ; then
          echo -ne '\b \n'
  fi
}

pause 'Press any key to continue...'
