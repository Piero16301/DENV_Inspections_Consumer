package configs

import (
	"DENV_Inspections_Consumer/models"
	"os"
)

func GetKafkaProperties() (models.KafkaProperties, error) {
	kfProperties := models.KafkaProperties{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}

	return kfProperties, nil
}
