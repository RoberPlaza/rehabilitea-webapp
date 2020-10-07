.PHONY: postgres stop-docker-windows

postgres:
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:13-alpine

curl-login:
	curl --header "Content-Type: application/json" --request POST --data '{"email":"Roberto.Plaza@alu.uclm.es","password":"6cf615d5bcaac778352a8f1f3360d23f02f34ec182e259897fd6ce485d7870d4"}' http://localhost:8080/login

curl-get:
	curl --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJvYmVydG8uUGxhemFAYWx1LnVjbG0uZXMiLCJleHAiOjE2MDIwODc2MzksIm9yaWdfaWF0IjoxNjAyMDg0MDM5fQ.iM6T675XDIIuMe7i2dp1yM1JEVKBX5QcpJvGudDA7Yk" --request GET http://localhost:8080/events

stop-docker-windows:
	docker ps -q | % { docker stop $_ }