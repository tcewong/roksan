package reviewct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type reviewBundle struct {
	Reviews []roksandb.Review `json:"review" bson:"review"`
	lang    string
}

func getReviewHandler() (handler rest.Handler) {
	handler.Bundle(&reviewBundle{})
	handler.PublicErrProcess(getReviewErr)
	handler.AddPublicProcess(getReviewReq)
	handler.AddPublicProcess(getReview)
	return
}

func getReviewBundle(bundle interface{}) (data *reviewBundle) {
	data, _ = bundle.(*reviewBundle)
	return
}

func getReviewErr(proto *rest.Protocol, bundle interface{}, err error) {
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getReviewReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getReviewBundle(bundle)
	if data.lang = Get.UrlVal("lang", proto.Req, ""); data.lang == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getReview(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getReviewBundle(bundle)
	data.Reviews, err = roksandb.FindReviews(data.lang)
	if err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
