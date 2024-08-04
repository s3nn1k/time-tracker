build-image:
	docker build --no-cache ./ -t time-tracker:latest

run-container:
	docker-compose --env-file ./.env up