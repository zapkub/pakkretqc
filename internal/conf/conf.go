package conf

import "os"

func ALMEndpoint() string {
	v, ok := os.LookupEnv("PAKKRETQC_ALM_ENDPOINT")
	if ok {
		return v
	}
	panic("PAKKRETQC_ALM_ENDPOINT is not set")
}
