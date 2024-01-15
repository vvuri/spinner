package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func ExampleCPUTimeMetric() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	go func() {
		for i := 0; i < xTimes; i++ {
			err := doOperation()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	err := http.ListenAndServe(
		":8080",
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	)
	if err != nil {
		fmt.Println(err)
	}
}

/*
	go get github.com/prometheus/client_golang

*/
