* Тестовое задание на позицию Junior Backend Developer

**Используемые технологии:**

- Go
- JWT
- MongoDB

**Задание:**

Написать часть сервиса аутентификации.

Два REST маршрута:

- Первый маршрут выдает пару Access, Refresh токенов для пользователя сидентификатором (GUID) указанным в параметре запроса
- Второй маршрут выполняет Refresh операцию на пару Access, Refreshтокенов

**Требования:**

Access токен тип JWT, алгоритм SHA512, хранить в базе строго запрещено.

Refresh токен тип произвольный, формат передачи base64, хранится в базеисключительно в виде bcrypt хеша, должен быть защищен от изменения настороне клиента и попыток повторного использования.

Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.

Сервер генерирует пару токенов и отсылает на клиент (клиент может их хранить где угодно). Клиент мониторит состояние acces токена по его времени жизни, как только он истечёт, клиент отправляет запрос на обновление ключей.
Сервер проверяет все сессии пользователя и если не находит, то делает logout и возвращает ошибку, если находит, то делает refresh и отсылает на клиент.
