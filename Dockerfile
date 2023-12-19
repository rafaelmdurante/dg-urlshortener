FROM golang:1.21-alpine3.18 as builder
LABEL authors="rafaelmdurante" title="DevGym URL Shortener" version="0.1.0"

WORKDIR /app
COPY . .

# When -d is used, go get will only manage dependencies in go.mod
# The ... expands to subdirectories, like ** in bash. More: https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Build executable for linux as in Dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM scratch

WORKDIR /
COPY --from=builder /app/api ./
EXPOSE 3000
ENTRYPOINT ["./api"]

