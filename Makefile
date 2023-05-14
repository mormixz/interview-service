
build-image-db:
	docker build -t interview-db ./dockerfile/database/

run-docker-db:
	docker run --name interview-db -d -p 27017:27017 interview-db

rm-docker-db:
	docker stop interview-db 
	docker rm interview-db 

build-image-service:
	docker build -t interview-service -f ./dockerfile/service/dockerfile .
	
run-service:
	go run main.go -configPath=./config/ -configName=config.local.json
