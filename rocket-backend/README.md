# Project rocket-backend

One Paragraph of project description goes here

## Getting Started
To get the backend up and running, you need to have the following installed:
- [Docker](https://www.docker.com/get-started/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install/source)

For the application to run, you need to create a .env file in the root directory of the backend with the following content:
```bash
PORT=8080
APP_ENV=local
BLUEPRINT_DB_HOST=postgres
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=blueprint
BLUEPRINT_DB_USERNAME=user
BLUEPRINT_DB_PASSWORD=password1234
BLUEPRINT_DB_SCHEMA=public

# generated with openssl rand -base64 64 (linux)
JWT_SECRET=

PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=admin
```

## How to run the backend:
If you want to run the backend completely isolated, you can run the complete application with Docker. This will run the backend and the database in a container.
If you want to run everything in a container, make sure that the `BLUEPRINT_DB_HOST` is set to `postgres` in the `.env` file. You can run the following command to start the application:
```bash

docker-compose up

```
If you only want the databases to run in a container,
make sure that the `BLUEPRINT_DB_HOST` is set to `localhost` in the `.env` file. You can run the following command to start the application:
```bash

docker-compose up postgres migrate pgadmin

```

## Docker Containers
1. **Postgres**: The database container.
2. **PgAdmin**: The database management tool.
3. **Backend**: The backend application container.
4. **Migrate**: The database migration tool.

### Postgres
- this container is used to run the database
- it uses the `postgis/postgis` image
- it supports the `postgis` extension
- `postgis` is helpful for geospatial queries

### PgAdmin
- this container is used to manage the database
- after the container is up, you can access it at `http://localhost:5050`
- enter the email and password from the `.env` file
- add a new server with the following settings:
  - Name: `Postgres`
  - Host: `postgres`
  - Port: `5432`
  - Username: `user`
  - Password: `password1234`
- you can now manage the database from the pgadmin interface


### Backend
- this container is used to run the backend application
- it uses the `golang:1.24.1-alpine` image
- it runs the application on port `8080`
- it uses the `migrate` tool to run the database migrations

### Migrate
- this container is used to run the database migrations
- it uses the files in the `migrations` folder to run the migrations

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
