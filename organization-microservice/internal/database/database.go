package database

import (
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := config.BuildDSN(config.CFG)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Could not connect with the database")
	}
}

func Migrate() error {
	return DB.AutoMigrate(&models.Organization{}, &models.Department{}, &models.WorkUnit{}, &models.Position{}, &models.PositionHierarchy{})
}
