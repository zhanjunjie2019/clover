#!/usr/bin/env bash

cd global
sh mod_tidy.sh
cd ..

cd share
sh mod_tidy.sh
cd ..

cd core
sh mod_tidy.sh
cd ..

cd starter-example
sh mod_tidy.sh
cd ..

cd starter-auth
sh mod_tidy.sh
cd ..
