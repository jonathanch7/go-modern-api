#!/bin/sh

swag init -g './cmd/restservice/server.go' --parseDependency --parseInternal -o './api/rest'
