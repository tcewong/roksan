package main

import (
	"network/protocol/http"
	"network/protocol/http/router"
	"roksan/recp/showroomct"
	"roksan/recp/webct"
)

func main() {
	rt := router.New()

	rt.Route("web", webct.Redirect())
	rt.Route("showroom", showroomct.Redirect())

	if err := http.NewListener(8080, rt).ListenAndServe(); err != nil {
		panic(err)
	}
}
