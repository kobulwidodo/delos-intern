# Delos Internship Submission
## Framework
- Http Web : Gin
- ORM Database : Gorm
- Database : PostgreSQL
- Mock : Mockery
- Testing : Testify
## Architecture
![Diagram](https://user-images.githubusercontent.com/53964878/172319411-468d8866-f2c6-4026-a371-b8e47714ed97.png)

This Project has 4 Domain layer :
- Models
- Repository
- Usecase
- Handler

Repository -> Usecase -> Handler
## How to run this project
```bash
# Clone project
$ git clone https://github.com/kobulwidodo/delos-intern.git

# Move to project
$ cd delos-intern

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
