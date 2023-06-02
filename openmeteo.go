package openmeteo

import (
	"errors"
	"fmt"
	"github.com/prongbang/callx"
	"net/http"
)

type OpenMeteo interface {
	Execute(parameter Parameter) (string, error)
}

type openMeteo struct {
	CallX callx.CallX
}

func (o *openMeteo) Execute(parameter Parameter) (string, error) {
	path := fmt.Sprintf("/v1/forecast?%s", parameter.ToQuery())
	resp := o.CallX.Get(path)
	if resp.Code == http.StatusOK {
		return string(resp.Data), nil
	}
	return "", errors.New(http.StatusText(resp.Code))
}

func New(args ...callx.CallX) OpenMeteo {
	var c callx.CallX
	if len(args) > 0 {
		c = args[0]
	} else {
		cfg := callx.Config{
			BaseURL: "https://api.open-meteo.com",
			Timeout: 60,
		}
		c = callx.New(cfg)
	}
	return &openMeteo{
		CallX: c,
	}
}
