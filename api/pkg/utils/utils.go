package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	PushgatewayUrl string = "PUSH_GATEWAY_URL"
	JobName string =  "JOB_NAME"
)

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("%s not set\n", key)
		panic(fmt.Sprintf("please set %s variable", key))
	} else {
		log.Printf("%s=%s\n", key, val)
		return val
	}
}

func CombineCollectors(arrays ...[]prometheus.Collector) []prometheus.Collector {
	// Calculate the total number of collectors across all arrays
    totalLen := 0
    for _, arr := range arrays {
        totalLen += len(arr)
    }

    // Create a new array to hold the combined collectors
    combined := make([]prometheus.Collector, 0, totalLen)

    // Iterate through the input arrays and append their elements to the combined array
    for _, arr := range arrays {
        combined = append(combined, arr...)
    }

    return combined
}