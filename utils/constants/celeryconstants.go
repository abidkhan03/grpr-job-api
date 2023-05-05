package constants

import (
	"os"
	"strconv"
)

var CeleryBrokerURL = getEnv("CELERY_BROKER_URL", "amqp://guest:guest@localhost:5672//")
var SnapshotsCount = getIntEnv("SNAP_SHOT_COUNT", 100)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if strValue, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(strValue); err == nil {
			return value
		}
	}
	return fallback
}
