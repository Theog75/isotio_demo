#!/bin/bash
`docker container stop $(docker container ls -a -q)`
`docker container rm $(docker container ls -a -q |grep -v $(docker container ls -a |grep mongo|awk '{print $1}'))`
`docker image rm $(docker image ls -q|grep -v $(docker image ls|grep mongo|awk '{print $3}'))`
