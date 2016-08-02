package sqrl

import "net/http"

var prefix map[string]string

func AddRoutes(sm *http.ServeMux) {
	prefix = map[string]string{
		"all":     "/api/v0/all/",
		"com":     "/api/v0/com/",
		"ros":     "/api/v0/ros/",
		"network": "/api/v0/network/",
	}

	sm.HandleFunc(prefix["all"], all)
	sm.HandleFunc(prefix["com"], com)
	sm.HandleFunc(prefix["ros"], ros)
	sm.HandleFunc(prefix["network"], network)
}
