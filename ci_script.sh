#!/usr/bin/env bash

docker build -t snabar/api-service .
echo "$(tput setaf 2)Building image $(tput sgr0)"
    docker stop snabar-api-service
    echo "$(tput setaf 1)Stopping Docker image $(tput sgr0)"
    docker rm snabar-api-service
    echo "$(tput setaf 1)Removing Docker image $(tput sgr0)"
docker run --name snabar-api-service -d -p 8081:8081 snabar/api-service
echo "$(tput setaf 2)Running container for the first time $(tput sgr0)"

if [ $? != 0 ]
then
    docker stop snabar-api-service
    echo "$(tput setaf 1)Stopping Docker image $(tput sgr0)"
    docker rm snabar-api-service
    echo "$(tput setaf 1)Removing Docker image $(tput sgr0)"
    docker run -d --rm --name snabar-api-service -p 8081:8081 -d snabar/api-service
    echo "$(tput setaf 1)Starting new Docker image $(tput sgr0)"
fi

echo "$(tput setaf 2)Deployed Coming-Soon \n Please visit http://snabar.com $(tput sgr0)"