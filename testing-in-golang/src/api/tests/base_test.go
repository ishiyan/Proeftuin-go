package test

import (
	"testing"
	"os"
	"proeftuin/testing-in-golang/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
