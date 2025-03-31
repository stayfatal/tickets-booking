all:

prod:
	docker-compose $(filter-out $@,$(MAKECMDGOALS))

test:
	docker-compose -f docker-compose.test.yml $(filter-out $@,$(MAKECMDGOALS))

pipe:
	docker compose -f docker-compose.test.yml up postgres redis migrations -d && docker compose -f docker-compose.test.yml up auth gateway test --exit-code-from test