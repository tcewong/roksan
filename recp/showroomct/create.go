package showroomct

import (
	"Utils/Data/Format"
	"fmt"
	"net/http"
	"network/protocol/http/api/rest"
	"roksan/roksandb"
)

type createShowRoomBundle struct {
	showroom roksandb.Showroom
}

func createShowRoomHandler() (handler rest.Handler) {
	handler.Bundle(&createShowRoomBundle{})
	handler.AddPublicProcess(getCreateShowRoomReq)
	handler.AddPublicProcess(createShowRoom)
	return
}

func getCreateShowRoomBundle(bundle interface{}) (data *createShowRoomBundle) {
	data, _ = bundle.(*createShowRoomBundle)
	return
}

func getCreateShowRoomReq(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateShowRoomBundle(bundle)
	if Format.Json2Struct(proto.Body, &data.showroom) != nil {
		proto.SetRespHeader(http.StatusNotAcceptable)
	}
	return
}

func createShowRoom(proto *rest.Protocol, bundle interface{}) (err error) {
	data := getCreateShowRoomBundle(bundle)
	if err := roksandb.InsertShowroom(data.showroom); err != nil {
		fmt.Println("err:", err)
		proto.SetRespHeader(http.StatusServiceUnavailable)
	} else {
		proto.SetRespHeader(http.StatusCreated)
	}
	return
}
