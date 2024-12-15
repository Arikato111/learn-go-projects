default:
	@echo no default command

run:
	go run .

watch:
	nodemon --exec go run . --signal SIGTERM -e go,html

up:
	docker compose up -d

down:
	docker compose down --rmi local
