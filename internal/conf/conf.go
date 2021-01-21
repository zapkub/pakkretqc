package conf

import (
	"os"
	"time"
)

func ALMEndpoint() string {
	v, ok := os.LookupEnv("PAKKRETQC_ALM_ENDPOINT")
	if ok {
		return v
	}
	panic("PAKKRETQC_ALM_ENDPOINT is not set")
}

var loc, _ = time.LoadLocation("Asia/Bangkok")

func Location() *time.Location {
	return loc
}
