package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HTTP() http.Handler {
	return promhttp.Handler()
}
