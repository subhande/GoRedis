clear
redis-benchmark -n 10000 -t ping_mbulk -c 1 -h localhost -p 6379
redis-benchmark -n 10000 -t ping_mbulk -c 1 -h localhost -p 7379