.PHONY: postgres stop-docker-windows

postgres:
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:13-alpine

stop-docker-windows:
	docker ps -q | % { docker stop $_ }