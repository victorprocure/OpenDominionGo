.PHONY: build-local build templ notify-templ-proxy tailwind test run

-include .env
-include .env.local

build-local:
	@go build -o bin/main.exe main.go

build:
	@npm run build
	@go build -o bin/main.exe main.go

templ:
	@go tool templ generate --watch --proxy=http://localhost:$(APP_PORT) --proxyport=8081 --proxybind="0.0.0.0" --open-browser=false

notify-templ-proxy:
	@go tool templ generate --notify-proxy --proxyport=8081

test:
	@go test -v ./...

air-build:
	@make notify-templ-proxy
	@npm run build
	@go build -o bin/main.exe main.go

run:
	@make templ & sleep 1
	@air

tailwind:
	@npx @tailwindcss/cli -i components/css/styles.css -o assets/styles.css --watch

