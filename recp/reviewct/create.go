package reviewct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type createReviewBundle struct {
	review roksandb.Review
}

func createReviewHandler() (handler rest.Handler) {
	handler.Bundle(&createReviewBundle{})
	handler.AddPublicProcess(getCreateReviewReq)
	handler.AddPublicProcess(createReview)
	return
}

func getCreateReviewBundle(bundle interface{}) (data *createReviewBundle) {
	data, _ = bundle.(*createReviewBundle)
	return
}

func getCreateReviewReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateReviewBundle(bundle)
	if Format.Json2Struct(proto.Body, &data.review) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func createReview(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateReviewBundle(bundle)
	if err := roksandb.InsertReview(data.review); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetRespHeader(http.StatusCreated)
	}
	return
}
