# 🏗 Stage 1: Build
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copy static assets into builder image
COPY ./static ./static

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN CGO_ENABLED=0 go build -o /server main.go

# 🧼 Stage 2: Minimal final image
FROM gcr.io/distroless/base-debian11 AS final

COPY --from=builder /server /server
COPY --from=builder /app/static /static

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/server"]

