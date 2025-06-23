###### [_↩ Back to `main` branch_](https://github.com/cuongpiger/golang)

<hr>

- Run postgreqSQL database in Docker:

  ```bash
  docker run --name postgres-db \
    -e POSTGRES_USER=developer \
    -e POSTGRES_PASSWORD=password123 \
    -e POSTGRES_DB=mydb \
    -p 30432:5432 \
    -d postgres:16
  ```

- Export the database schema to HCL format using Atlas:

  ```bash
  atlas schema inspect \
    --url "postgres://developer:password123@127.0.0.1:30432/mydb?sslmode=disable" \
    --format "{{ hcl . }}" > schema.hcl
  ```
