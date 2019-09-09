build:
	docker build -t go-tutorial .
debug:
	docker run -v ${CURDIR}:/root --name debug -it --rm golang:1.13.0 /bin/bash
stopd:
	docker stop debug
killd:
	docker kill debug
run:
	docker run -p 80:8080 --name go --rm go-tutorial
stopr:
	docker stop go
killr:
	docker kill go