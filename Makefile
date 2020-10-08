.PHONY: postgres stop-docker-windows

postgres:
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:13-alpine

curl-login:
	curl --header "Content-Type: application/json" --request POST --data '{"email":"Roberto.Plaza@alu.uclm.es","password":"6cf615d5bcaac778352a8f1f3360d23f02f34ec182e259897fd6ce485d7870d4"}' http://localhost:8080/login

curl-get:
	curl --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJvYmVydG8uUGxhemFAYWx1LnVjbG0uZXMiLCJleHAiOjE2MDIxNTc0MzcsIm9yaWdfaWF0IjoxNjAyMTUzODM3fQ.MeMuGFwZI1XcMqMzXq_h7V626udN1jZtDD6uSgkhK5M" --request GET http://localhost:8080/difficulties/

stop-docker:
	docker stop $(docker ps -aq)