package constants

import (
	"fmt"
	"log"
	"os"
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