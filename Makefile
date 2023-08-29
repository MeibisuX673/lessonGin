
test:
	docker-compose -f docker-compose.test.yml up --abort-on-container-exit --force-recreate || docker-compose -f docker-compose.test.yml down