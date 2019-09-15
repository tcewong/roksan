package brandct

import (
	"Utils/Data/Format"
	"Utils/Data/Get"
	"fmt"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type showRoomBundle struct {
	Showrooms []roksandb.Showroom `json:"showroom" bson:"showroom"`
	lang      string
}

func getShowRoomHandler() (handler rest.Handler) {
	handler.Bundle(&showRoomBundle{})
	handler.PublicErrProcess(getShowRoomErr)
	handler.AddPublicProcess(getShowRoomReq)
	handler.AddPublicProcess(getShowRoom)
	return
}

func getShowRoomBundle(bundle interface{}) (data *showRoomBundle) {
	data, _ = bundle.(*showRoomBundle)
	return
}

func getShowRoomErr(proto *rest.Protocol, bundle interface{}, err error) {
	fmt.Println("err:", err)
	proto.SetRespHeader(http.StatusServiceUnavailable)
	return
}

func getShowRoomReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getShowRoomBundle(bundle)
	if data.lang = Get.UrlVal("lang", proto.Req, ""); data.lang == "" {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func getShowRoom(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getShowRoomBundle(bundle)
	data.Showrooms, err = roksandb.FindShowrooms(data.lang)
	if err == nil {
		proto.SetResp(http.StatusAccepted, Format.Struct2Json(data))
	}
	return
}
