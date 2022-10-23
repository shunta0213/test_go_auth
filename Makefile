up:
	docker-compose up --build

down:
	docker-compose down

core:
	docker attach test_go_auth-core-1
