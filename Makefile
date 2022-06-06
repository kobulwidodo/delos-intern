BINARY=engine
engine:
	go build -o ${BINARY} app/*.go

run-db:
	docker-compose up --build -d

run-app:
	go run app/http.go

stop:
	docker-compose down

test: 
	go test -v -cover -covermode=atomic ./...
