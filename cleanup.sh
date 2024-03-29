#! /bin/bash

#Killing containers which has label of "hkn"

#for CONTAINER in $(docker ps -q --filter "label=hkn")
#do
#	docker kill $CONTAINER
#	docker rm $CONTAINER
#done

docker kill $(docker ps -q --filter "label=hkn")
#time docker ps --no-trunc --filter "label=hkn" --format '{{.ID}}' | xargs -n 1 docker inspect --format '{{.State.Pid}}' $1 | xargs -n 1 sudo kill -9
# Removing killed containers which have label of "hkn"

docker rm $(docker ps -q -a --filter "label=hkn" --filter status=exited)
# Prune Network bridges which are not used by any container

docker network prune -f

#Remove wireguard config and interfaces


for INTERFACE in $(sudo wg | grep interface: | awk '{print $2}')
do
	echo $INTERFACE
	sudo wg-quick down $INTERFACE
done

rm -rf /etc/wireguard/*

rm -rf filetransfer/*
