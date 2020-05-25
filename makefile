dev:
	go run main.go

build-image:
	docker build -t marcfinserv .

run-local: build-image server

server:
	docker run -p 80:80
