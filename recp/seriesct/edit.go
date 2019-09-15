package seriesct

import (
	"Utils/Data/Format"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type editShowRoomBundle struct {
	Name     string            `json:"name" bson:"name"`
	Showroom roksandb.Showroom `json:"showroom" bson:"showroom"`
}

func editShowRoomHandler() (handler rest.Handler) {
	handler.Bundle(&editShowRoomBundle{})
	handler.AddPublicProcess(getEditShowRoomReq)
	handler.AddPublicProcess(editShowRoom)
	return
}

func getEditShowRoomBundle(bundle interface{}) (data *editShowRoomBundle) {
	data, _ = bundle.(*editShowRoomBundle)
	return
}

func getEditShowRoomReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditShowRoomBundle(bundle)
	if Format.Json2Struct(proto.Body, data) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func editShowRoom(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getEditShowRoomBundle(bundle)
	if err := roksandb.UpdateShowroom(data.Name, data.Showroom); err != nil {
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
