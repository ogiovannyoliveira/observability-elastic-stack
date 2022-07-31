package events

import (
	"github.com/giovannylucas/observability-elastic-stack/golang-server/database"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/models"
)

func Create(event *models.Event) {
	database.DB.Create(&event)
}