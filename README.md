## Packages
- Client : HTTP Client to make REST API calls with retry and back-off mechanism
- DB: Create and stores in-memory object later used to respond to GET requests to the service. Also performs caching for data less than a days old to avoid multiple network call to Alphavantage API
- Domain: Response schema
- DTO : Parser to deserialize data from Alphavantage API
- Handlers: Middleware to handle incoming GET request to service
- Routes: Route(s) management

## Main
`server.go`

## Docker
- To build the image locally run following command
`docker build --tag <TAG_NAME> .`
- To run without need to build the docker image `docker run --publish 8080:8080 -e SYMBOL='<TICKER-SYMBOL>' -e API_KEY='<API-KEY>' -e NDAYS='<NUM-DAYS>' stock-aggregator:latest`
- From another terminal `curl localhost:8080`
- Public Docker repo [https://hub.docker.com/r/abhibroto0402/stock-aggregator](https://hub.docker.com/r/abhibroto0402/stock-aggregator)

## Development Setup
### Running stock aggregator 
- Build the executable file and store under target dir `go build -o target/`
- ```
  $ export SYMBOL=MSFT
  $ export API_KEY=<ALPHAVANTAGE-API-KEY>
  $ export NDAYS=7
  $ ./target/stock-aggregator
  ```
### Test
`go test -v ./tests`
