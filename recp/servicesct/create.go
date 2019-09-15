package servicesct

import (
	"Utils/Data/Format"
	"fmt"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type createSerivceBundle struct {
	service roksandb.Service
}

func createServiceHandler() (handler rest.Handler) {
	handler.Bundle(&createSerivceBundle{})
	handler.AddPublicProcess(getCreateSerivceReq)
	handler.AddPublicProcess(createService)
	return
}

func getCreateSerivceBundle(bundle interface{}) (data *createSerivceBundle) {
	data, _ = bundle.(*createSerivceBundle)
	return
}

func getCreateSerivceReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateSerivceBundle(bundle)
	if Format.Json2Struct(proto.Body, &data.service) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func createService(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateSerivceBundle(bundle)
	if err := roksandb.InsertService(data.service); err != nil {
		fmt.Println("err:", err)
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetRespHeader(http.StatusCreated)
	}
	return
}
