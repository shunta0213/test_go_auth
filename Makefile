up:
	docker-compose up --build

down:
	docker-compose down
	docker-compose -f docker-compose.test.yaml down --volumes

core:
	docker attach test_go_auth-core-1


test:
	make uptest
	go test -v -cover ./...

uptest:
	docker-compose -f docker-compose.test.yaml up --build -d

downtest:
		docker-compose -f docker-compose.test.yaml down --volumes