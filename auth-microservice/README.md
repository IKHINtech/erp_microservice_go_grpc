# Auth Microservice

[![Go Version](https://img.shields.io/github/go-mod/go-version/IKHINtech/erp_microservice_go_grpc)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Microservice autentikasi dan otorisasi untuk sistem ERP, dibangun dengan Golang, gRPC, dan PostgreSQL.

## 🚀 Fitur Utama

- Registrasi & Login User dengan JWT
- Manajemen Role-Based Access Control (RBAC)
- Validasi Token gRPC untuk layanan internal
- Integrasi dengan:
  - HRD Microservice (Employee Data)
  - Finance Microservice
  - Organization Microservice
- Dukungan Multi-Factor Authentication (MFA) _(coming soon)_

## 📦 Struktur Proyek

```bash
auth-microservice/
├── config/           # Konfigurasi environment
├── cmd/              # Entry point aplikasi
├── internal/         # Kode inti
│   ├── handlers/     # gRPC/HTTP handlers
│   ├── models/       # Database models
│   ├── repositories/ # Database operations
│   └── utils/        # JWT, password hashing
├── pb/               # Protobuf definitions
├── .env.example      # Template environment variables
├── go.mod            # Go module dependencies
└── README.md         # Dokumentasi ini
```
