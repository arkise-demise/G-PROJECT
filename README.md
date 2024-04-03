# GO_PROJECT1

This is a simple Go project for handling user authentication and image uploads.

## Project Structure

The project consists of several packages and files:

- `db.go`: Defines the `Database` struct and methods for managing users and images.
- `auth_handler.go`: Contains handler functions for user authentication (login and registration).
- `image_handler.go`: Contains handler functions for uploading and retrieving images.
- `user_handler.go`: Contains handler functions for listing users.
- `models`: Package containing the data models (`User` and `Image`) used throughout the project.
- `main.go`: Main entry point of the application, where HTTP server is initialized and routes are defined.

## Usage

To run the application, execute the following command in your terminal:

```bash
go run main.go

This will start the server listening on localhost:8080.

Endpoints

POST /login: Endpoint for user login. Requires a JSON payload containing the username and password.

POST /register: Endpoint for user registration. Requires a JSON payload containing the username and password.

GET /users: Endpoint for listing all users.

POST /upload: Endpoint for uploading an image. Requires a valid JWT token in the Authorization header and a JSON payload containing image data.

GET /image: Endpoint for retrieving images. Requires a valid JWT token in the Authorization header.

Dependencies


This project relies on the following third-party packages:

github.com/dgrijalva/jwt-go: For JSON Web Token (JWT) generation and validation.

Ensure you have these dependencies installed before running the application.

```
