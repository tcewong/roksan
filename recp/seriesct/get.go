package seriesct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"fmt"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type seriesBundle struct {
	Serieses []roksandb.Series `json:"series" bson:"series"`
	name     string
}

func getSeriesHandler() (handler rest.Handler) {
	handler.Bundle(&seriesBundle{})
	handler.PublicErrProcess(getSeriesErr)
	handler.AddPublicProcess(getSeriesReq)
	handler.AddPublicProcess(getSeries)
	return
}

func getSeriesBundle(bundle interface{}) (data *seriesBundle) {
	data, _ = bundle.(*seriesBundle)
	return
}

func getSeriesErr(proto *rest.Protocol, bundle interface{}, err error) {
	fmt.Println("err:", err)
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getSeriesReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getSeriesBundle(bundle)
	if data.name = Get.UrlVal("name", proto.Req, ""); data.name == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getSeries(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getSeriesBundle(bundle)
	data.Serieses, err = roksandb.FindSerieses(data.name)
	if err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
