docker build -f Dockerfile -t ascii-art-web-docker .

docker container run -p 8080:8080 --detach --name ascii-art-web-docker ascii-art-web-docker

docker rmi -f golang:1.22-alpine

docker rmi -f alpine

docker system prune -f