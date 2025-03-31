## Project: Go-Movies-Crud

### Overview
`go-movies-crud` is a RESTful API built with Go that allows users to manage a collection of movies. It supports full CRUD operations (Create, Read, Update, Delete) and uses the Gorilla Mux router for handling HTTP requests. The project stores movie data in memory (a slice) and includes fields for movie ID, ISBN, title, and director details.

### Features
- **GET /movies**: Retrieve all movies.
- **GET /movies/{id}**: Retrieve a specific movie by ID.
- **POST /movies**: Create a new movie with a randomly generated ID.
- **PUT /movies/{id}**: Update an existing movie.
- **DELETE /movies/{id}**: Delete a movie by ID.

### Technologies Used
- **Go (Golang)**: Backend language.
- **Gorilla Mux**: HTTP router for handling routes.
- **Postman**: For testing API endpoints.
