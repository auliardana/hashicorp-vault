# FROM golang:alpine AS builder
# RUN apk update && apk add --no-cache git
# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build -o main ./Api/main.go

# FROM alpine
# WORKDIR /app
# COPY --from=builder /app/main /app
# COPY --from=builder /app/.env /app
# CMD [ "./main" ]