# Go-book-review
書評APIをGo言語で開発

## preparation
### prepare .env file
```
AWS_ACCESS_KEY_ID=<<prepare your aws access key id>>
AWS_SECRET_ACCESS_KEY=<<preapre your aws secret access key>>
AWS_REGION=ap-northeast-1
AWS_COGNITO_CLIENT_ID=<<prepare your aws cognito client id>>
AWS_COGNITO_USER_POOL_ID=<<prepare your aws cognito user pool id>>
AWS_COGNITO_USER_POOL_ISS=<<prepare your aws cognito user pool iss>>

APP_PORT=8080
APP_HOST=localhost

DB_HOST=db
DB_PASS=passw0rd
```

### dependencies
```
go mod tidy
```

## serve
### run MySQL using Docker (terminal 1)
```
docker compose up
```

### run Go API server (terminal 2)
```
export $(cat .env | grep -v ^#)
go run cmd/main.go
```
