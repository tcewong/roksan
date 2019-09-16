package reviewct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type editReviewBundle struct {
	Name   string          `json:"name" bson:"name"`
	Review roksandb.Review `json:"review" bson:"review"`
}

func editReviewHandler() (handler rest.Handler) {
	handler.Bundle(&editReviewBundle{})
	handler.AddPublicProcess(getEditReviewmReq)
	handler.AddPublicProcess(editReview)
	return
}

func getEditReviewBundle(bundle interface{}) (data *editReviewBundle) {
	data, _ = bundle.(*editReviewBundle)
	return
}

func getEditReviewmReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditReviewBundle(bundle)
	if Format.Json2Struct(proto.Body, data) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func editReview(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditReviewBundle(bundle)
	if err := roksandb.UpdateReview(data.Name, data.Review); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
