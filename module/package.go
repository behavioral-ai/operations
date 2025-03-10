package module

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operations/operative1"
)

const (
	Domain         = "github/behavioral-ai/operations"
	ResiliencyPath = "/resiliency"
)

func Startup(hosts []string, do collective.HttpExchange, appHostName string) {
	collective.Startup(hosts, do, appHostName)
	AgentMessage(messaging.StartupEvent)
}

func AgentMessage(event string) error {
	return operative1.AgentMessage(event)
}
