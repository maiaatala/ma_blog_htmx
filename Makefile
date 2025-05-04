# Variables
APP_NAME=ssr-blog
PROXY_PORT=8080
GO_PORT=3001

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

# Clean Templ-generated files
clean:
	@echo "🧹 Cleaning up..."
	find . -type f -name "*_templ.go" -delete

# Format all Go code
fmt:
	@echo "✨ Formatting Go code..."
	go fmt ./...

# Install Templ CLI
install-templ:
	go install github.com/a-h/templ/cmd/templ@latest

