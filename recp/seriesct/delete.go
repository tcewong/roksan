package seriesct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type deleteSeriesBundle struct {
	id string
}

func deleteSeriesHandler() (handler rest.Handler) {
	handler.Bundle(&deleteSeriesBundle{})
	handler.PublicErrProcess(getDeleteSeriesErr)
	handler.AddPublicProcess(getDeleteSeriesReq)
	handler.AddPublicProcess(getDeleteSeries)
	return
}

func getDeleteSeriesBundle(bundle interface{}) (data *deleteSeriesBundle) {
	data, _ = bundle.(*deleteSeriesBundle)
	return
}

func getDeleteSeriesErr(proto *rest.Protocol, bundle interface{}, err error) {
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getDeleteSeriesReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteSeriesBundle(bundle)
	if data.id = Get.UrlVal("id", proto.Req, ""); data.id == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getDeleteSeries(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteSeriesBundle(bundle)
	if err := roksandb.DeleteSeries(data.id); err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
