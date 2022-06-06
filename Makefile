BINARY=engine
engine:
	go build -o ${BINARY} app/*.go

docker:
	docker build -t delos-intern .

run:
	docker-compose up --build -d

stop:
	docker-compose down

test: 
	go test -v -cover -covermode=atomic ./...
