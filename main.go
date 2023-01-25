package main

import (
	"github.com/gorilla/mux"
	"github.tools.sap/atom-cfs/mock-api/utils"
	"log"
	"net/http"
	"sync"
)

var (
	brokers map[string]*utils.BrokerSettings
	mutex   *sync.Mutex
)

func main() {
	brokers = make(map[string]*utils.BrokerSettings)

	router := mux.NewRouter()
	log.Printf("starting ro-broker server")

	/*
		POST:
		 Create a new test configuration, using BrokerSettings struct
		 Response will contain the ID of the test config
		DELETE:
		 Optional flags:
			* all=true: delete all test configurations
			* id=uuid: delete single test configuration
	*/
	router.HandleFunc("/broker/{brokerID}", getBroker).Methods("GET")
	router.HandleFunc("/broker/{brokerID}", updateBroker).Methods("PATCH")
	router.HandleFunc("/broker/{brokerID}", deleteBroker).Methods("DELETE")
	router.HandleFunc("/broker", createBroker).Methods("POST")
	router.HandleFunc("/broker", getAllBrokers).Methods("GET")

	/*
		Mandatory flag:
		* id=testUUID: with define which test config to work with
		POST:
		 Execute create with the test config
		DELETE:
		 Execute delete with the test config
	*/
	router.HandleFunc("/broker/{brokerID}/v2/catalog", getCatalog).Methods("GET")

	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}/service_bindings/{bindingID}", bind).Methods("PUT")
	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}/service_bindings/{bindingID}", unbind).Methods("DELETE")
	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}/service_bindings/{bindingID}", getBinding).Methods("GET")

	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}", provision).Methods("PUT")
	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}", deprovision).Methods("DELETE")
	router.HandleFunc("/broker/{brokerID}/v2/service_instances/{instanceID}", getInstance).Methods("GET")

	router.HandleFunc("/broker/{brokerID}", getCatalog).Methods("GET")

	//http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
