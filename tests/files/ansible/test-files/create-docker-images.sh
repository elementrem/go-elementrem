#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="elementrem/cppjit-testrunner" docker-cppjit
docker build --tag="elementrem/python-testrunner" docker-python
docker build --tag="elementrem/go-testrunner" docker-go
