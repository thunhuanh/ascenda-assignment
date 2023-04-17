# Hotel Finder
This is a simple web server written in `Golang` that retrieves, clean and store hotels data from multiple suppliers in the background and provides a simple API to search for hotels based on the destination and hotel ids.

## Requirements
- `Golang` version 1.20 or higher

## Installation
- Clone the repository
- Run `go run main.go` to start the server
- Use `curl` or `Postman` to test the API:
    - `curl -X GET http://localhost:8080/hotels?destination=5432`
    - `curl -X GET http://localhost:8080/hotels?destination=5432&hotels=iJhz`

## Testing
- Run `go test ./... -v` to run all tests

## Documentation
The documentation is available in the `docs` folder.

## Deployment Strategy
- Option 1:
  - We can deploy the whole application in a single container using `docker` to containerize the application and `Kubernetes` to manage the containers. In order to automate the deployment, we can use `Jenkins` or `github actions` or `gitlab-ci` to build the application and deploy it to `GKE`. We can also run the test pipeline automatically before deploying the application to production.
- Option 2:
  - Since the `worker` and `api server` is not depend on each other, we can deploy them seperately, `worker` as a cronjob that run every 5 minutes and `api server` as a deployment.

- Cloud Provider: `GCP` or `AWS` or we can consider other option like `Heroku`