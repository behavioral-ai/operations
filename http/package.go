package http

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/agency/module"
	"github.com/behavioral-ai/agency/resiliency"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/uri"
	"net/http"
)

// http://localhost:8085/github/advanced-go/agency:resiliency?action=send

const (
	PkgPath = "github/advanced-go/operations/http"
	ver1    = "v1"
	ver2    = "v2"

	resiliencyResource  = "resiliency"
	healthLivenessPath  = "health/liveness"
	healthReadinessPath = "health/readiness"
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	p, err := uri.ValidateURL(r.URL, module.Domain)
	if err != nil {
		status := core.NewStatusError(http.StatusBadRequest, err)
		resp1, _ := httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return resp1, status
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case resiliencyResource:
		return resiliency.Post(r)
	case healthReadinessPath, healthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, testresource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}
