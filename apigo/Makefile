build:
	@go build -o bin/apigo src/main.go

test:
	@go test -v ./...

run: build
	@./bin/apigo

# Cria uma nova migração de banco de dados utilizando no cli make migration <nome_da_migração>
migration:
	@migrate create -ext sql -dir src/migrate/migrations $(filter-out $@,$(MAKECMDGOALS)) 

migrate-up:
	@go run src/migrate/main.go up

migrate-down:
	@go run src/migrate/main.go down