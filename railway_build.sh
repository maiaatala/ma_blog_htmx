#!/bin/bash
echo "📦 Installing Templ and Generating Components..."
go install github.com/a-h/templ/cmd/templ@latest
templ generate

echo "🛠️ Building Go App..."
go build -o blog-app

