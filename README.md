# Todo-App
Working with the gin-gonic/gin framework.

Dependency injection technique.

Working with the Postgresql database. Launching from Docker. 

Generation of migration files.

Application configuration using the spf13/viper library. Working with environment variables.

Working with the database using the sqlx library.

Writing SQL queries

Registration and authentication. Working with JWT. Middleware.


# View data from docker container
docker ps
docker exec -it <container_id> /bin/bash
psql -U postgres

# Create a new migration
migrate create -ext sql -dir < directory> -seq <name>

# Run migration Up/Down
migrate -path <directory> -database <database_url> up/down
