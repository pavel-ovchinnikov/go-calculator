

run:
	$(info Run all dockers)
	docker-compose up --force-recreate --build --remove-orphans
