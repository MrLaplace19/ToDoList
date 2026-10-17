# Details

Date : 2026-10-17 09:52:49

Directory /Users/mrlaplace/git_rep/REST

Total : 90 files,  3671 codes, 28 comments, 714 blanks, all 4413 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Jenkinsfile](/Jenkinsfile) | Groovy | 30 | 0 | 6 | 36 |
| [Makefile](/Makefile) | Makefile | 63 | 0 | 21 | 84 |
| [README.md](/README.md) | Markdown | 0 | 0 | 1 | 1 |
| [cmd/todoapp/Dockerfile](/cmd/todoapp/Dockerfile) | Docker | 10 | 0 | 1 | 11 |
| [cmd/todoapp/main.go](/cmd/todoapp/main.go) | Go | 75 | 5 | 17 | 97 |
| [docker-compose.jenkins.yaml](/docker-compose.jenkins.yaml) | YAML | 15 | 0 | 0 | 15 |
| [docker-compose.yaml](/docker-compose.yaml) | YAML | 53 | 0 | 4 | 57 |
| [docs/docs.go](/docs/docs.go) | Go | 125 | 2 | 5 | 132 |
| [docs/swagger.json](/docs/swagger.json) | JSON | 107 | 0 | 0 | 107 |
| [docs/swagger.yaml](/docs/swagger.yaml) | YAML | 71 | 0 | 1 | 72 |
| [go.mod](/go.mod) | Go Module File | 41 | 0 | 4 | 45 |
| [go.sum](/go.sum) | Go Checksum File | 124 | 0 | 1 | 125 |
| [internal/core/domain/nullable.go](/internal/core/domain/nullable.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/domain/statistics.go](/internal/core/domain/statistics.go) | Go | 8 | 0 | 2 | 10 |
| [internal/core/domain/task.go](/internal/core/domain/task.go) | Go | 146 | 0 | 27 | 173 |
| [internal/core/domain/uninitialized.go](/internal/core/domain/uninitialized.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/domain/user.go](/internal/core/domain/user.go) | Go | 90 | 0 | 24 | 114 |
| [internal/core/domain/user\_test.go](/internal/core/domain/user_test.go) | Go | 61 | 0 | 6 | 67 |
| [internal/core/errors/common.go](/internal/core/errors/common.go) | Go | 7 | 0 | 3 | 10 |
| [internal/core/logger/config.go](/internal/core/logger/config.go) | Go | 20 | 0 | 8 | 28 |
| [internal/core/logger/logger.go](/internal/core/logger/logger.go) | Go | 73 | 0 | 18 | 91 |
| [internal/core/repository/postgres/pool/errors.go](/internal/core/repository/postgres/pool/errors.go) | Go | 7 | 0 | 2 | 9 |
| [internal/core/repository/postgres/pool/pgx/adapters.go](/internal/core/repository/postgres/pool/pgx/adapters.go) | Go | 46 | 0 | 10 | 56 |
| [internal/core/repository/postgres/pool/pgx/config.go](/internal/core/repository/postgres/pool/pgx/config.go) | Go | 28 | 0 | 8 | 36 |
| [internal/core/repository/postgres/pool/pgx/pool.go](/internal/core/repository/postgres/pool/pgx/pool.go) | Go | 63 | 0 | 15 | 78 |
| [internal/core/repository/postgres/pool/pool.go](/internal/core/repository/postgres/pool/pool.go) | Go | 24 | 0 | 5 | 29 |
| [internal/core/transport/http/middelware/common.go](/internal/core/transport/http/middelware/common.go) | Go | 77 | 0 | 19 | 96 |
| [internal/core/transport/http/middelware/dummy.go](/internal/core/transport/http/middelware/dummy.go) | Go | 17 | 0 | 4 | 21 |
| [internal/core/transport/http/middelware/middleware.go](/internal/core/transport/http/middelware/middleware.go) | Go | 12 | 0 | 4 | 16 |
| [internal/core/transport/http/request/decode.go](/internal/core/transport/http/request/decode.go) | Go | 30 | 0 | 10 | 40 |
| [internal/core/transport/http/request/path\_values.go](/internal/core/transport/http/request/path_values.go) | Go | 28 | 0 | 6 | 34 |
| [internal/core/transport/http/request/query\_params.go](/internal/core/transport/http/request/query_params.go) | Go | 42 | 0 | 7 | 49 |
| [internal/core/transport/http/response/errors.go](/internal/core/transport/http/response/errors.go) | Go | 5 | 0 | 1 | 6 |
| [internal/core/transport/http/response/handler.go](/internal/core/transport/http/response/handler.go) | Go | 72 | 0 | 23 | 95 |
| [internal/core/transport/http/response/writer.go](/internal/core/transport/http/response/writer.go) | Go | 25 | 0 | 7 | 32 |
| [internal/core/transport/http/server/config.go](/internal/core/transport/http/server/config.go) | Go | 24 | 0 | 8 | 32 |
| [internal/core/transport/http/server/route.go](/internal/core/transport/http/server/route.go) | Go | 25 | 0 | 6 | 31 |
| [internal/core/transport/http/server/router.go](/internal/core/transport/http/server/router.go) | Go | 36 | 0 | 9 | 45 |
| [internal/core/transport/http/server/server.go](/internal/core/transport/http/server/server.go) | Go | 86 | 0 | 18 | 104 |
| [internal/core/transport/http/types/nullable.go](/internal/core/transport/http/types/nullable.go) | Go | 27 | 10 | 8 | 45 |
| [internal/features/statistics/repository/postgres/get\_statistics.go](/internal/features/statistics/repository/postgres/get_statistics.go) | Go | 71 | 0 | 12 | 83 |
| [internal/features/statistics/repository/postgres/model.go](/internal/features/statistics/repository/postgres/model.go) | Go | 31 | 0 | 7 | 38 |
| [internal/features/statistics/repository/postgres/repository.go](/internal/features/statistics/repository/postgres/repository.go) | Go | 12 | 0 | 4 | 16 |
| [internal/features/statistics/service/get\_statistics.go](/internal/features/statistics/service/get_statistics.go) | Go | 66 | 0 | 14 | 80 |
| [internal/features/statistics/service/service.go](/internal/features/statistics/service/service.go) | Go | 24 | 0 | 5 | 29 |
| [internal/features/statistics/transport/http/get\_statistics.go](/internal/features/statistics/transport/http/get_statistics.go) | Go | 72 | 0 | 16 | 88 |
| [internal/features/statistics/transport/http/transport.go](/internal/features/statistics/transport/http/transport.go) | Go | 35 | 0 | 7 | 42 |
| [internal/features/tasks/repository/postgres/create\_task.go](/internal/features/tasks/repository/postgres/create_task.go) | Go | 75 | 0 | 11 | 86 |
| [internal/features/tasks/repository/postgres/delete\_task.go](/internal/features/tasks/repository/postgres/delete_task.go) | Go | 25 | 0 | 6 | 31 |
| [internal/features/tasks/repository/postgres/get\_task.go](/internal/features/tasks/repository/postgres/get_task.go) | Go | 54 | 0 | 8 | 62 |
| [internal/features/tasks/repository/postgres/get\_tasks.go](/internal/features/tasks/repository/postgres/get_tasks.go) | Go | 49 | 0 | 9 | 58 |
| [internal/features/tasks/repository/postgres/models.go](/internal/features/tasks/repository/postgres/models.go) | Go | 31 | 0 | 7 | 38 |
| [internal/features/tasks/repository/postgres/patch\_task.go](/internal/features/tasks/repository/postgres/patch_task.go) | Go | 78 | 0 | 11 | 89 |
| [internal/features/tasks/repository/postgres/repository.go](/internal/features/tasks/repository/postgres/repository.go) | Go | 12 | 0 | 5 | 17 |
| [internal/features/tasks/service/create\_task.go](/internal/features/tasks/service/create_task.go) | Go | 19 | 0 | 6 | 25 |
| [internal/features/tasks/service/delete\_task.go](/internal/features/tasks/service/delete_task.go) | Go | 17 | 0 | 4 | 21 |
| [internal/features/tasks/service/get\_task.go](/internal/features/tasks/service/get_task.go) | Go | 19 | 0 | 4 | 23 |
| [internal/features/tasks/service/get\_tasks.go](/internal/features/tasks/service/get_tasks.go) | Go | 32 | 0 | 6 | 38 |
| [internal/features/tasks/service/patch\_task.go](/internal/features/tasks/service/patch_task.go) | Go | 33 | 0 | 6 | 39 |
| [internal/features/tasks/service/service.go](/internal/features/tasks/service/service.go) | Go | 39 | 0 | 6 | 45 |
| [internal/features/tasks/transport/http/create\_task.go](/internal/features/tasks/transport/http/create_task.go) | Go | 45 | 0 | 15 | 60 |
| [internal/features/tasks/transport/http/delete\_task.go](/internal/features/tasks/transport/http/delete_task.go) | Go | 28 | 0 | 6 | 34 |
| [internal/features/tasks/transport/http/dto\_common.go](/internal/features/tasks/transport/http/dto_common.go) | Go | 34 | 0 | 11 | 45 |
| [internal/features/tasks/transport/http/get\_task.go](/internal/features/tasks/transport/http/get_task.go) | Go | 31 | 0 | 7 | 38 |
| [internal/features/tasks/transport/http/get\_tasks.go](/internal/features/tasks/transport/http/get_tasks.go) | Go | 47 | 0 | 11 | 58 |
| [internal/features/tasks/transport/http/patch\_task.go](/internal/features/tasks/transport/http/patch_task.go) | Go | 80 | 0 | 19 | 99 |
| [internal/features/tasks/transport/http/transport.go](/internal/features/tasks/transport/http/transport.go) | Go | 70 | 0 | 5 | 75 |
| [internal/features/users/repository/postgres/create\_user.go](/internal/features/users/repository/postgres/create_user.go) | Go | 34 | 0 | 10 | 44 |
| [internal/features/users/repository/postgres/delete\_user.go](/internal/features/users/repository/postgres/delete_user.go) | Go | 25 | 0 | 6 | 31 |
| [internal/features/users/repository/postgres/get\_user.go](/internal/features/users/repository/postgres/get_user.go) | Go | 46 | 0 | 9 | 55 |
| [internal/features/users/repository/postgres/get\_users.go](/internal/features/users/repository/postgres/get_users.go) | Go | 45 | 0 | 8 | 53 |
| [internal/features/users/repository/postgres/model.go](/internal/features/users/repository/postgres/model.go) | Go | 20 | 0 | 4 | 24 |
| [internal/features/users/repository/postgres/patch\_user.go](/internal/features/users/repository/postgres/patch_user.go) | Go | 49 | 0 | 7 | 56 |
| [internal/features/users/repository/postgres/repository.go](/internal/features/users/repository/postgres/repository.go) | Go | 12 | 0 | 4 | 16 |
| [internal/features/users/service/create\_user.go](/internal/features/users/service/create_user.go) | Go | 19 | 0 | 5 | 24 |
| [internal/features/users/service/create\_user\_test.go](/internal/features/users/service/create_user_test.go) | Go | 34 | 0 | 7 | 41 |
| [internal/features/users/service/delete\_user.go](/internal/features/users/service/delete_user.go) | Go | 14 | 0 | 3 | 17 |
| [internal/features/users/service/get\_user.go](/internal/features/users/service/get_user.go) | Go | 16 | 0 | 5 | 21 |
| [internal/features/users/service/get\_users.go](/internal/features/users/service/get_users.go) | Go | 30 | 0 | 7 | 37 |
| [internal/features/users/service/patch\_user.go](/internal/features/users/service/patch_user.go) | Go | 24 | 0 | 6 | 30 |
| [internal/features/users/service/service.go](/internal/features/users/service/service.go) | Go | 37 | 0 | 5 | 42 |
| [internal/features/users/transport/http/create\_user.go](/internal/features/users/transport/http/create_user.go) | Go | 40 | 11 | 14 | 65 |
| [internal/features/users/transport/http/delete\_user.go](/internal/features/users/transport/http/delete_user.go) | Go | 27 | 0 | 7 | 34 |
| [internal/features/users/transport/http/dto\_common.go](/internal/features/users/transport/http/dto_common.go) | Go | 23 | 0 | 7 | 30 |
| [internal/features/users/transport/http/get\_user.go](/internal/features/users/transport/http/get_user.go) | Go | 31 | 0 | 11 | 42 |
| [internal/features/users/transport/http/get\_users.go](/internal/features/users/transport/http/get_users.go) | Go | 40 | 0 | 13 | 53 |
| [internal/features/users/transport/http/path\_user.go](/internal/features/users/transport/http/path_user.go) | Go | 76 | 0 | 15 | 91 |
| [internal/features/users/transport/http/transport.go](/internal/features/users/transport/http/transport.go) | Go | 68 | 0 | 7 | 75 |
| [migrations/000001\_.down.sql](/migrations/000001_.down.sql) | MS SQL | 3 | 0 | 0 | 3 |
| [migrations/000001\_.up.sql](/migrations/000001_.up.sql) | MS SQL | 26 | 0 | 3 | 29 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)