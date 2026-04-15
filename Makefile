.PHONY: dev build css tailwind clean

## Start the dev server and Tailwind watcher together
dev:
	@./tailwindcss -i internal/web/static/css/input.css -o internal/web/static/css/style.css --watch &
	@$(shell go env GOPATH)/bin/air

## Build a production binary with minified CSS
build: css
	go build -o server ./cmd/server

## Build minified CSS only
css:
	./tailwindcss -i internal/web/static/css/input.css -o internal/web/static/css/style.css --minify

## Download the Tailwind CLI binary for this platform (Linux x64)
tailwind:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

## Remove build artifacts
clean:
	rm -f server internal/web/static/css/style.css
