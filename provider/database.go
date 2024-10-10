package provider

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DatabaseKey = "database"

func WithDatabase(ctx context.Context) context.Context {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

    return context.WithValue(ctx, DatabaseKey, db)
}
