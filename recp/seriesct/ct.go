package seriesct

import (
	"net/http"
	"network/protocol/http/api/rest"
	"network/protocol/http/router"
)

// Redirect path (showroom) to corresponding policy
func Redirect() (policy router.RoutingPolicy) {
	policy.Handle(handler)
	return
}

func handler(w http.ResponseWriter, req *http.Request) {
	rest.Handlers{
		GET:    getSeriesHandler(),
		POST:   createSeriesHandler(),
		PATCH:  editSeriesHandler(),
		DELETE: deleteSeriesHandler(),
	}.Handle(w, req)
}
