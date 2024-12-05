
- run postgres

```
docker run -d --name=pg1 -p 5432:5432 -e POSTGRES_USER="admin" -e POSTGRES_PASSWORD="admin123" -e POSTGRES_DB="demodb" postgres
```

- adminer

```
docker -d --name dbui -p 8089:8080 adminer
```