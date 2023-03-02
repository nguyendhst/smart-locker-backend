<h2 align="center"> Server-side code for the Smart Locker project </h2>

<h3 align="left"> Table of Contents </h3>

-   [API Documentation](#api-documentation)
-   [Installation](#installation)
-   [Database](#database)
-   [Development](#development)

### Installation
-   Go 1.19 or higher is recommended

-   Clone the repository and run `go mod tidy` to install dependencies

-   Run `go run main.go` to start the server


### Database
-   MySQL 8

-   PlanetScale is used as the database platform for this project, access token is shared in the team chat (`config.json`)

    -   [PlanetScale](https://planetscale.com/)

-   Tools used:
    -   `sqlc` - Generate Go code from SQL queries
    -   `dbdiagram.io` - Generate database schema diagram
    -   `go-migrate` - Database migration tool

-   [Database Schema](https://dbdiagram.io/d/635783f4fa2755667d6744c7)

- TODOs:

    -   [x] Define database schema
        -   Cloud-based databases do not support foreign keys constraint, so we need to manually check the foreign key constraint in the application layer.

    -   [x] Create database

### Development
-   Create a new fork and work on your own changes

-   Create an issue for any feature you want to work on


### API Documentation

-   [x] `GET /api/hello` - Hello world
    - Resp:
        ```json
        {
            "message": "Hello world!"
        }
        ```
-   [ ] `POST /api/users/login` - Login
    - Body:
        ```json
        {
            "username": "string",
            "password": "string"
        }
        ```
-   [ ] `POST /api/users/register` - Register
    - Body:
        ```json
        {
            "username": "string",
            "password": "string"
        }
        ```
    - Resp:
        ```json
        {   
            "success": true,
            "token": "string"
        }
        ```
        ```json
        {   
            "success": false,
            "message": "string"
        }
        ```
    
-   [ ] `GET /api/locker` - Get all lockers
    - Body:
        ```json
        {
            "token": "string"
        }
        ```
    - Resp:
        ```json
        {
            "lockers": [
                {
                    "id": "string",
                    "name": "string",
                    "description": "string",
                    "location": "string",
                    "status": "string"
                }
            ]
        }
        ```
-   [ ] `GET /api/locker/:id` - Get a locker by id
    - Body:
        ```json
        {
            "id": "string",
            "token": "string"
        }
        ```
-   [ ] `POST /api/locker` - Create a new locker
    - Body:
        ```json
        {
            "id": "string",
            "name": "string",
            "description": "string",
            "location": "string",
            "status": "string",
            "token": "string"
        }
        ```
-   [ ] `PUT /api/locker/:id` - Update a locker by id
    - Body:
        ```json
        {
            "id": "string",
            "name": "string",
            "description": "string",
            "location": "string",
            "status": "string",
            "token": "string"
        }
        ```


<blockquote>
    <p>“Assumption is the mother of all fuck ups”</p>
    <p>— Neck Deep (Heartbreak Of The Century)</p>
</blockquote>