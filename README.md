# golangAdvace

# Modular Movie Database in Go

A robust, modular Go application designed to master professional architectural patterns, including interfaces, dependency injection, file-based persistence, thread-safe concurrency control, and clean separation of concerns.

---

## 🚀 Project Overview & Evolution

This project has evolved from a simple in-memory data store into a modular application capable of polymorphic storage swapping. It demonstrates how real-world Go systems decouple business logic from underlying persistence mechanisms.

### Core Architectural Features:

- **Interface-Driven Design:** The application layer interacts with storage via Go interfaces (`DBInterface`), ensuring complete decoupling.
- **Polymorphic Storage Swapping:** Easily switch between different database backends (`db_mock` RAM vs. `filedb` JSON persistence) without modifying application code or `main.go` logic.
- **Thread-Safe File I/O:** Uses `sync.RWMutex` to protect local file assets from data corruption during concurrent read and write operations.
- **Efficient JSON Streaming:** Implements `json.NewEncoder` and `json.NewDecoder` for optimized memory performance when persisting data to disk.
- **Modern Best Practices:** Avoids deprecated functions in favor of robust string manipulation and clean error handling.

---

## 📂 Project Architecture & Components

```text
golangAdvace/
├── db_mock/        # In-memory database implementation (RAM)
├── fileDB/         # File-backed JSON database implementation
├── model/          # Shared domain models (e.g., Movie structs)
└── main.go         # Application entry point and dependency injection container

```

### 1. The Domain Model (`model`)

Defines the core data structure used across all database implementations.

```go
type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}

```

### 2. The Abstract Contract (`DBInterface`)

Standardizes behavior across storage engines so components remain interchangeable.

```go
type DBInterface interface {
	ShowMovies() error
	ShowMovie(movieID int) error
	AddMovie(name string, description string, year int) error
	UpdateMovie(movieID int, name, description string, year int) error
	DeleteMovie(movieID int) error
}

```

### 3. File-Based Persistence & Concurrency (`fileDB`)

Implements file I/O safely using a Reader/Writer Mutual Exclusion Lock (`sync.RWMutex`) and helper functions to maintain a clean JSON array structure on disk.

- **`RLock()` / `RUnlock()`:** Allows multiple concurrent readers (e.g., fetching movie lists) without blocking each other.
- **`Lock()` / `Unlock()`:** Ensures exclusive access when mutating data (adding, updating, or deleting movies).

---

## 💡 Core Architectural Concepts (Cheat Sheet)

### 1. Abstraction

- **What it is:** Hiding complex low-level implementation details behind a clean, simple interface so the rest of your program only has to care _what_ something does, not _how_ it does it.
- **In this project:** The application core calls `app.db.AddMovie(...)`. It doesn't need to know whether that movie is being saved to a volatile RAM slice or streamed into a physical JSON file on a hard drive.

### 2. Polymorphism

- **What it is:** The ability of different types to implement the exact same interface, allowing them to be used interchangeably.
- **In this project:** Both `dbmock.DB` and `fileDB.FileDB` satisfy the `DBInterface` contract. Because they share the exact same method signatures, your application container can swap storage engines seamlessly in `main.go`.

### 3. Embedding (Composition)

- **What it is:** In Go, embedding means including a field inside a struct without giving it an explicit field name (anonymous field). This automatically forwards or "promotes" the inner type's methods up to the outer struct.
- **Example in Go:**

```go
type App struct {
    dbdbmock.DBInterface // Interface embedding or field composition
}

```

---

## 💡 Key Technical Learnings

1. **`sync.Mutex` vs `sync.RWMutex`:** Use exclusive `Mutex` for hard-locked writes, and `RWMutex` when you want multiple simultaneous readers (`RLock`) to improve performance on read-heavy tasks like viewing lists.
2. **`json.Marshal` vs `json.NewEncoder`:** Use `Marshal` when you need an in-memory byte slice (`[]byte`), but use `NewEncoder` for direct streaming straight to files or network sockets for better memory efficiency.
3. **Dependency Injection:** Injecting concrete database implementations into the `App` container at runtime rather than hardcoding dependencies.
