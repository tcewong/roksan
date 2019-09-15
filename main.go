package main

import (
	"network/protocol/http"
	"network/protocol/http/router"
	"roksan/recp/brandct"
	"roksan/recp/reviewct"
	"roksan/recp/seriesct"
	"roksan/recp/servicesct"
	"roksan/recp/showroomct"
	"roksan/recp/webct"
)

func main() {
	rt := router.New()

	rt.Route("web", webct.Redirect())
	rt.Route("showroom", showroomct.Redirect())
	rt.Route("series", seriesct.Redirect())
	rt.Route("service", servicesct.Redirect())
	rt.Route("review", reviewct.Redirect())
	rt.Route("brand", brandct.Redirect())

	if err := http.NewListener(8080, rt).ListenAndServe(); err != nil {
		panic(err)
	}
}
