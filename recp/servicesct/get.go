package servicesct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"fmt"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type serviceBundle struct {
	Services []roksandb.Service `json:"service" bson:"service"`
	lang     string
}

func getServiceHandler() (handler rest.Handler) {
	handler.Bundle(&serviceBundle{})
	handler.PublicErrProcess(getShowRoomErr)
	handler.AddPublicProcess(getServiceReq)
	handler.AddPublicProcess(getService)
	return
}

func getServiceBundle(bundle interface{}) (data *serviceBundle) {
	data, _ = bundle.(*serviceBundle)
	return
}

func getShowRoomErr(proto *rest.Protocol, bundle interface{}, err error) {
	fmt.Println("err:", err)
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getServiceReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getServiceBundle(bundle)
	if data.lang = Get.UrlVal("lang", proto.Req, ""); data.lang == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getService(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getServiceBundle(bundle)
	data.Services, err = roksandb.FindServices(data.lang)
	if err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
