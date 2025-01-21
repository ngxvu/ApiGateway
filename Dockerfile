FROM golang:1.22-alpine AS builder

WORKDIR /3rd-party-gateway/

COPY . .

# Print Go environment and module information
RUN go env
RUN go mod tidy
RUN go mod download

# Build the Go application
RUN go build -o 3rd-party-gateway.bin .

ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_PORT
ARG DB_SSLMODE
ARG SERVICE_A_API_KEY

ENV DB_HOST=$DB_HOST
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME
ENV DB_PORT=$DB_PORT
ENV DB_SSLMODE=$DB_SSLMODE
ENV GOONG_API_KEY=$SERVICE_A_API_KEY

RUN chmod +x 3rd-party-gateway.bin

EXPOSE 8081

CMD ["./3rd-party-gateway.bin"]
