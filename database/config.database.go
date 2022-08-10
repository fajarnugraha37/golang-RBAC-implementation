package database

import (
	"log"
	"os"
	"time"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/database/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func NewDB() *gorm.DB {
	return util.Singleton(db, func() *gorm.DB {
		return initialize()
	})
}

func initialize() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	gormDB, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database, case: " + err.Error())
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic("failed to get sqlDB, case: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if res := os.Getenv("DB_MIGRATION"); res == "true" {
		err = gormDB.AutoMigrate(
			&entity.Gift{},
			&entity.Permission{},
			&entity.Rating{},
			&entity.Redeem{},
			&entity.Role{},
			&entity.User{},
		)
		if err != nil {
			panic("failed to migrate database, cause:  " + err.Error())
		}
	}

	if res := os.Getenv("DB_SEED"); res == "true" {
		for _, v := range SeedAll() {
			err := v.Run(gormDB)
			if err != nil {
				panic("failed to seed database, cause: " + err.Error())
			}
		}
	}

	return gormDB
}
