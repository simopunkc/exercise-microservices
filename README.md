# microservices

**Tested**

- Go 1.19.2
- Docker Compose v2.12.2

**Start Docker Compose**

```shellscript
bash start.sh
```

**Rebuild image (optional)**

```shellscript
docker compose up -d --no-deps --build exercise-container user-container
```

**Stop image (optional)**

```shellscript
docker compose stop exercise-container user-container
```

**Teardown**

```shellscript
bash stop.sh
```
