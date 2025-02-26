package operative1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath = "github/behavioral-ai/agency/operative1"
)

var (
	westOrigin    = common.Origin{Region: "us-west", Host: "www.west-host1.com"}
	centralOrigin = common.Origin{Region: "us-central", Host: "www.central-host1.com"}
	opsAgent      = New()
)

func AgentMessage(event string) error {
	switch event {
	case messaging.StartupEvent:
		opsAgent.Run()
	case messaging.ShutdownEvent:
	case messaging.StartEvent:
	case messaging.StopEvent:
	case messaging.PauseEvent:
	case messaging.ResumeEvent:
		opsAgent.Message(messaging.NewMessage(messaging.ControlChannel, event))
	default:
		return errors.New(fmt.Sprintf("AgentMessage() -> [err:%v] [event:%v]\n", "error: invalid event", event))
	}
	return nil
}
