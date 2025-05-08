# Pinned Front

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-front)](https://goreportcard.com/report/github.com/patricksferraz/pinned-front)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-front?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-front)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, high-performance web application built with Go and Fiber framework, designed to manage guests, employees, places, and menus in a restaurant or hospitality environment.

## 🚀 Features

- **High Performance**: Built with [Fiber](https://github.com/gofiber/fiber), a fast and efficient web framework
- **Modern Architecture**: Clean code structure following domain-driven design principles
- **Security**: Built-in CSRF protection and input validation
- **Scalable**: Containerized with Docker and Kubernetes support
- **Developer Friendly**: Hot-reload support with Air
- **Environment Configuration**: Flexible configuration management with environment variables

## 🛠️ Tech Stack

- **Backend**: Go 1.18+
- **Web Framework**: Fiber v2
- **Template Engine**: Fiber Template
- **Containerization**: Docker & Docker Compose
- **Orchestration**: Kubernetes
- **Development**: Air for hot-reload
- **Environment**: Go-ENV for configuration management

## 📋 Prerequisites

- Go 1.18 or higher
- Docker and Docker Compose (for containerized deployment)
- Make (for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/pinned-front.git
cd pinned-front
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Run the application:

For development:
```bash
make dev
```

For production:
```bash
make build
make run
```

Using Docker:
```bash
docker-compose up
```

## 🏗️ Project Structure

```
.
├── app/            # Application layer (handlers, routes)
├── config/         # Configuration management
├── domain/         # Domain layer (services, models)
├── utils/          # Utility functions
├── k8s/            # Kubernetes configurations
├── .docker/        # Docker-related files
└── .github/        # GitHub workflows and templates
```

## 🔧 Configuration

The application can be configured using environment variables. See `.env.example` for available options.

## 🧪 Testing

```bash
make test
```

## 📦 Deployment

### Docker
```bash
docker-compose up -d
```

### Kubernetes
```bash
kubectl apply -f k8s/
```

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/patricksferraz/pinned-front/issues).

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

Patrick Sferraz

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber) - Fast and efficient web framework
- [Air](https://github.com/cosmtrek/air) - Live reload for Go applications
- All other open-source projects that made this possible
