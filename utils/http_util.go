package utils

import (
	"encoding/json"
	"net/http"
)

const (
	CatalogKey = "catalog"

	ProvisionKey   = "provision"
	DeprovisionKey = "deprovision"

	BindKey   = "bind"
	UnbindKey = "unbind"

	StatusKey = "status"

	AllKey  = "all"

	HostURL = "https://robroker.cert.cfapps.stagingaws.hanavlab.ondemand.com/broker/%s/"
)

func GenerateResponse(w http.ResponseWriter, statusCode int, object interface{}) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	marshal, _ := json.Marshal(object)
	w.Write(marshal)
	return w
}
