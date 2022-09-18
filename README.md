# BQ Connect

A small connector service for querying a BigQuery dataset. 

### Related Infrastructure
[bq-connect-infra](https://github.com/VinceDeslo/bq-connect-infra)

### Related Frontend
[bq-connect-frontend](https://github.com/VinceDeslo/bq-connect-frontend)

### Roadmap
- [x] Hexagonal architecture
- [x] HTTP API
- [x] Dockerfile and Docker Compose up and running
- [ ] Github actions to push container to GCP Artifact Registry
- [ ] gRPC API
- [ ] Event streaming to a GCP Pub/Sub topic

### Purpose
This connector acts as a microservice to fetch and parse data from a BigQuery data lake.
The data lake contains a dataset of video game sales accross the years.
