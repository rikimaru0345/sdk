package sdk_test

import (
	"fmt"
	"github.com/rikimaru0345/sdk"
	"os"
	"testing"
)

func getClient() *sdk.Client {
	addr := "http://localhost:3000"
	user := "admin"
	pass := "admin"

	if a := os.Getenv("GRAFANA_ADDR"); a != "" {
		addr = a
	}
	if u := os.Getenv("GRAFANA_USER"); u != "" {
		user = u
	}
	if p := os.Getenv("GRAFANA_PASS"); p != "" {
		pass = p
	}

	return sdk.NewClient(addr, fmt.Sprintf("%s:%s", user, pass), sdk.DefaultHTTPClient)
}

func shouldSkip(t *testing.T) {
	if v := os.Getenv("GRAFANA_INTEGRATION"); v != "1" {
		t.Skipf("skipping because GRAFANA_INTEGRATION is %s, not 1", v)
	}
}
