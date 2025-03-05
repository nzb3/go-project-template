# Go Project Template Documentation

This repository provides a clean and structured template for quickly bootstrapping new Go projects. It includes best practices, organized directory structure, and essential tooling to help you start coding immediately.

---

## 📌 Features

- **Standardized Project Structure**: Clear and maintainable layout following Go community conventions.
- **Dependency Management**: Integrated Go modules for easy dependency handling.
- **Testing Setup**: Pre-configured test structure to encourage test-driven development (TDD).
- **Configuration Management**: Simple configuration setup with environment variables support.
- **Docker Integration**: Dockerfile included for containerization and easy deployment.

---

## 📂 Project Structure

```
.
├── cmd/                # Application entry points
│   └── app/            # Main application
│       └── main.go     # Main application entry point
├── Dockerfile          # Docker configuration
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── internal/           # Private application code
│   ├── app/            # Application core components
│   │   └── app.go      # Application initialization
│   └── config/         # Configuration handling
│       └── config.go   # Configuration implementation
├── LICENSE             # License file
├── Makefile            # Build automation
├── pkg/                # Public libraries that can be used by external applications
├── README.md           # Project documentation
└── tools/              # Supporting tools for the project
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.22+](https://golang.org/dl/)
- [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) tool
- [Docker](https://www.docker.com/get-started) (optional)

### Installation

Create a new project using `gonew`:

```bash
# Install gonew if you don't have it
go install golang.org/x/tools/cmd/gonew@latest

# Create a new project from the template
gonew github.com/nzb3/go-project-template github.com/yourusername/your-project-name
cd your-project-name
```

### Running the Application

Run the application locally:

```bash
go run cmd/app/main.go
```

### Testing

Run unit tests:

```bash
go test ./...
```

---

## ⚙️ Configuration

The project uses environment variables for configuration. You can set these variables directly or use a `.env` file in the project root.

Example `.env` file:

```env
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙌 Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have improvements or suggestions.

---

## 📬 Contact

Author: [Nikolay Bukhalov](https://github.com/nzb3)

Feel free to reach out via GitHub issues or discussions if you have questions or need assistance.
