package operative1

import (
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/resiliency/guidance"
)

const (
	startAgentsEvent = "event:start-agents"
	stopAgentsEvent  = "event:stop-agents"
)

type newOfficerAgent func(origin core.Origin, handler messaging.OpsAgent, dispatcher messaging.Dispatcher) messaging.OpsAgent

// emissary attention
func emissaryAttend(agent *ops, newAgent newOfficerAgent) {
	// Agent is always running
	//agent.dispatch(messaging.StartupEvent)
	//agent.dispatch(msg.Event())
	for {
		select {
		case msg := <-agent.emissary.C:
			agent.setup(msg.Event())
			switch msg.Event() {
			case messaging.ShutdownEvent:
				agent.finalize()
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(guidance.ContentTypeCalendar) {
					agent.caseOfficers.Broadcast(msg)
					agent.dispatch(msg.Event())
				}
			case stopAgentsEvent:
				agent.caseOfficers.Shutdown()
				agent.dispatch(msg.Event())
			case startAgentsEvent:
				if agent.caseOfficers.Count() == 0 {
					createAssignments(agent, newAgent)
					agent.dispatch(msg.Event())
				}
			default:
				agent.Notify(messaging.EventErrorStatus(agent.agentId, msg))
			}
		default:
		}
	}
}
