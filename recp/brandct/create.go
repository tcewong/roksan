package brandct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type createBrandBundle struct {
	brand roksandb.Brand
}

func createBrandHandler() (handler rest.Handler) {
	handler.Bundle(&createBrandBundle{})
	handler.AddPublicProcess(getCreateBrandReq)
	handler.AddPublicProcess(createBrand)
	return
}

func getCreateBrandBundle(bundle interface{}) (data *createBrandBundle) {
	data, _ = bundle.(*createBrandBundle)
	return
}

func getCreateBrandReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateBrandBundle(bundle)
	if Format.Json2Struct(proto.Body, &data.brand) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func createBrand(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateBrandBundle(bundle)
	if err := roksandb.InsertBrand(data.brand); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetRespHeader(http.StatusCreated)
	}
	return
}
