## SimpleAuth

### Описание
Простенький сервис по генерации токенов для JWT авторизации

### Сборка
`❗ Команда выполняется из корневой директории проекта`
```sh
docker-compose up -d --build
```
`Сервер будет доступен по адресу localhost:8080`

### API
- `POST /api/token` - генерирует пару `Access` и `Refresh` токенов. Обязательно указать параметр `guid` содержащий guid юзера
  - Пример: `localhost:8080/api/token?guid=fi-zz-bu-zz`
- `POST /api/refresh` - обновляет `Access` токен. Обязательно указать параметр `guid` содержащий guid юзера.
В теле запроса указать пару `Access` и `Refresh` токенов.
```json
{
  "access_token": "your old access token",
  "refresh_token": "your refresh token"
}
```