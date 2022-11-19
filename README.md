
# Driver Service for Taxi app

Auth/taxi ordering


## Installation

Clone repository

```bash
git clone github.com/zxcghoulhunter/InnoTaxi-Driver
```

Install dependencies
```bash
go mod tidy
```

Run application
```bash
go run main.go
```

OR

Use docker
```bash
docker compose up
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file
- Cassandra for users data storage
`DB_HOST`(localhost)
`DB_PORT`(9842) 

- secret key for JWT auth
`SECRET`(secret) 

- App gRPC and REST ports
`RPC_PORT`(4343)
`REST_PORT`(8081)


`JWT_EXPIRATION_TIME` Expiration time for JWT token, example:2h, 1d, 30m 

