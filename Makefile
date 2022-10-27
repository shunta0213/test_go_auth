up:
	docker-compose up --build -d
	go run main.go

down:
	docker-compose down --rmi local
	docker-compose -f docker-compose.test.yaml down --volumes --rmi local

test:
	make uptest
	go test -v -cover ./...

uptest:
	docker-compose -f docker-compose.test.yaml up --build -d

downtest:
		docker-compose -f docker-compose.test.yaml down --volumes