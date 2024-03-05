# url-shortener Service

This project is a URL shortening service implemented in Go. It allows for the generation of short URLs that redirect to longer URLs, making it easier to share and manage links.

## Technologies and Frameworks

The project uses the following technologies and frameworks:

- Go programming language
- Docker for containerization
- Redis for caching
- MySQL for database storage
- Echo framework for building HTTP APIs

## Project Structure

The project is organized into several directories and files, each with a specific purpose:

- `Dockerfile`: This file contains the configuration for building the Docker image of the application.

### main.go
- `main.go`: This is the main entry point of the application. It calls the `cmd.Execute()` function to start the application.

- `cmd/`: This directory contains the main entry points of the application.

    ### cmd/root.go
  - `root.go`: This file initializes the application.

    ### cmd/serve.go
  - `serve.go`: This file starts the HTTP server.

- `config/`: This directory contains configuration files and related code.

    ### config/mysql.go
  - `mysql.go`: This file contains the code for connecting to and interacting with the MySQL database.

    ### config/redis.go
  - `redis.go`: This file contains the code for connecting to and interacting with the Redis cache.

### handlers/handlers.go

This file defines the HTTP request handlers for the application. It defines a `Handlers` struct that has a `Service` field. The `Service` field is a pointer to a `Service` struct defined in the `Service` package. This struct is used to handle the business logic of the application.

The `Handlers` struct has two methods: `ShortenUrlHandler` and `ResolvUrlHandler`.

- `ShortenUrlHandler`: This method returns a function that handles HTTP requests to shorten a URL. It takes a long URL from the request body, generates a short URL, and returns it in the response. If there is an error during this process, it returns the error.

- `ResolvUrlHandler`: This method returns a function that handles HTTP requests to resolve a short URL. It takes a short URL from the request path, retrieves the corresponding long URL, and redirects to the long URL. If there is an error during this process, it returns the error.

The `NewHandlers` function is used to create a new instance of the `Handlers` struct. It takes a pointer to a `Service` struct as input and returns a pointer to a `Handlers` struct.

  ### service/service.go
- `service/`: This directory contains the business logic code.
  - `service.go`:
  - This file contains the core business logic of the URL shortening service. It defines a `service` struct that has a `Repository` field. The `Repository` field is a pointer to a `Repository` struct defined in the `repository` package. This struct is used to interact with the database and cache.
   The `service` struct has two methods: `ShortenUrl` and `ResolveUrl`.

  - `ShortenUrl`: This method takes a context and a long URL as input. It generates a short URL by creating an MD5 hash of the long URL and encoding it as a hexadecimal string. It then inserts the long URL and its corresponding short URL into the database. If the insertion is successful, it returns the short URL. If there is an error during the insertion, it returns the error.

  - `ResolveUrl`: This method takes a context and a short URL as input. It first tries to retrieve the corresponding long URL from the cache. If the long URL is not in the cache, it retrieves it from the database and then inserts it into the cache. If there is an error during this process, it returns the error. If the process is successful, it returns the long URL.

  - The `NewLogics` function is used to create a new instance of the `Logics` struct. It takes a pointer to a `Repository` struct as input and returns a pointer to a `Logics` struct.

### models/models.go
- `models/`: This directory contains data models and related code.
  - `model.go`: This file defines and interacts with data models.

- `repository/`: This directory contains code for data access and storage.
### repository/cache/cache.go
  - `cache/`: This subdirectory contains code for interacting with the cache.

  ### repository/database/database.go
  - `database/`: This subdirectory contains code for interacting with the database.

  ### repository/repository.go
  - `repository.go`: This file manages data access and storage.


### routes/api.go
- This file is responsible for defining the API routes for the application. It contains a function `InitializeGroup` which takes an instance of `echo.Echo` and a `Handlers` struct as parameters. Inside this function, a new group of routes is created under the path `/api/v1`.

Two routes are defined within this group:

- `GET /:shortUrl`: This route is handled by the `ResolvUrlHandler` method of the `Handlers` struct. It resolves a short URL to its original URL.
- `POST /shorten`: This route is handled by the `ShortenUrlHandler` method of the `Handlers` struct. It creates a short URL from a given long URL.

## Running the Application

To run the application, you need to have Docker installed on your machine. Once Docker is installed, you can build and run the application using the following commands:

bash docker build -t urlshortener . docker run -p 8080:8080 urlshortener

The application will be accessible at `http://localhost:8080`.

## API Endpoints

The application provides the following API endpoints:

- `GET /api/v1/:shortUrl`: Resolves a short URL to its original URL.
- `POST /api/v1/shorten`: Creates a short URL from a given long URL.

## Contributing

Contributions to this project are welcome. Please feel free to open an issue or submit a pull request.
