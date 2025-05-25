```
$ ./build.sh
```

# Docker
```
$ docker compose -f infra/docker-compose.yaml up -d
```

# Makefile
```
$ make build       # Compila
$ make run         # Executa
$ make up          # Sobe Docker
$ make clean       # Remove binários
$ make rebuild     # Clean + Build
$ make help        # Mostra ajuda
```