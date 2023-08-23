package controllers

import (
	"DENV_Inspections_Consumer/configs"
	"DENV_Inspections_Consumer/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendHomeInspectionToBackend(homeInspection *models.HomeInspection) {
	// Endpoint HTTP POST
	postUrl := configs.BaseUrl + "/home-inspections"

	// JSON Body
	body, err := json.Marshal(homeInspection)
	if err != nil {
		panic(err)
	}

	// Send POST request
	request, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	// Set Headers
	request.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != 201 {
		panic("Error al registrar la Inspección de Vivienda en la Base de Datos")
	}
}
