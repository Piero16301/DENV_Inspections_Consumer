package main

import (
	"DENV_Inspections_Consumer/configs"
	"DENV_Inspections_Consumer/controllers"
	"DENV_Inspections_Consumer/models"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	message, err := configs.Consumer.ConsumePartition("register-home-inspection", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	for {
		var homeInspection models.HomeInspection
		messages := <-message.Messages()
		err = json.Unmarshal(messages.Value, &homeInspection)
		if err != nil {
			panic(err)
		}

		// Enviar Inspección de Vivienda a Backend
		controllers.SendHomeInspectionToBackend(&homeInspection)

		fmt.Println("Inspección de Vivienda registrada en la Base de Datos")
	}
}
