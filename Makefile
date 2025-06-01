# Variables
APP_NAME=ssr-blog
PROXY_PORT=8080
GO_PORT=3001
SCSS_SRC = static/style.scss
CSS_OUT = static/style.css
CSS_MIN = static/style.min.css

# Runs the dev server with Templ + Go and auto-reload
dev:
	@echo "🚀 Starting development server on http://localhost:$(PROXY_PORT)"
	templ generate \
		-watch \
		-proxy="http://localhost:$(GO_PORT)" \
		-proxyport="$(PROXY_PORT)" \
		-open-browser=false \
		-cmd="go run ."

# Generate templates manually
generate:
	@echo "🛠️ Generating Templ components..."
	templ generate

# Run Go app directly (no proxy, no templ watch)
run:
	@echo "🎯 Running app directly at http://localhost:$(GO_PORT)"
	go run .

minify:
	sass static/style.scss | cleancss -o static/style.min.css

# Clean Templ-generated files
clean:
	@echo "🧹 Cleaning up..."
	find . -type f -name "*_templ.go" -delete
	rm -f $(CSS_OUT) $(CSS_MIN)

# Format all Go code
fmt:
	@echo "✨ Formatting Go code..."
	go fmt ./...

# Install Templ CLI
install-templ:
	go install github.com/a-h/templ/cmd/templ@latest

docker-build:
	@echo "🐳 Building Docker image locally..."
	docker build -t ssr-htmx-app .

docker-run:
	@echo "🚀 Running Docker container..."
	docker run -p 8080:8080 ssr-htmx-app
