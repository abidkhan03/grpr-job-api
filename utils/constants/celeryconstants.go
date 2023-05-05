package constants

import (
	"fmt"
	"os"
	"strconv"
)

func getConfig(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = defaultValue
	}
	return value
}

func CeleryConstants() {
	celeryBrokerURL := getConfig("CELERY_BROKER_URL", "amqp://guest:guest@localhost:5672//")
	snapshotsCount, err := strconv.Atoi(getConfig("SNAP_SHOT_COUNT", "100"))
	if err != nil {
		fmt.Println("Error parsing SNAP_SHOT_COUNT:", err)
		return
	}

	fmt.Println("CELERY_BROKER_URL:", celeryBrokerURL)
	fmt.Println("SNAP_SHOT_COUNT:", snapshotsCount)
}
