package servicesct

import (
	"net/http"
	"network/protocol/http/api/rest"
	"network/protocol/http/router"
)

// Redirect path (service) to corresponding policy
func Redirect() (policy router.RoutingPolicy) {
	policy.Handle(handler)
	return
}

func handler(w http.ResponseWriter, req *http.Request) {
	rest.Handlers{
		GET:    getServiceHandler(),
		POST:   createServiceHandler(),
		PATCH:  editServiceHandler(),
		DELETE: deleteServiceHandler(),
	}.Handle(w, req)
}
