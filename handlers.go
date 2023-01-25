package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.tools.sap/atom-cfs/mock-api/utils"
	"io/ioutil"
	"net/http"
	"time"
)

func getCatalog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	w = utils.GenerateResponse(w, 200, brokers[brokerID].Catalog)
}

func getInstance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]
	w = utils.GenerateResponse(w, 200, brokers[brokerID].Instances[instanceID])
}

func provision(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]
	print(instanceID)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var provisioningReq utils.ProvisionRequestBody
	err = json.Unmarshal(b, &provisioningReq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	broker := brokers[brokerID]
	broker.Instances[instanceID] = &utils.Instance{
		ID:         instanceID,
		ServiceID:  provisioningReq.ServiceID,
		PlanID:     provisioningReq.PlanID,
		CreateTime: time.Now().String(),
		Bindings:   map[string]*utils.Binding{},
	}

	w = utils.GenerateResponse(w, 201, utils.ProvisionResponseBody{})
}

func deprovision(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]

	delete(brokers[brokerID].Instances, instanceID)

	w = utils.GenerateResponse(w, 200, utils.DeprovisionResponseBody{})
}

func getBinding(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]
	bindingID := vars["bindingID"]
	w = utils.GenerateResponse(w, 200, brokers[brokerID].Instances[instanceID].Bindings[bindingID])
}

func bind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]
	bindingID := vars["bindingID"]
	print(instanceID)
	print(bindingID)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var provisioningReq utils.ProvisionRequestBody
	err = json.Unmarshal(b, &provisioningReq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	broker := brokers[brokerID]
	broker.Instances[instanceID].Bindings[bindingID] = &utils.Binding{
		ID: bindingID,
	}

	credentials := map[string]interface{}{
		"ClientID":     "client",
		"ClientSecret": "secret",
	}

	w = utils.GenerateResponse(w, 201, utils.BindResponseBody{
		Credentials: credentials,
	})
}

func unbind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brokerID := vars["brokerID"]
	instanceID := vars["instanceID"]
	bindingID := vars["bindingID"]

	delete(brokers[brokerID].Instances[instanceID].Bindings, bindingID)

	w = utils.GenerateResponse(w, 200, utils.DeprovisionResponseBody{})
}

