# SimpleAuth

## Сборка
### Локальная:
Небходимо изменить переменную окружения DB_URL и иметь запущенную базу данных PostgreSQL.
```sh
make build
make run
```
### В Docker:
```sh
docker-compose up --build
```