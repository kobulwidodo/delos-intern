# Delos Internship Submission
## Framework
- Http Web : Gin
- ORM Database : Gorm
- Database : PostgreSQL
- Mock : Mockery
- Testing : Testify
## Architecture
This Project has 4 Domain layer :
- Models
- Repository
- Usecase
- Handler

Repository -> Usecase -> Handler
## How to run this project
```bash
# Build Database with compose
$ make run-db

# wait until postgre container run perfectly
# check container status
$ docker ps

# Run test
$ make test

# Run App
$ make run-app
```
## Endpoint Documentation
Postman : https://documenter.getpostman.com/view/14494329/Uz5JHFBj
## Bug Found
- sometimes test failed to run because mysql >= 200ms when truncate table executed
