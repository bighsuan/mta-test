
# RUN = docker exec -it server1
FROM-MTA = sender@mta-send
FROM-SWAKS = sender@swaks-send
TO = user@mta-receive


help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  build           to build three container, two run MTA server on localhost:2525, and one run nothing but install Swaks."
	@echo "  sendfrommta     to send mail from one of MTA server to another"
	@echo "  sendfromswaks   to send mail by Swaks, from the container which without MTA server"

build:
	docker-compose -f ./docker/docker-compose.yml up -d

sendfrommta:
	docker exec -it mta-send curl "localhost:8081/sendemail?from=$(FROM-MTA)&to=$(TO)"

sendfromswaks:
	docker exec -it swaks-send swaks --to $(TO) --port 2525