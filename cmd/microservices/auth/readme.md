## Сервис регистрации и аутентификации
Аутентификация с помощью JWT-токена.

[Swagger](https://github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/blob/polina-auth/docs/api/openapi.yaml)

#### Запуск
```bash
cd cmd/microservices/auth
go run main.go
```

### Структура
- `main.go` - файл запуска сервиса
- `api` - "_уровень представлений_": ручки, роутинг
  - `api/test` - быстрый тест на доступность, не судите строго
  - `api/models` - фронтовые модели
- `behavior` - "_бизнес-логика_", логика авторизации
- `utils` - утилиты
  - `utils/logging.go` - оболочка хэндлера для логирования
- `errors` - используемые ошибки
### TODO:
- services
  - [ ] добавить модуль работы с JWT
  - [ ] выбрать способ хэширования
  - [ ] сделать регистрацию
  - [ ] сделать аутентификацию

- storage
  - [ ] добавтить интерфейс взаимодействия с бд
  - [ ] интегрировать по интерфейсу заглушку

- transport
  - [ ] сделать валидацию данных
  - [ ] парсинг данных из тела запроса

### Политика
__Требования к валидации данных__ - [тут](https://github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/blob/polina-auth/docs/%D0%BF%D0%BE%D0%BB%D0%B8%D1%82%D0%B8%D0%BA%D0%B0%20%D0%B0%D1%83%D1%82%D0%B5%D0%BD%D1%82%D0%B8%D1%84%D1%82%D0%BA%D0%B0%D1%86%D0%B8%D0%B8.md)

- [ ] дописать политику
  - процесс регистрации
  - процесс аутентификации
  - структура токена

___
by `PtFux`

#### Ресурсы:
- логгирование с `slog` - [тут](https://habr.com/ru/companies/slurm/articles/798207/)
- роутинг с `gorilla/mux` - [тут](https://habr.com/ru/companies/ruvds/articles/561108/)
