package events

import (
	"github.com/giovannylucas/observability-elastic-stack/golang-server/database"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/models"
)

func Create(event *models.Event) {
	database.DB.Raw(
		"INSERT INTO events (name, date) VALUES (?, ?) RETURNING *",
		event.Name,
		event.Date,
	).Scan(&event)
}