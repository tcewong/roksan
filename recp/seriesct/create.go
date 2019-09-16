package seriesct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type createSeriesBundle struct {
	series roksandb.Series
}

func createSeriesHandler() (handler rest.Handler) {
	handler.Bundle(&createSeriesBundle{})
	handler.AddPublicProcess(getCreateSeriesReq)
	handler.AddPublicProcess(createSeries)
	return
}

func getCreateSeriesBundle(bundle interface{}) (data *createSeriesBundle) {
	data, _ = bundle.(*createSeriesBundle)
	return
}

func getCreateSeriesReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateSeriesBundle(bundle)
	if Format.Json2Struct(proto.Body, &data.series) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func createSeries(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateSeriesBundle(bundle)
	if err := roksandb.InsertSeries(data.series); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetRespHeader(http.StatusCreated)
	}
	return
}
