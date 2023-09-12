run-interactive-os:
	go run . interactive os
run-interactive-pgs:
	make pg-start
	go run . interactive os

dev-run-interactive-os:
	export MODE="dev" && go run . interactive os
dev-run-interactive-pgs:
	export MODE="dev" && go run . interactive postgres ./configs/postgres-config.json
	make pg-start

pg-refresh:
	docker-compose -f tools/docker/docker-compose.yaml stop
	docker-compose -f tools/docker/docker-compose.yaml rm -f
	docker-compose -f tools/docker/docker-compose.yaml up -d
pg-start:
	docker-compose -f tools/docker/docker-compose.yaml up -d
pg-attach:
	docker container exec -it pg-jarvis-test bash -c 'PGPASSWORD=mamadspass psql -U mamad -h 30.0.0.10 mamad_db'
pg-restart:
	docker-compose -f tools/docker/docker-compose.yaml restart
