FROM golang:1.13.5-stretch

WORKDIR apps/pebble-data-container

RUN apt-get install -y --no-install-recommends make

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build && \
    cp ./bin/pebble-data-container /usr/local/bin/pebble-data-container
CMD [ "pebble-data-container"]