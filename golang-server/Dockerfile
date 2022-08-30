# Stage 1

FROM golang:1.17-alpine AS builder

ARG DB_USER=${DB_USER}
ARG DB_PASS=${DB_PASS}
ARG DB_NAME=${DB_NAME}
ARG DB_HOST=${DB_HOST}
ARG DB_PORT=${DB_PORT}
ARG AMQP_URL=${AMQP_URL}

ENV DB_USER ${DB_USER}
ENV DB_PASS ${DB_PASS}
ENV DB_NAME ${DB_NAME}
ENV DB_HOST ${DB_HOST}
ENV DB_PORT ${DB_PORT}
ENV AMQP_URL ${AMQP_URL}

RUN apk update && apk add --no-cache

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o main main.go

# Stage 2

FROM scratch AS production

WORKDIR /app

COPY --from=builder /src/ .

# EXPOSE 3334

CMD [ "./main" ]