<h2 align="center"> Server-side code for the Smart Locker project </h2>

<h3 align="left"> Table of Contents </h3>

-   [API Documentation](#api-documentation)
-   [Installation](#installation)
-   [Configuration](#configuration)
-   [Database](#database)
-   [CI/CD](#cicd)
-   [Development](#development)

### Installation

-   Go 1.19 or higher is recommended

-   Clone the repository and run `go mod tidy` to install dependencies

-   Run `go run main.go` to start the server

### Configuration

-   `config.json` is used to store configuration for the project
-   `config.json` is not included in the repository, please create your own `config.json` file in the root directory of the project.
-   ```json
    {
        "port": "8080",
        "dsn": "xxxxxxxxxxxxxxx:pscale_pw_xxxxxxxxxxxxxxxxxxx@tcp(ap-southeast.connect.psdb.cloud)/smart-locker?tls=true",
        "planetscale_db": "smart-locker",
        "adafruit_username": "xxxxxxxxxxx",
        "adafruit_key": "aio_xxxxxxxxxxxxxxxxxxxxxxxxxx"
    }
    ```

### Database

-   MySQL 8

-   PlanetScale is used as the database platform for this project, access token is shared in the team chat (`config.json`)

    -   [PlanetScale](https://planetscale.com/)

-   Tools used:

    -   `sqlc` - Generate Go code from SQL queries
    -   `dbdiagram.io` - Generate database schema diagram
    -   `go-migrate` - Database migration tool

-   [Database Schema](https://dbdiagram.io/d/635783f4fa2755667d6744c7)

-   TODOs:

    -   [x] Define database schema

        -   Cloud-based databases do not support foreign keys constraint, so we need to manually check the foreign key constraint in the application layer.

    -   [x] Create database

### CI/CD

-   Under consideration

### Development

-   Create a new fork and work on your own changes

-   Create an issue for any feature you want to work on

### API Documentation

-   [Swagger](https://app.swaggerhub.com/apis-docs/nguyendhst/smart-locker/1.0.0)

<blockquote>
    <p>“Assumption is the mother of all fuck ups”</p>
    <p>— Neck Deep (Heartbreak Of The Century)</p>
</blockquote>
````
