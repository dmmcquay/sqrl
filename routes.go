package sqrl

import "net/http"

var prefix map[string]string

func AddRoutes(sm *http.ServeMux) {
	prefix = map[string]string{
		"all":  "/api/v0/all/",
		"com":  "/api/v0/com/",
		"ram":  "/api/v0/ram/",
		"ops":  "/api/v0/ops/",
		"swap": "/api/v0/swap/",
		//"network": "/api/v0/network/",
	}

	sm.HandleFunc(prefix["all"], all)
	sm.HandleFunc(prefix["com"], com)
	sm.HandleFunc(prefix["ram"], ram)
	sm.HandleFunc(prefix["ops"], ops)
	sm.HandleFunc(prefix["swap"], swap)
	//sm.HandleFunc(prefix["network"], network)
}
