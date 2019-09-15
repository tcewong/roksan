package reviewct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type deleteShowRoomBundle struct {
	name string
}

func deleteShowRommHandler() (handler rest.Handler) {
	handler.Bundle(&deleteShowRoomBundle{})
	handler.PublicErrProcess(getDeleteShowRoomErr)
	handler.AddPublicProcess(getDeleteShowRoomReq)
	handler.AddPublicProcess(getDeleteShowRoom)
	return
}

func getDeleteShowRoomBundle(bundle interface{}) (data *deleteShowRoomBundle) {
	data, _ = bundle.(*deleteShowRoomBundle)
	return
}

func getDeleteShowRoomErr(proto *rest.Protocol, bundle interface{}, err error) {
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getDeleteShowRoomReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteShowRoomBundle(bundle)
	if data.name = Get.UrlVal("name", proto.Req, ""); data.name == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getDeleteShowRoom(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getDeleteShowRoomBundle(bundle)
	if err := roksandb.DeleteShowrooms(data.name); err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
