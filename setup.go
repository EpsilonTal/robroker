package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"github.tools.sap/atom-cfs/mock-api/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func deleteBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]

	delete(brokers, brokerID)
	utils.GenerateResponse(w, 200, nil)
}

func getAllBrokers(w http.ResponseWriter, r *http.Request) {
	utils.GenerateResponse(w, 200, brokers)
}

func createBroker(w http.ResponseWriter, r *http.Request) {
	UUID := uuid.New().String()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var reqTestSettings utils.BrokerSettings
	err = json.Unmarshal(b, &reqTestSettings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reqCatalog, err := setCatalog(b, utils.CatalogKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqProvisionSettings, err := setMethodSettings(b, utils.ProvisionKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqDeprovisionSettings, err := setMethodSettings(b, utils.DeprovisionKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqBindSettings, err := setMethodSettings(b, utils.BindKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqUnbindSettings, err := setMethodSettings(b, utils.UnbindKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	url := fmt.Sprintf(utils.HostURL, UUID)

	brokerSettings := utils.BrokerSettings{
		ID:          UUID,
		Catalog:     reqCatalog,
		Provision:   reqProvisionSettings,
		Deprovision: reqDeprovisionSettings,
		Bind:        reqBindSettings,
		Unbind:      reqUnbindSettings,
		Instances:   map[string]*utils.Instance{},
		URL:         &url,
		Username:    "admin",
		Password:    "admin",
	}
	brokers[UUID] = &brokerSettings

	brokerResponse := map[string]string{
		"ID":       UUID,
		"Username": brokerSettings.Username,
		"Password": brokerSettings.Password,
		"URL":      url,
	}
	log.Printf("Creating a test with ID: %s", UUID)
	utils.GenerateResponse(w, 201, brokerResponse)
}

func setMethodSettings(b []byte, key string) (utils.MethodConfig, error) {
	value := []byte(gjson.GetBytes(b, key).Raw)
	if len(value) == 0 {
		return getDefaultMock(), nil
	}
	// Unmarshal
	var methodConfig utils.MethodConfig
	err := json.Unmarshal(value, &methodConfig)
	if err != nil {
		return utils.MethodConfig{}, err
	}
	return methodConfig, err
}

func setCatalog(b []byte, key string) (utils.Catalog, error) {
	value := []byte(gjson.GetBytes(b, key).Raw)
	if len(value) == 0 {
		return getDefaultCatalog(), nil
	}
	// Unmarshal
	var catalog utils.Catalog
	err := json.Unmarshal(value, &catalog)
	if err != nil {
		return utils.Catalog{}, err
	}
	return catalog, err
}

func getDefaultMock() utils.MethodConfig {
	return utils.MethodConfig{
		Status: 200,
		Body:   map[string]string{},
	}
}

func getDefaultCatalog() utils.Catalog {
	return utils.Catalog{
		Services: map[string]string{},
	}
}

func updateBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var reqTestSettings utils.BrokerSettings
	err = json.Unmarshal(b, &reqTestSettings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reqProvisionSettings, err := setMethodSettings(b, utils.ProvisionKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqDeprovisionSettings, err := setMethodSettings(b, utils.DeprovisionKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqBindSettings, err := setMethodSettings(b, utils.BindKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqUnbindSettings, err := setMethodSettings(b, utils.UnbindKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	reqCatalog, err := setCatalog(b, utils.CatalogKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	settings := &utils.BrokerSettings{
		ID:          brokerID,
		Name:        brokers[brokerID].Name,
		Provision:   reqProvisionSettings,
		Deprovision: reqDeprovisionSettings,
		Bind:        reqBindSettings,
		Unbind:      reqUnbindSettings,
		Catalog:     reqCatalog,
	}
	brokers[brokerID] = settings
	log.Printf("Updating a test with ID: %s", brokerID)
	utils.GenerateResponse(w, 200, brokers[brokerID])
}

func getBroker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]

	broker := brokers[brokerID]

	catalogURL := "v2/catalog"

	if broker == nil {
		utils.GenerateResponse(w, 404, nil)
	} else {
		brokerResponse := utils.BrokerSettings{
			ID:          broker.ID,
			Name:        broker.Name,
			Username:    broker.Username,
			Password:    broker.Password,
			Catalog:     catalogURL,
			URL:         broker.URL,
			Instances:   broker.Instances,
			Provision:   broker.Provision,
			Deprovision: broker.Deprovision,
			Bind:        broker.Bind,
			Unbind:      broker.Unbind,
		}
		utils.GenerateResponse(w, 200, brokerResponse)
	}
}
