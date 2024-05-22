# Tinder Like Apps


### Prerequisite

- Git (See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))
- Go 1.22 or later (See [Golang Installation](https://golang.org/doc/install))
- [create migration](https://github.com/golang-migrate/migrate) installed

### How To Create Migrations

  ```sh 
    migrate create -ext sql -dir files/db/migrations -seq your_file_name
  ```

### Setup Development Database

  ```sh
  PGPASSWORD=postgres psql -h localhost -U postgres -w -c "create database database_name;"
  ```


### Run Migration

```sh 
    export POSTGRESQL_URL='' && migrate -database ${POSTGRESQL_URL} -path files/db/migration
  ```


### Run Test

```sh 
    go test ./...
  ```

### Run Application Locally

```sh 
    go run cmd/api/main.go
  ```


