package http

import (
	"github.com/behavioral-ai/operations/module"
	"net/http"
)

// http://localhost:8080/resiliency?event=startup

const (
	resiliencyResource = "resiliency"
	eventKey           = "event"
)

// Exchange - HTTP exchange function
func Exchange(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"+resiliencyResource {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: invalid path"))
		return
	}
	values := r.URL.Query()
	if len(values) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: no query args"))
		return
	}
	event := values.Get(eventKey)
	if event == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: event query key not found"))
		return
	}
	err := module.AgentMessage(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*

func Exchange(w http.ResponseWriter, r *http.Request)  {
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

*/
