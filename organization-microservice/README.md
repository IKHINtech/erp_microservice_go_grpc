# Organization Microservice

[![Go Version](https://img.shields.io/github/go-mod/go-version/IKHINtech/erp_microservice_go_grpc)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Microservice Mangajemen Organisasi (Position, Department, Work Unit) dibangun dengan Golang, gRPC, dan PostgreSQL.

## 🚀 Fitur Utama

- Manajemen Organisasi
- Manajemen Department
- Manajemen Work Unit
- Integrasi dengan:
  - HRD Microservice (Employee Data)
  - Finance Microservice

## 📦 Struktur Proyek

```bash
auth-microservice/
├── client/           # Klien gRPC
├── config/           # Konfigurasi environment
├── cmd/              # Entry point aplikasi
├── internal/         # Kode inti
│   ├── database/     # Database operations
│   ├── dto/          # Data Transfer Objects
│   ├── handlers/     # gRPC/HTTP handlers
│   ├── models/       # Database models
│   ├── repositories/ # Database operations
│   └── utils/        # JWT, password hashing
├── pb/               # Protobuf definitions
├── .env.example      # Template environment variables
├── go.mod            # Go module dependencies
└── README.md         # Dokumentasi ini
```
