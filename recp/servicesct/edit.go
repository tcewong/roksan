package servicesct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type editServiceBundle struct {
	Name    string           `json:"name" bson:"name"`
	Service roksandb.Service `json:"service" bson:"service"`
}

func editServiceHandler() (handler rest.Handler) {
	handler.Bundle(&editServiceBundle{})
	handler.AddPublicProcess(getEditServiceReq)
	handler.AddPublicProcess(editService)
	return
}

func getEditServiceBundle(bundle interface{}) (data *editServiceBundle) {
	data, _ = bundle.(*editServiceBundle)
	return
}

func getEditServiceReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditServiceBundle(bundle)
	if Format.Json2Struct(proto.Body, data) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func editService(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditServiceBundle(bundle)
	if err := roksandb.UpdateService(data.Name, data.Service); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
