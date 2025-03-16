package module

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/test"
	"github.com/behavioral-ai/operations/operative1"
)

const (
	ResiliencyPath = "/resiliency"
)

func Startup(hostName string) {
	test.Startup()
	AgentMessage(messaging.StartupEvent)
}

func AgentMessage(event string) error {
	return operative1.Message(event)
}
