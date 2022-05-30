# Code Flow 

# Docker
- To build the image locally run following command
`docker build --tag <TAG_NAME> .`
- To run without need to build the docker image `docker run --publish 8080:8080 -e SYMBOL='<TICKER-SYMBOL>' -e API_KEY='<API-KEY>' -e NDAYS='<NUM-DAYS>' stock-aggregator:latest`
  - From another terminal `curl localhost:8080/stock-aggregate`