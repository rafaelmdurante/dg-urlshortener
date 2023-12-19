up:
	docker compose up -d

down:
	docker compose down

run:
	go run cmd/api/main.go

logdb:
	docker logs -f urlshortener_db
