FROM golang:1.22-alpine

WORKDIR /usr/src/note-api

RUN go install github.com/air-verse/air@latest

COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

EXPOSE 3000