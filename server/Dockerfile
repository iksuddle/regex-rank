FROM golang:1.24.5-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

EXPOSE 8080

CMD [ "make run" ]
