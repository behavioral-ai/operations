package operative1

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

type newOfficerAgent func(origin common.Origin, resolver content.Resolution, dispatcher messaging.Dispatcher) messaging.Agent

// emissary attention
func emissaryAttend(agent *agentT, newAgent newOfficerAgent) {
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
				agent.caseOfficers.Broadcast(messaging.Pause)
				content.Agent.Message(msg)
			case messaging.ResumeEvent:
				paused = false
				agent.caseOfficers.Broadcast(messaging.Resume)
				content.Agent.Message(msg)
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
