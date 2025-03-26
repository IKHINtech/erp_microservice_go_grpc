package main

import (
	"log"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Koneksi PostgreSQL

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	// Jalankan migrations
	authRepo := repositories.NewAuthRepository(db)
	if err := authRepo.Migrate(); err != nil {
		log.Fatal("Gagal migrate database:", err)
	}
	log.Printf("Auth Service berjalan di port %s", cfg.AuthPort)
}

func buildDSN(cfg config.Config) string {
	return "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=" + cfg.DBSSLMode
}
