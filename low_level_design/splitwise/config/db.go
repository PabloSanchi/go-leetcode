package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	models "splitwise/domain/db"
)

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("splitwise.db"), &gorm.Config{})
	if err != nil {
		slog.Error("error connecting to database", slog.Any("err", err))
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Group{},
	); err != nil {
		slog.Error("error migrating user table", slog.Any("err", err))
		return nil, err
	}

	return db, nil
}
