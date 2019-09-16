package seriesct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type editSeriesBundle struct {
	ID     string          `json:"id" bson:"id"`
	Series roksandb.Series `json:"series" bson:"series"`
}

func editSeriesHandler() (handler rest.Handler) {
	handler.Bundle(&editSeriesBundle{})
	handler.AddPublicProcess(getEditSeriesReq)
	handler.AddPublicProcess(editSeries)
	return
}

func getEditSeriesBundle(bundle interface{}) (data *editSeriesBundle) {
	data, _ = bundle.(*editSeriesBundle)
	return
}

func getEditSeriesReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditSeriesBundle(bundle)
	if Format.Json2Struct(proto.Body, data) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func editSeries(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditSeriesBundle(bundle)
	if err := roksandb.UpdateSeries(data.ID, data.Series); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
