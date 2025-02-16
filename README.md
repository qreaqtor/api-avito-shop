# Backend-trainee-assignment-winter-2025

[Тех. задание](<https://github.com/avito-tech/tech-internship/blob/main/Tech Internships/Backend/Backend-trainee-assignment-winter-2025/Backend-trainee-assignment-winter-2025.md>)

## Deploy

Docker-compose файл лежит в папке deployment.
Для запуска можно использовать `make up` команду из Makefile - она запустит манифест с необходимыми переменными окружения.
Для завершения можно использовать также команду из Makefile `make down` или `make down-clear`, 
последняя очищает образы и тома связанные с текущим docker-compose файлом.

## API

Т.к. необходимо было реализовать API по предоставленной документации, то статус коды возвращаемых ответов были взяты оттуда, без изменения.
(В случае создания пользователя возвращается 200, вместо 201 с указанием на эндпоинт /api/info).

## Docs

Предоставленные файлы документации были скопированы и лежат в папке `./docs`, 
воспользоваться ими можно перейдя на [swagger editor](https://editor.swagger.io/) или запустив образ командой
```
docker run -p 8081:8080 -e SWAGGER_JSON=/api/schema.json -v $(pwd)/docs/schema.json:/api/schema.json swaggerapi/swagger-ui
```
