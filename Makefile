run-interactive-os:
	go build .
	export MODE="dev" && go run . interactive os
pg-refresh:
	docker-compose -f tools/docker/docker-compose.yaml stop
	docker-compose -f tools/docker/docker-compose.yaml rm -f
	docker-compose -f tools/docker/docker-compose.yaml up -d
pg-connect:
	docker container exec -it pg-jarvis-test bash -c 'PGPASSWORD=mamadspass psql -U mamad -h 30.0.0.10 mamad_db'
pg-restart:
	docker-compose -f tools/docker/docker-compose.yaml restart
pg-start:
	docker-compose -f tools/docker/docker-compose.yaml up -d