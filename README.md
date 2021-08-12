# Any Service Template
A template for creating Golang based microservice from scratch.

## Folder Hierarchy
The hierachy for this project is based on the below assumptions.

### cmd/
This is the entrypoint for a service. This package defines the type or kind of service that is being developed. They can be a job, daemon, api based service or any other.

### internal/
This package contains the code that the user intends to hide or disable others from importing them in other projects.

### pkg/
This package contains sub packages that can be used across different projects, packages.

## Components
### logger
### config
### cache
### server
### database
### broker
