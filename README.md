# Fullstack Go - Template

This is a template for a fullstack Go web applications. I'm using this
for my own projects.

- [echo](https://echo.labstack.com/) for web routing capabilities
- [templ](https://github.com/a-h/templ) as a template engine
- [sqlx](https://github.com/jmoiron/sqlx) for database access
- [validator](https://github.com/go-playground/validator) for validation
- [air](https://github.com/cosmtrek/air) for auto code reloading during dev
- [go-migrate](https://github.com/golang-migrate/migrate) for migrations
- Tailwind for styling
- `make` for commands (build, test, watch, ...)

<img width="1003" alt="image" src="https://github.com/rverton/fullstack-go/assets/1506585/8d92d592-2a14-4ee9-b0c3-f6442965311d">

## Development

This project is setup with and in-memory sqlite3 DB. Migrations are auto run on start. 

For development, a simple `make watch` will run all necessary watch tasks:
templ, air and tailwind. The dev server is then restarted, as soon as a change
happens.

If a persistent DB is needed, the `make migrate` command can be used.

Additional useful commands:

- `make test` - run all tests
- `make build` - build
- `make run` - run 
- `make audit` - audit code
