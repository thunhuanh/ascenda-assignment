# Hotel Finder
This is a simple web server written in `Golang` that retrieves, clean and store hotels data from multiple suppliers in the background and provides a simple API to search for hotels based on the destination and hotel ids.

## Requirements
- `Golang` version 1.20 or higher

## Installation
- Clone the repository
- Run `go run main.go` to start the server
- Use `curl` or `Postman` to test the API:
    - `curl -X GET http://34.87.111.208/api/v1/hotels?destination=5432`
    - `curl -X GET http://34.87.111.208/api/v1/hotels?destination=5432&hotels=iJhz`

## Testing
- Run `go test ./... -v` to run all tests

## Documentation
The documentation is available in the `docs` folder.

## Deployment Strategy
- Option 1:
  - We can deploy the whole application in a single container using `docker` to containerize the application and `Kubernetes` to manage the containers. In order to automate the deployment, we can use `Jenkins`, `github actions`,`circle-ci` or `gitlab-ci` to build the application and deploy it to Cloud provider like `GCP`. We can also run the test pipeline automatically before deploying the application to production.
- Option 2:
  - Since the `worker` and `api server` is not depend on each other, we can deploy them seperately, `worker` as a cronjob that run every 5 minutes and `api server` as a deployment.

- Cloud Provider: `GCP` or `AWS` or we can consider other option like `Heroku`

## Deployment
- The application is deployed on `GCP` using `Kubernetes` and `Docker`.
- The whole process is automated using `github actions`.
- App URL: `http://34.87.111.208/api/v1/`

## Further Improvements
- We can use `Redis` to cache the data and improve the performance of the API.
- We can store the config using `Vault` or `KMS` to make the application more secure.
- We can make this into a microservice architecture and use `gRPC` to communicate between the services.
- We can setup a gateway to handle the traffic and route the request to the correct service.
- We can run the repo through `SonarQube` to check the code quality and fix the issues.
- We can use `Prometheus` to monitor the application and `Grafana` to visualize the metrics.
- We can use `Helm` to manage the deployment of the application.
- etc.