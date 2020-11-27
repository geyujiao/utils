package library

import (
	"testing"

	"github.com/vgmdj/utils/logger"
)

func Test_Tel(t *testing.T) {
	data, err := Tel("18732380601")
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func Test_TccTel(t *testing.T) {
	TccTel("18732380601")
}

func TestTelTmp(t *testing.T) {
	data, err := TelTmp("18732380601")
	if err != nil {
		t.Error(err)
	}

	logger.Info("data", data)
}