package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
)

type newOfficerAgent func(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher) messaging.Agent

// emissary attention
func emissaryAttend(agent *ops, newAgent newOfficerAgent) {
	agent.dispatch(agent.emissary, messaging.StartupEvent)
	paused := false
	createAssignments(agent, newAgent)

	for {
		select {
		case msg := <-agent.emissary.C:
			agent.dispatch(agent.emissary, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.StopEvent:
				if !paused {
					agent.caseOfficers.Shutdown()
				}
			case messaging.StartEvent:
				if !paused {
					if agent.caseOfficers.Count() == 0 {
						createAssignments(agent, newAgent)
					}
				}
			case messaging.ShutdownEvent:
				agent.finalize()
				return
			default:
			}
		default:
		}
	}
}
