
# RUN = docker exec -it server1

help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  guerrillad   to build the main binary for current platform"

rundocker:
	docker-compose -f ./docker/docker-compose.yml up -d

runserver1:
	docker exec -it server1 go run app/main.go

runserver2:
	docker exec -it server2 go run app/main.go

enters1:
	docker exec -it server1 bash
	# apt-get update
	# apt-get install telnet


enters2:
	docker exec -it server2 bash


sendmail:
	telnet 0.0.0.0:2525
