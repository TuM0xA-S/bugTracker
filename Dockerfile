FROM golang

WORKDIR /go/src/bugTracker
COPY . .

RUN go build .

CMD ["./bugTracker"]
