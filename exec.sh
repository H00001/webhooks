#!/bin/bash
echo "push:$(date)" >> data
export GOPATH=$PWD/work
rm -rf ./work/src/gunplan.top/SCI
git clone git@github.com:H00001/SCI ./work/src/gunplan.top/SCI
cd ./work/src/gunplan.top/SCI/client
go build
mv client SCI_client_linux
scp SCI_client_linux root@123.56.223.150:/var/www/physics/static
cd ../../
rm -rf SCI

