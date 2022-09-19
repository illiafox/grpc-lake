package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	grpcTotalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_requests_total",
		Help: "The total number of processed grpc requests",
	})

	grpcSuccessRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_requests_success",
		Help: "The total number of successful grpc requests",
	})

	grpcInternalErrorRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_requests_error_internal",
		Help: "The total number of grpc requests with internal errors",
	})
)

func IncGrpcTotalRequests() {
	grpcTotalRequests.Inc()
}

func IncGrpcSuccessRequests() {
	grpcSuccessRequests.Inc()
}

func IncGrpcErrorRequests() {
	grpcInternalErrorRequests.Inc()
}
