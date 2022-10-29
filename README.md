# microservices

**Start Docker Compose**

```shellscript
docker compose up -d
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
docker compose down
docker compose rm -f
sleep 5
docker volume rm $(docker volume ls -q)
sleep 5
```
