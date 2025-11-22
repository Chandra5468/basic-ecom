# basic-ecom
tiago : https://youtu.be/s3XItrqfccw

Using sqlc - a sql code generator
/*
    You write plain SQL queries in .sql files.
    You define your database schema.
    sqlc then generates type-safe Go code (DAO functions) based on your SQL queries and schema. This generated code provides the interfaces to execute those queries from your Go application.
*/

Migrations
    - Package

    - goose -s create create_products sql

After adding migrations file
- Do sqlc generate


# COMMANDS ORDER OF EXECUTION

1. install sqlc and goose using go install commands

2. in root -> Create sqlc.yaml

    2.2. In internal/adapters/postgresql
            Add two folders manually
                sqlc and migrations
                In sqlc folder 
                    add this file queries.sql
    2.3. In queries.sql file add commands in specific pattern which should be viable for sqlc.
        :many means we get an array in generated go.
        :one means we get single result in generated go.

    2.4. use command 
            goose -s create file_name sql. To generate a migration file. Drag and drop into migrations folder.

            Add mirgations up and down for like create and drop a table. In a goose specific pattern.

3. In sqlc.yaml mention
    queries path : ./internal/adapters/postgresql/sqlc/queries.sql (Sqlc turns sql mentioned in queries.sql into go methods)
    schema path
    out path
    sql_package
    emit_json_tags
    emit_json_interface


4. Now enter sqlc generate command
    Based on 
        internal/postgresql/sqlc/queries.sql
        and
        internal/postgresql/mirgations/...

    Go files will be generated based on sqlc.yaml configuration.

5. Add these in .env
    GOOSE_DBSTRING="host=localhost user=postgres password=user dbname=postgres sslmode=disable"
    GOOSE_DRIVER=postgres
    GOOSE_MIGRATION_DIR=./internal/adapters/postgresql/migrations

6. In new terminal or if server not runnning in same terminal do
    goose up

7. Install this for all dependency and migrations to work
    github.com/jackc/pgx/v5

---------------------------------
    For new tables

1. goose -s create create_orders sql

2. goose up

3. sqlc generate