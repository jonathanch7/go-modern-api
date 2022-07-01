#!/bin/sh

enviroment="$1"

if [ $enviroment = 'dev' ]
then
  echo "running in mode development"
  ENV_FILE=.dev.env go run cmd/main.go
fi
