# Aplikasi Forum Sederhana

## Cara menjalankan :
1. Buat Docker Container untuk database mysql menggunakan docker compose

    ```
    docker-compose up
    ```

2. Jalankan Docker container
    ```
    docker container start simple-forum-sql
    ```

3. Lakukan database migrations
    ```
    make migrate-up
    ```

4. Jalankan project
    ```
    go run cmd/main.go
    ```
    atau
    ```
    make gorun
    ```

## Tech Stack:
- Go
- Gin (http framework for golang)
- Go Viper (library for config management)
- Golang JWT
- Zerolog (logger)
- MySQL