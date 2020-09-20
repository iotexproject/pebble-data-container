FROM golang:1.13.5-stretch

WORKDIR apps/iotex-blockchain-iot

RUN apt-get install -y --no-install-recommends make

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build && \
    cp ./bin/iotex-blockchain-iot /usr/local/bin/iotex-blockchain-iot
CMD [ "iotex-blockchain-iot"]