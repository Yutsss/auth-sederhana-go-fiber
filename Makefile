APP_NAME=backend

dep:
	go mod tidy

run:
	GO_ENV=development go run main.go

dev:
	CompileDaemon -command="go run main.go" -build="go build -o $(APP_NAME)"

build:
	go build -o $(APP_NAME)

migrate-dev:
	GO_ENV=development  go run main.go --migrate

migrate-prod:
	GO_ENV=production  go run main.go --migrate

run-build:
	GO_ENV=development ./$(APP_NAME)

run-start-dev:
	GO_ENV=development ./$(APP_NAME)

run-start-prod:
	GO_ENV=production ./$(APP_NAME)
deploy-dev:
	pm2 start ecosystem.config.js --env development

deploy-prod:
	pm2 start ecosystem.config.js --env production

clean:
	rm -f $(APP_NAME)