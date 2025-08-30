# Go Movies CRUD API

A simple RESTful API built with Go for managing a movie collection. This project demonstrates basic CRUD (Create, Read, Update, Delete) operations using Go's built-in HTTP server and the Gorilla Mux router.

## 🎬 Features

- **Create** new movies with director information
- **Read** all movies or get a specific movie by ID
- **Update** existing movie details
- **Delete** movies from the collection
- **JSON** API responses
- **In-memory** data storage
- **RESTful** endpoint design

## 🛠️ Tech Stack

- **Go** (Golang) - Programming language
- **Gorilla Mux** - HTTP router and URL matcher
- **JSON** - Data exchange format
- **HTTP** - REST API protocol

## 📁 Project Structure

```
go-movies-crud/
├── main.go          # Main application file with all API handlers
├── go.mod          # Go module file with dependencies
├── go.sum          # Go dependencies checksum file
└── README.md       # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher installed on your machine
- Git (optional, for cloning)

### Installation

1. **Clone the repository** (or download the source code):
   ```bash
   git clone https://github.com/Amritanshu05/go-movies-crud.git
   cd go-movies-crud
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

4. **Server will start** on `http://localhost:8000`
   ```
   Starting server at port 8000
   ```

## 📡 API Endpoints

### Base URL: `http://localhost:8000`

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `GET` | `/movies` | Get all movies | None |
| `GET` | `/movies/{id}` | Get movie by ID | None |
| `POST` | `/movies` | Create new movie | Movie JSON |
| `PUT` | `/movies/{id}` | Update movie by ID | Movie JSON |
| `DELETE` | `/movies/{id}` | Delete movie by ID | None |

### 📝 Data Structure

#### Movie Object
```json
{
  "id": "string",
  "isbn": "string",
  "title": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

## 🧪 API Usage Examples

### 1. Get All Movies
```bash
curl -X GET http://localhost:8000/movies
```

**Response:**
```json
[
  {
    "id": "1",
    "isbn": "438227",
    "title": "Movie One",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  },
  {
    "id": "2",
    "isbn": "45455",
    "title": "Movie Two",
    "director": {
      "firstname": "Steve",
      "lastname": "Smith"
    }
  }
]
```

### 2. Get Movie by ID
```bash
curl -X GET http://localhost:8000/movies/1
```

### 3. Create New Movie
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "123456",
    "title": "New Movie",
    "director": {
      "firstname": "Jane",
      "lastname": "Director"
    }
  }'
```

### 4. Update Movie
```bash
curl -X PUT http://localhost:8000/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "999999",
    "title": "Updated Movie Title",
    "director": {
      "firstname": "Updated",
      "lastname": "Director"
    }
  }'
```

### 5. Delete Movie
```bash
curl -X DELETE http://localhost:8000/movies/1
```

## 🏗️ Code Structure

### Main Components

- **Movie Struct**: Represents a movie with ID, ISBN, title, and director
- **Director Struct**: Contains director's first and last name
- **Handler Functions**:
  - `getMovies()` - Retrieves all movies
  - `getMovie()` - Retrieves a specific movie by ID
  - `createMovie()` - Creates a new movie with auto-generated ID
  - `updateMovie()` - Updates an existing movie
  - `deleteMovie()` - Removes a movie from collection

### Key Features

- **Auto-generated IDs**: New movies get random IDs (0-9,999,999)
- **JSON Serialization**: Automatic conversion between Go structs and JSON
- **Route Parameters**: ID extraction from URL paths
- **HTTP Methods**: Proper REST verb usage
- **Content-Type Headers**: Correct JSON response headers

## 🧪 Testing with Postman or Thunder Client

1. Import the following collection or create requests manually:
   - GET `http://localhost:8000/movies`
   - GET `http://localhost:8000/movies/1`
   - POST `http://localhost:8000/movies` (with JSON body)
   - PUT `http://localhost:8000/movies/1` (with JSON body)
   - DELETE `http://localhost:8000/movies/1`

## 📚 Learning Resources

This project demonstrates:
- Go HTTP server creation
- RESTful API design
- JSON handling in Go
- Third-party package usage (Gorilla Mux)
- Struct definitions and tags
- Slice manipulation
- Error handling (basic)

## 🔮 Future Enhancements

- [ ] Database integration (PostgreSQL, MongoDB)
- [ ] Input validation and error handling
- [ ] Authentication and authorization
- [ ] Pagination for large datasets
- [ ] Unit tests
- [ ] Docker containerization
- [ ] API documentation with Swagger
- [ ] Logging middleware
- [ ] CORS support

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

**Amritanshu05**
- GitHub: [@Amritanshu05](https://github.com/Amritanshu05)

---

⭐ If you found this project helpful, please consider giving it a star!