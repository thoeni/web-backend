#!/usr/bin/env bash

while [[ $# > 1 ]]
do
key="$1"

case $key in
    -p|--password)
    MARIADB_PWD="$2"
    echo $MARIADB_PWD
    shift
    ;;
    *)
    echo "Parameter $1 unknown"
    exit 1
    ;;
esac
shift
done

if [ -z $MARIADB_PWD ]; then
    echo "Please, specify a valid password for MariaDB providing the parameter -p password"
    exit 1
fi

docker build --rm --tag web-backend . &&
docker run --rm -p 8080:8080 --name web-backend --link web-mariadb:web-mariadb -e MARIADB_PWD=$MARIADB_PWD -e MARIADB_CONTAINER="web-mariadb" web-backend