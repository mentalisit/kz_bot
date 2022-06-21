FROM golang:latest

COPY ./ ./
RUN go build -o kzbot .
CMD ["./kzbot"]