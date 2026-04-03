# 🛒 Go Web Application - Product CRUD

![Go](https://img.shields.io/badge/Go-1.22-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue)
![Architecture](https://img.shields.io/badge/Architecture-MVC-green)
![Status](https://img.shields.io/badge/status-in%20progress-yellow)
![License](https://img.shields.io/badge/license-MIT-lightgrey)

## About

This project is a web application built with Go using only the standard library.

It implements a complete product CRUD (Create, Read, Update, Delete), applying the MVC architecture and persisting data in a PostgreSQL database.

---

## Features

- Create products
- List products
- Update product data
- Delete products
- Server-side rendering with HTML templates
- Database integration with PostgreSQL

---

## Architecture

The application follows the MVC pattern:

- **Model** → Database access and data logic  
- **View** → HTML templates  
- **Controller** → Request handling and business logic  
- **Routes** → Endpoint definitions  

---

## Tech Stack

- Go (Golang)
- Standard Library (`net/http`, `html/template`)
- PostgreSQL
- HTML

---

## Running the project

```bash
git clone <repository-url>
cd project-name
go run main.go
```

## License

This project is licensed under the MIT License.

## Notes

This project is part of a [Dev Journey](https://github.com/FernandaIshida/dev-journey/blob/main/README.md) focused on exploring backend development across multiple technologies and architectural approaches:

- Go (standard library) – building web applications without frameworks and applying MVC architecture  
- Go, Gin, Gorm – building REST APIs and structuring backend services  
- PostgreSQL – working with relational databases and data persistence  
- RabbitMQ – implementing asynchronous communication and event-driven architecture  
- Python (Django/DRF, FastAPI) – building APIs, microservices, and automations  
- Docker – containerization and environment management  

This project specifically demonstrates server-side rendering, full CRUD operations, and application structuring without frameworks, reinforcing core backend concepts and complementing my broader learning path in backend engineering.



