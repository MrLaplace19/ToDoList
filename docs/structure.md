# Граф структур проекта

Схемы составлены по структурам и конструкторам в исходном коде.
`main` в `cmd/todoapp/main.go` создаёт пул, репозитории, сервисы,
обработчики и HTTP-сервер, затем регистрирует маршруты `/api/v1`.

## Основные зависимости

Сплошная стрелка показывает зависимость от поля или вызов при регистрации
маршрутов. Пунктирная стрелка с подписью «реализует» показывает реализацию
Go-интерфейса. Имена интерфейсов дополнены пакетом, чтобы отличать их
от одноимённых структур.

```mermaid
flowchart TB
    Server["server.HTTPServer"] -->|регистрирует| Router["server.APIVersionRouter"]
    Server --> Logger["logger.Logger"]
    Server --> Config["server.Config"]
    Router -->|регистрирует Router с методами обработчика| UH
    Router -->|регистрирует Router с методами обработчика| TH
    Router -->|регистрирует Router с методами обработчика| SH

    subgraph Users["internal/features/users"]
        UH["http.UserHttpHandler"] --> UI["http.UserService · интерфейс"]
        US["service.UsersService"] -. реализует .-> UI
        US --> URI["service.UsersRepository · интерфейс"]
        UR["postgres.UsersRepository"] -. реализует .-> URI
    end

    subgraph Tasks["internal/features/tasks"]
        TH["http.TasksHTTPHandler"] --> TI["http.TasksService · интерфейс"]
        TS["service.TasksService"] -. реализует .-> TI
        TS --> TRI["service.TasksRepository · интерфейс"]
        TR["postgres.TasksRepository"] -. реализует .-> TRI
    end

    subgraph Statistics["internal/features/statistics"]
        SH["http.StatisticsHTTPHandler"] --> SI["http.StatisticsService · интерфейс"]
        SS["service.StatisticsService"] -. реализует .-> SI
        SS --> SRI["service.StatisticsRepository · интерфейс"]
        SR["postgres.StatisticsRepository"] -. реализует .-> SRI
    end

    UR --> PI["pool.Pool · интерфейс"]
    TR --> PI
    SR --> PI
    P["pgx.Pool"] -. реализует .-> PI
    P -->|встраивает| Native["pgxpool.Pool"]
    Native --> DB[(PostgreSQL)]
```

Обработчики хранят интерфейсы сервисов, сервисы хранят интерфейсы
репозиториев, репозитории хранят интерфейс `pool.Pool`. Конкретные
реализации передаются через конструкторы в `main`.
`APIVersionRouter` регистрирует функции обработчиков через `Router`;
отдельных полей для обработчиков в нём нет.

## Доменные структуры

```mermaid
classDiagram
    class User {
        ID int
        Version int
        FullName string
        PhoneNumber *string
        Validate() error
        ApplyPatch(UserPatch) error
    }
    class Task {
        ID int
        Version int
        Title string
        Description *string
        Completed bool
        CompletedAt *time.Time
        CreatedAt time.Time
        AuthorUserID int
        Validate() error
        ApplyPatch(PatchTask) error
        CompletedDuration() *time.Duration
    }
    class Statistics {
        CreatedTask int
        CompletedTask int
        CompletedTaskRate *float64
        ArgTime *time.Duration
    }
    class UserPatch {
        FullName Nullable~string~
        PhoneNumber Nullable~string~
        Validate() error
    }
    class PatchTask {
        Title Nullable~string~
        Description Nullable~string~
        Completed Nullable~bool~
        Validate() error
    }
    class Nullable~T~ {
        Value *T
        Set bool
    }
    User "1" <-- "0..*" Task : AuthorUserID соответствует User.ID
    User ..> UserPatch : ApplyPatch
    Task ..> PatchTask : ApplyPatch
    UserPatch *-- Nullable~T~ : поля
    PatchTask *-- Nullable~T~ : поля
```

Связь `Task → User` реализована идентификатором `AuthorUserID`, а не
вложенной структурой `User`. В БД это внешний ключ `tasks.author_id →
users.id` с `ON DELETE CASCADE`.

`StatisticsService` получает `[]domain.Task` из своего репозитория и
вычисляет `domain.Statistics`. Отдельной таблицы статистики нет.
`Nullable[T]` различает отсутствующее поле (`Set == false`), явный `null`
(`Set == true`, `Value == nil`) и переданное значение.

## Преобразование данных

```mermaid
flowchart LR
    JSON["JSON запроса"] --> Request["HTTP Request DTO"]
    Request --> Domain["domain.User / Task / UserPatch / PatchTask"]
    Domain --> Service["Сервис: проверка и применение изменений"]
    Service --> Repository["PostgreSQL-репозиторий"]
    Repository --> Model["UserModel / TaskModel"]
    Model --> Result["domain.User / Task"]
    Result --> Response["HTTP Response DTO"]
    Response --> Output["JSON ответа"]
```

Схема преобразования обобщает операции создания, чтения и изменения.
Удаление возвращает ошибку или успешный статус без доменного результата.
Для статистики сервис преобразует список задач в `domain.Statistics`,
а обработчик — в `GetStatisticsResponse`.
