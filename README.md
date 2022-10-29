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
bash stop.sh
```
