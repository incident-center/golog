# golog

Данный модуль представляет собой универсальный логгер для приложений на Go. Он инкапсулирует в себе функционал библиотеки [uber-go/zap](https://github.com/uber-go/zap) и предоставляет более упрощённый интерфейс для работы с логированием.

## Использование
Сначала необходимо инстанциировать объект логгера при помощи функции New:

```go
logger := logger.New("debug")
```

Здесь `debug` - это уровень логирования. Вместо него может быть указано любое другое значение из: `debug`, `info`, `warn`, `error` или `fatal`.
Далее можно использовать этот объект для ведения логов:

```go
context := map[string]any{
    "user_id": 123,
    "role": "admin",
}

logger.Debug("User logged in", context)
```

В данном примере мы использовали метод Debug, но также доступны методы Info, Warn, Error и Fatal.

Каждый метод принимает два аргумента: сообщение в виде строки и контекст в виде словаря.

## Лицензия
Этот проект лицензирован по лицензии MIT - подробности в файле LICENSE.md, в корне проекта.