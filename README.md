## Versions

Go 1.18.1



## Instructions on Running the Server


### Setting up the DB

1. Install MySQL CLI
2. Create a database called edusystem with the command `CREATE DATABASE edusystem;`
3. Install golang-migrate CLI. More detailed instructions can be found here for your specific operating system. https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```
brew install golang-migrate
```

4. From the project route, enter the migrations folder. It should contain another folder called migrations. Use the following command to run the migrations.
```
migrate -path "./migrations" -database  "mysql://root@tcp(localhost:3306)/edusystem" up
```


In case of error, the schema can be dropped by running the following command
```
migrate -path "./migrations" -database  "mysql://root@tcp(localhost:3306)/edusystem" drop
```

