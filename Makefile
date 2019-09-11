build:
	docker build -t go_tutorial .
compose:
	docker-compose up --build --force-recreate
debug:
	docker run -v ${CURDIR}:/root --name debug -it --rm golang:1.13.0 /bin/bash
killapi:
	docker kill go_tutorial_api_1
killdebug:
	docker kill debug
killrun:
	docker kill api
push:
	docker push glennmeyer/go_tutorial:latest
run:
	docker run -p 80:8080 --name api --rm go_tutorial
stopdebug:
	docker stop debug
stoprun:
	docker stop api
tag:
	docker tag go_tutorial glennmeyer/go_tutorial:latest
