Build the Go application with the following:

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

scratch is empty, so there are no libraries and no loadpath for it to look in. What we have to do is modify our build script to statically compile our app with all libraries built in


before running again after you change the containers run:
docker-compose.exe  down --rmi all

or

docker rm $(docker ps -a -q)
 docker image rm $(docker image ls -q)

To remove old images from cache
