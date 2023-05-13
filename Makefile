run-service:
	go run main.go -configPath=./config/ -configName=config.local.json

build-db:
	docker build -t interview-db ./int-db/

build-service:
	docker build -t interview-service .

run-db:
	docker run --name interview-db -d -p 27017:27017 interview-db

rm-db:
	docker stop interview-db 
	docker rm interview-db 