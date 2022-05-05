# BQ Connect

A small connector service for querying a BigQuery dataset of plants 

### Infrastructure
[Infrastructure of the project](https://github.com/VinceDeslo/bq-connect-infra)

### Roadmap
- [ ] Hexagonal architecture
- [ ] gRPC communications for ports and adapters
- [ ] Redis as a caching layer
- [ ] Docker config and image deployed to DockerHub

### Purpose
This connector acts as a microservice to fetch and parse data from a BigQuery data lake.
The data lake contains a dataset of plants and their corresponding states/provinces in North America.
It should do the translation of state/province acronyms into proper names.



