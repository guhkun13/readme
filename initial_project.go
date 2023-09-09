package init_project

import (
	"event-service/config"
	"event-service/pkg/database"
	"event-service/pkg/logger"
	"log"
	"path/filepath"

	"go.uber.org/zap"
)

func NewProject() {
	l, err := logger.NewZapFileLogger(filepath.Join("./log", "error.log"))
	zap.ReplaceGlobals(zap.Must(l, err))

	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("config error", err)
	}
	database.ConnectMYSQL("mysql", config)
}
