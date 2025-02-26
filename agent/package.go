package agent

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavorial-ai/operations/operative1"
)

var opsAgent messaging.Agent

func Run() {
	opsAgent = operative1.New()
	opsAgent.Run()
}
