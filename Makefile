up:
	docker-compose up --build -d
	make core

down:
	docker-compose down
	docker-compose -f docker-compose.test.yaml down --volumes

core:
	docker attach core

test:
	make uptest
	go test -v -cover ./...

uptest:
	docker-compose -f docker-compose.test.yaml up --build -d

downtest:
		docker-compose -f docker-compose.test.yaml down --volumes