# imageUpload S3 [Backend Application]

## Build & Run (Locally)
### Prerequisites
- go 1.17
- docker & docker-compose

Create .env file in root directory and add following values:
```dotenv
export DB_HOST=db
export DB_PORT=5432
export DB_USERNAME=postgres
export DB_NAME=postgres
export DB_SSLMODE=disable
export DB_PASSWORD=

export SALT=<random string>
export SECRET=<random string>

export AWS_REGION=eu-central-1
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
```

Use `make run` to build&run project.
