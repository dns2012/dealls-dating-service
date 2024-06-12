generate:
	./buf.gen.sh

migrate:
	docker-compose exec -it httpd go run database/migration.go

start:
	docker-compose up -d
	docker-compose logs -f httpd

stop:
	docker-compose down

restart:
	docker-compose restart httpd
	docker-compose logs -f httpd

mock:
	./mock.sh

test:
	./test.sh