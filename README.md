# BQ Connect

A small connector service for querying a BigQuery dataset. 

### State
:construction: Project currently a work in progress

### Related Infrastructure
[bq-connect-infra](https://github.com/VinceDeslo/bq-connect-infra)

### Related Frontend
[bq-connect-frontend](https://github.com/VinceDeslo/bq-connect-frontend)

### Roadmap
- [x] Hexagonal architecture
- [x] HTTP API
- [x] Dockerfile and Docker Compose up and running
- [ ] gRPC API
- [ ] Redis for event streaming

### Purpose
This connector acts as a microservice to fetch and parse data from a BigQuery data lake.
The data lake contains a dataset of video game sales accross the years.
