.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i static/css/tailwind.css -o static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i static/css/tailwind.css -o static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go
