package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Conn *gorm.DB
	DB   *sql.DB
}

func New(ctx context.Context) (*Gorm, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres gorm, err: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	db, err := conn.WithContext(ctx).DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres gorm, err: %w", err)
	}

	db.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	db.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres gorm, err: %w", err)
	}

	return &Gorm{
		Conn: conn,
		DB:   db,
	}, err
}
