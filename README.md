# Demo Smart Contracts for Pebble Data

## API server (golang)
Provide http request, accept the data submitted by thingsboard's api request, call the contract on the iotex blockchain, and hand the data to the contract for processing.

### Setup
Ready to configure

You need to provide these environment variables in the current operating environment or in a file:
```
IO_ENDPOINT="<iotex endpoint:port>"
CONTRACT_ADDRESS="<The address of the contract to be executed>"
VAULT_ADDRESS="<The address of the user who wants to execute the contract>"
VAULT_PASSWORD="<Password of the user to execute the contract>"
GAS_PRICE=<gas price>
GAS_LIMIT=<gas limit>
SLEEP_INTERVAL=<Waiting time for obtaining contract execution result>
```

And you need to provide The keystore file of the user who executed the contract

Example
```
UTC--<date time string>--<account address string>.json
```

Build and start operation, choose one
- Local Build:
```
make build
./bin/pebble-data-container
```

- Docker build:
```
docker build . -t <yourname>:<yourtag>
docker run -d -e <key>=<value> -e <key>=<value>... <yourname>:<yourtag>
or
docker run -d --env-file <envfile> <yourname>:<yourtag>
```

## Contract (solidity)
A contract that uses JsmnSolLib to parse json strings.


