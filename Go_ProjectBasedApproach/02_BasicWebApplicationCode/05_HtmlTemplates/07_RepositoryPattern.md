## Repository Pattern in Golang

The **Repository Pattern** is a design pattern commonly used in software development to abstract the data layer and provide a more domain-specific interface to work with data sources such as databases, APIs, or even in-memory storage. In **Go (Golang)**, the repository pattern is widely used to organize code, especially in web applications with layered architectures like **Clean Architecture**, **DDD (Domain-Driven Design)**, or **Hexagonal Architecture**.

Let’s dive into the concept, followed by examples and use cases.

---

### **What is the Repository Pattern?**

The Repository Pattern acts as a bridge between the domain/business logic and the data access layer. It abstracts database queries, ensuring that:

1. The business logic doesn’t directly interact with the database.
2. Code is more testable and maintainable by injecting dependencies.
3. Switching data sources (e.g., from MySQL to MongoDB) requires minimal changes.

---

### **Core Components**

1. **Domain Entities**: Represent the core business objects.
2. **Repository Interface**: Defines methods to interact with the data layer.
3. **Repository Implementation**: Implements the interface with specific data access logic (e.g., SQL queries).
4. **Service Layer**: Calls the repository to perform operations.

---

### **Example: Building a Repository Pattern in Golang**

#### Use Case: A Bookstore application with a `Book` entity.

#### 1. **Define the Domain Entity**

```go
package models

type Book struct {
    ID          int64  `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    PublishedAt string `json:"published_at"`
}
```

---

#### 2. **Create the Repository Interface**

The repository interface defines the operations we can perform on `Book` data.

```go
package repositories

import "github.com/yourapp/models"

type BookRepository interface {
    GetByID(id int64) (*models.Book, error)
    GetAll() ([]*models.Book, error)
    Create(book *models.Book) error
    Update(book *models.Book) error
    Delete(id int64) error
}
```

---

#### 3. **Implement the Repository**

The implementation contains the actual logic to interact with a data source. Here’s an example with a PostgreSQL database using the `pgx` library:

```go
package repositories

import (
    "database/sql"
    "github.com/yourapp/models"
    "fmt"
)

type PostgresBookRepository struct {
    DB *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
    return &PostgresBookRepository{DB: db}
}

func (r *PostgresBookRepository) GetByID(id int64) (*models.Book, error) {
    query := "SELECT id, title, author, published_at FROM books WHERE id = $1"
    row := r.DB.QueryRow(query, id)

    book := &models.Book{}
    err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedAt)
    if err != nil {
        return nil, err
    }

    return book, nil
}

func (r *PostgresBookRepository) GetAll() ([]*models.Book, error) {
    query := "SELECT id, title, author, published_at FROM books"
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []*models.Book
    for rows.Next() {
        book := &models.Book{}
        err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedAt)
        if err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}

func (r *PostgresBookRepository) Create(book *models.Book) error {
    query := "INSERT INTO books (title, author, published_at) VALUES ($1, $2, $3)"
    _, err := r.DB.Exec(query, book.Title, book.Author, book.PublishedAt)
    return err
}

func (r *PostgresBookRepository) Update(book *models.Book) error {
    query := "UPDATE books SET title = $1, author = $2, published_at = $3 WHERE id = $4"
    _, err := r.DB.Exec(query, book.Title, book.Author, book.PublishedAt, book.ID)
    return err
}

func (r *PostgresBookRepository) Delete(id int64) error {
    query := "DELETE FROM books WHERE id = $1"
    _, err := r.DB.Exec(query, id)
    return err
}
```

---

#### 4. **Service Layer (Business Logic)**

The service layer interacts with the repository and performs the business logic.

```go
package services

import (
    "github.com/yourapp/models"
    "github.com/yourapp/repositories"
)

type BookService struct {
    Repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) *BookService {
    return &BookService{Repo: repo}
}

func (s *BookService) GetBookByID(id int64) (*models.Book, error) {
    return s.Repo.GetByID(id)
}

func (s *BookService) ListBooks() ([]*models.Book, error) {
    return s.Repo.GetAll()
}

func (s *BookService) AddBook(book *models.Book) error {
    return s.Repo.Create(book)
}

func (s *BookService) UpdateBook(book *models.Book) error {
    return s.Repo.Update(book)
}

func (s *BookService) RemoveBook(id int64) error {
    return s.Repo.Delete(id)
}
```

---

#### 5. **Usage in the Main Application**

```go
package main

import (
    "database/sql"
    "fmt"
    "github.com/yourapp/models"
    "github.com/yourapp/repositories"
    "github.com/yourapp/services"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
    // Database connection
    db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/bookstore?sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Initialize repository and service
    repo := repositories.NewPostgresBookRepository(db)
    service := services.NewBookService(repo)

    // Add a new book
    newBook := &models.Book{
        Title:       "Golang Basics",
        Author:      "John Doe",
        PublishedAt: "2023-01-01",
    }
    if err := service.AddBook(newBook); err != nil {
        fmt.Println("Error adding book:", err)
    } else {
        fmt.Println("Book added successfully")
    }

    // List all books
    books, err := service.ListBooks()
    if err != nil {
        fmt.Println("Error listing books:", err)
    } else {
        for _, book := range books {
            fmt.Printf("Book: %+v\n", book)
        }
    }
}
```

---

### **Advantages of Using Repository Pattern**

1. **Abstraction**: Hides the complexities of the data source.
2. **Testability**: Easier to mock the repository interface for unit testing.
3. **Maintainability**: Changes to the data source or schema affect only the repository implementation.
4. **Reusability**: Promotes code reuse by separating concerns.

---

### **Use Cases**

1. **Microservices**: Each service might use the repository pattern to handle its data independently.
2. **Web Applications**: Common in REST APIs for managing CRUD operations.
3. **Data-Driven Applications**: Abstracting complex database queries, joins, and transactions.
4. **Switching Data Sources**: If you need to migrate from SQL to NoSQL, you only update the repository implementation.

---

### **Testing the Repository**

Mocking the repository interface allows for testing the service layer without connecting to the database.

```go
package tests

import (
    "testing"
    "github.com/yourapp/models"
    "github.com/yourapp/services"
    "github.com/stretchr/testify/mock"
)

// Mock repository
type MockBookRepository struct {
    mock.Mock
}

func (m *MockBookRepository) GetByID(id int64) (*models.Book, error) {
    args := m.Called(id)
    return args.Get(0).(*models.Book), args.Error(1)
}

// ... Implement other methods similarly

func TestGetBookByID(t *testing.T) {
    repo := new(MockBookRepository)
    service := services.NewBookService(repo)

    // Mock behavior
    repo.On("GetByID", int64(1)).Return(&models.Book{
        ID:    1,
        Title: "Mock Book",
    }, nil)

    // Test
    book, err := service.GetBookByID(1)
    if err != nil || book.Title != "Mock Book" {
        t.Error("Expected Mock Book")
    }
}
```

---

This setup demonstrates a robust implementation of the Repository Pattern in Go. It ensures clear separation of concerns and promotes scalable, testable, and maintainable code.
