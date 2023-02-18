# GovTech Technical Assessment

## Versions

This project uses Go 1.18.1 and MySQL 8.0.19

## Instructions on Running the Server

### Cloning the project
1. Clone the project onto your local machine
2. In the `src/` folder, run `go get` to install dependencies


### Setting up the DB

1. Create a MySQL database, and specify the following settings in the .env file

```
DBUSER=<fill in user>
DBPASS=<fill in password>
DBNET=<fill in protocol>
DBADDR=<fill in address>
DBNAME=<fill in database name>
```

For more details on filling in the database settings, check the documentation for the MySQL driver here: https://github.com/go-sql-driver/mysql


2. Install golang-migrate CLI. More detailed instructions can be found here for your specific operating system. https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```
brew install golang-migrate
```

3. In the `src/migrations` folder, run the following command to run the migrations.
```
migrate -path "./db_migrations" -database  "mysql://username:password@protocol(address)/dbname" up
```

The schema can be dropped by running the following command if need be.
```
migrate -path "./db_migrations" -database  "mysql://username:password@protocol(address)/dbname" drop
```

## Running the server

In the `src/` folder, run the command `go run main.go`

## Testing Instructions

1. Create a separate testing database and run migrations with the instructions above, but changing the database name to the testing database
2. Specify the following settings of the testing DB in the .env file

```
TEST_DBUSER=<fill in user>
TEST_DBPASS=<fill in password>
TEST_DBNET=<fill in protocol>
TEST_DBADDR=<fill in address>
TEST_DBNAME=<fill in database name>
```
3. In the `src/` folder, run the command `go test ./... -v`
