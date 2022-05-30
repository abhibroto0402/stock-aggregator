FROM golang:1.16-alpine
WORKDIR /app/stock-aggregator
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/server .
CMD [ "./out/server" ]
EXPOSE 8080