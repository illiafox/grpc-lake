package http

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"server/internal/metrics"
)

func NewServer(host string, port int) *http.Server {
	router := http.NewServeMux()

	// pprof
	{
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/heap", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	// prometheus metrics
	{
		router.Handle("/metrics", metrics.HTTP())
	}

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}

	return server
}
