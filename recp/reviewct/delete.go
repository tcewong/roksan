package reviewct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type deleteReviewBundle struct {
	name string
}

func deleteReviewHandler() (handler rest.Handler) {
	handler.Bundle(&deleteReviewBundle{})
	handler.PublicErrProcess(getDeleteReviewErr)
	handler.AddPublicProcess(getDeleteReviewReq)
	handler.AddPublicProcess(getDeleteReview)
	return
}

func getDeleteReviewBundle(bundle interface{}) (data *deleteReviewBundle) {
	data, _ = bundle.(*deleteReviewBundle)
	return
}

func getDeleteReviewErr(proto *rest.Protocol, bundle interface{}, err error) {
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getDeleteReviewReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteReviewBundle(bundle)
	if data.name = Get.UrlVal("name", proto.Req, ""); data.name == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getDeleteReview(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteReviewBundle(bundle)
	if err := roksandb.DeleteShowrooms(data.name); err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
