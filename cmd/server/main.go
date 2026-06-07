package main

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/config"
	"github.com/Abrahamthefirst/finecore-practice/internal/db"
	"github.com/Abrahamthefirst/finecore-practice/pkg/logger"
)

func main() {

	cfg := config.Load()
	db := db.NewPgDB(cfg.DATABASE_URL)

	app := NewApp(db, cfg, logger.New(false))

	app.Bootstrap()
}
