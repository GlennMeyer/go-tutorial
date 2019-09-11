build:
	docker build -t go_tutorial .
compose:
	docker-compose up --build --force-recreate
debug:
	docker run --rm -v ${CURDIR}:/usr/src/myapp -w /usr/src/myapp --name debug -it golang:1.13.0 /bin/bash
killapi:
	docker kill go_tutorial_api_1
killdb:
	docker kill go_tutorial_db_1
killdebug:
	docker kill debug
killrun:
	docker kill api
push:
	docker push glennmeyer/go_tutorial:latest
run:
	docker run --rm -p 80:8080 --name api go_tutorial
stopapi:
	docker stop go_tutorial_api_1
stopdb:
	docker stop go_tutorial_db_1
stopdebug:
	docker stop debug
stoprun:
	docker stop api
tag:
	docker tag go_tutorial glennmeyer/go_tutorial:latest
