.PHONY: postgres stop-docker-windows

postgres:
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:13-alpine

curl-login:
	curl --header "Content-Type: application/json" --request POST --data '{"email":"Roberto.Plaza@alu.uclm.es","password":"5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"}' http://localhost:8080/login

stop-docker-windows:
	docker ps -q | % { docker stop $_ }