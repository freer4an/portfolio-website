mongo:
	docker run --name mongo-docker -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=password mongo
test:
	go test -v -cover ./...
go:
	go run cmd/*