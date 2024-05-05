### challenge-stress-test

# para buildar esta imagem
```
docker build -t loadtester .
```

# para realizar os testes:
```
docker run loadtester --url=http://google.com --requests=200 --concurrency=10
```

# para remover esta imagem docker basta rodar:
```
docker rm -f $(docker ps -aqf "ancestor=loadtester")
docker rmi loadtester
```
