package server

import (
	"encoding/json"
	"fmt"
	"os"
	"server/config"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB(conf *config.Config) (*gorm.DB, error) {
	db, err := connectDB(conf)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Album{},
		&models.Post{},
		// &models.User{},
	)
	if err != nil {
		return nil, err
	}

	err = seedDB(db, conf)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func seedDB(db *gorm.DB, conf *config.Config) error {
	var count int64
	db.Model(&models.Album{}).Count(&count)
	if count == 0 {
		err := load[models.Album](db, "albums", conf)
		if err != nil {
			return err
		}
	}

	db.Model(&models.Post{}).Count(&count)
	if count == 0 {
		err := load[models.Post](db, "posts", conf)
		if err != nil {
			return err
		}
	}

	return nil
}

func load[T any](db *gorm.DB, tableName string, conf *config.Config) error {
	raw, err := os.ReadFile(fmt.Sprintf("%sdata/%s.json", conf.BaseDir, tableName))
	if err != nil {
		return err
	}

	var models []*T
	err = json.Unmarshal(raw, &models)
	if err != nil {
		return err
	}

	db.Create(models)
	// TODO handle SQL errors too?
	db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', (SELECT MAX(id) from %s))", tableName, tableName))
	return nil
}

func connectDB(conf *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN: conf.DatabaseUrl,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
