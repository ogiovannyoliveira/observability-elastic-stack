# Stage 1

FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o main main.go

# Stage 2

FROM scratch AS production

WORKDIR /app

COPY --from=builder /src/ .

EXPOSE 3334

CMD [ "./main" ]