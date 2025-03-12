package operative1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
)

const (
	PkgPath = "github/behavioral-ai/agency/operative1"
)

var (
	opsAgent = New()
)

func AgentMessage(event string) error {
	switch event {
	case messaging.StartupEvent:
		opsAgent.Run()
	case messaging.ShutdownEvent:
	case messaging.StartEvent:
	case messaging.StopEvent:
	case messaging.PauseEvent:
		opsAgent.Message(messaging.Pause)
	case messaging.ResumeEvent:
		opsAgent.Message(messaging.Resume)
	default:
		return errors.New(fmt.Sprintf("AgentMessage() -> [%v] [%v]", "error: invalid event", event))
	}
	return nil
}
