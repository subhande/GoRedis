docker stop goredis
docker container rm goredis
docker build -t goredis .
# docker run -d -p 7379:7379 --net host --name goredis goredis
docker run -d -p 7379:7379 --name goredis goredis