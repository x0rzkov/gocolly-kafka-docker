# reezocar_test

### Prerequisite

- GoLang 1.13
- Kafka 2.5.0
- Go Colly scrapper library
- Docker

### Initialisation

A) Create a _Cars_ Kafka topic :
 - `./kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic cars`

B) Use following Docker commands to run a MongoDB container :
 - Create a volume to persist data `docker volume create cars-volume`
 - Run a MongoDB container `docker run --name cars-mongodb -v cars-volume:/data/db -p 27017:27017 -d mongo` (for the sake of the POC we use no auth. We could have set environment variables _MONGO_INITDB_ROOT_USERNAME_ and _MONGO_INITDB_ROOT_PASSWORD_ to set credentials)
 
C) Download Go Colly library :
 - `go get -u github.com/gocolly/colly/...`