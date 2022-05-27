# BQ Connect

A small connector service for querying a BigQuery dataset 

### State
:construction: Project currently a work in progress

### Infrastructure
[Infrastructure of the project](https://github.com/VinceDeslo/bq-connect-infra)

### Roadmap
- [x] Hexagonal architecture
- [x] HTTP API
- [x] Dockerfile and Docker Compose up and running
- [ ] gRPC API
- [ ] Redis as a caching layer

### Purpose
This connector acts as a microservice to fetch and parse data from a BigQuery data lake.
The data lake contains a dataset of video game sales accross the years.