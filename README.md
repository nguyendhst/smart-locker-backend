<h2 align="center"> Server-side code for the Smart Locker project </h2>

<h3 align="left"> Table of Contents </h3>

-   [API Documentation](#api-documentation)
-   [Installation](#installation)
-   [Database](#database)
-   [Development](#development)

### API Documentation

-   [x] `GET /api/hello` - Hello world
    - Resp:
        ```json
        {
            "message": "Hello world!"
        }
        ```
-   [ ] `GET /api/locker` - Get all lockers
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
            "id": "string"
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
            "status": "string"
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
            "status": "string"
        }
        ```

### Installation
-   Go 1.19 or higher is recommended

-   Clone the repository and run `go mod tidy` to install dependencies

-   Run `go run main.go` to start the server


### Database
- MySQL 8

- PlanetScale is used as the database platform for this project, access token is shared in the team chat (`config.json`)

    - [PlanetScale](https://planetscale.com/)

- TODOs:

    -   [ ] Define database schema

    -   [ ] Create database


### Development
-   Create a new branch for your feature

-   Create an issue for any feature you want to work on


<blockquote>
    <p>“Assumption is the mother of all fuck ups”</p>
    <p>— Neck Deep (Heartbreak Of The Century)</p>
</blockquote>