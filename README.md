# Enterprise Resource Planning (ERP) System

[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/erp-system)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Microservices](https://img.shields.io/badge/arch-microservices-brightgreen)](https://microservices.io)

Monorepo untuk sistem ERP berbasis microservices dengan Go, gRPC, dan PostgreSQL.

## 🌟 Fitur Utama
| Modul              | Deskripsi                                  | Status       |
|--------------------|--------------------------------------------|-------------|
| Auth Service       | Autentikasi & otorisasi terpusat           | ✅ Production Ready |
| HRD Module         | Manajemen karyawan, absensi, payroll       | 🚧 Development |
| Finance Module     | Invoice, laporan keuangan                  | ⏳ Planned    |
| Inventory Module   | Manajemen stok & gudang                    | ⏳ Planned    |

## 🏗️ Struktur Monorepo
```bash
erp-system/
├── 📁 auth-microservice/      # Layanan autentikasi (JWT, RBAC)
├── 📁 hrd-microservice/       # Modul HRD (employee, attendance, payroll)
│   ├── employee-service/     
│   ├── attendance-service/    
│   └── payroll-service/      
├── 📁 proto/                  # Protobuf definitions (shared)
├── 📁 deployments/            # Docker/K8s configs
├── 📁 api-gateway/            # Kong/Traefik configuration
├── 📁 docs/                   # Dokumentasi arsitektur
├── Makefile                  # Build semua services
└── README.md                 # File ini
