FROM golang:alpine AS build

WORKDIR /app

#COPY .kafka ./.kafka
#COPY .postgresql ./.postgresql

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
#COPY pkg/ ./pkg/

RUN go build -o fare-parser cmd/app/main.go

FROM alpine

WORKDIR /app
ENV PROJECT_ABS_PATH="/app"

COPY .env .
COPY --from=build /app/fare-parser .

CMD ["./fare-parser"]
