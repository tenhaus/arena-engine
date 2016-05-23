#!/bin/bash

export BOTPIT_ENV="development"
export GOOGLE_APPLICATION_CREDENTIALS="./botpit-development-authentication.json"
go run Server.go
