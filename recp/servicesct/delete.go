package servicesct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type deleteServiceBundle struct {
	id string
}

func deleteServiceHandler() (handler rest.Handler) {
	handler.Bundle(&deleteServiceBundle{})
	handler.PublicErrProcess(getDeleteServiceErr)
	handler.AddPublicProcess(getDeleteServiceReq)
	handler.AddPublicProcess(getDeleteService)
	return
}

func getDeleteServiceBundle(bundle interface{}) (data *deleteServiceBundle) {
	data, _ = bundle.(*deleteServiceBundle)
	return
}

func getDeleteServiceErr(proto *rest.Protocol, bundle interface{}, err error) {
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getDeleteServiceReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteServiceBundle(bundle)
	if data.id = Get.UrlVal("id", proto.Req, ""); data.id == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getDeleteService(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteServiceBundle(bundle)
	if err := roksandb.DeleteSeries(data.id); err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
