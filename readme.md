docker build -t loadtester .

docker run loadtester --url=http://example.com --requests=100 --concurrency=10
