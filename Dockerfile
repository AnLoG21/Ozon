FROM golang:1.19
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["go", "run", "main.go"]
