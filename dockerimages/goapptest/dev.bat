docker rm --force $(docker ps -a -q)
docker image rm --force  $(docker image ls -q)
cd frontend
go build main.go
cd ..
docker-compose up
