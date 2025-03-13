package module

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/test"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operations/operative1"
)

const (
	Domain         = "github/behavioral-ai/operations"
	ResiliencyPath = "/resiliency"
)

func Startup(hosts []string, do content.HttpExchange, appHostName string) {
	content.Startup(hosts, do, appHostName)
	test.Startup()
	AgentMessage(messaging.StartupEvent)
}

func AgentMessage(event string) error {
	return operative1.Message(event)
}
