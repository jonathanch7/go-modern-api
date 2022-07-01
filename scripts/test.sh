#!/bin/sh

typeTest="$1"
coverMode="$2"
if [ -z "$typeTest" ]
  then
    typeTest='none'
fi
if [ -z "$coverMode" ]
  then
    coverMode='none'
fi

if [ $typeTest = 'unit' ]
then
  echo "running unit test..."
  if [ $coverMode = 'coverHTML' ]
  then
    go test -count=1 -covermode=count -coverprofile=coverage.txt $(go list ./internal/...); go tool cover -html=coverage.txt -o coverage.html
  else
    if [ $coverMode = 'cover' ]
    then
      go test -count=1 -covermode=count -coverprofile=coverage.txt $(go list ./internal/...)
    else
      go test -count=1 $(go list ./internal/...)
    fi
  fi
fi

if [ $typeTest = 'it' ]
then
  echo "running integration test..."
  ENV_FILE=$(pwd)/.it.env go test -count=1 $(go list ./test/integration/...)
fi

if [ $typeTest = 'e2e' ]
then
  echo "running end to end test..."
  ENV_FILE=$(pwd)/.e2et.env go test -count=1 $(go list ./test/e2e/...)
fi
