# Documentation

This is a simple web server written in `Golang` that retrieves, clean and store hotels data from multiple suppliers in the background and provides a simple API to search for hotels based on the destination and hotel ids.

## Components
There are 2 main components in the application:
- `worker` that retrieves, clean and store hotels data from multiple suppliers in the background. There is also a cronjob that runs every 5 minutes to fetch the data from suppliers and send it to the `worker` to be processed. The `worker` then merges the data from all suppliers, clean it and store it in the database.
  
- `api server` that provides a simple API to search for hotels based on the destination and hotel ids

## Design Decisions
- The `worker` is a separate component that runs in the background and retrieves, clean and store hotels data from multiple suppliers. This is to decouple the data retrieval and storage from the API server. This way, the API server can be scaled independently from the `worker` and the `worker` can be scaled independently from the API server. Another reason for this is to increase the performance of the API since we dont have to fetch the data from suppliers every time we want to search for hotels. And also to reduce the load on the suppliers.  (for this project i didn't separate the `worker` and the `api server`)
  
- The `api server` is built using `go-chi` router. This is to provide a simple and fast router that is easy to use and extend. And leverage the `mongodb` to store the data. I chose `mongodb` because it is a document database and it is easy to store and query the data. I also chose `mongodb` because it is easy to scale horizontally and because of its flexibility. And it also supports `geospatial queries` which is needed to search for hotels based on the latitude and longitude if we ever want to add this feature later.

- For optimizing the application, i leverage go routines to run the `worker` and the `api server` concurrently. And also to run the `worker` concurrently for each supplier. This is to increase the performance of the application. And also Decouple the data retrieval and storage from the API server. This way, the API server has no dependency on the `worker` and the `worker` has no dependency on the API server, it should make the API faster and more reliable.

- I also used `docker` to containerize the application and `Kubernetes` to manage the containers. This is to make it easy to run the application in any environment. And also to make it easy to scale the application horizontally.
