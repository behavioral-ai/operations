package operative1

import (
	"github.com/behavioral-ai/core/messaging"
)

func createAssignments(agent *ops, newAgent newOfficerAgent) {
	a := newAgent(westOrigin, agent.resolver, agent.dispatcher)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.resolver.Notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
	}
	a = newAgent(centralOrigin, agent.resolver, agent.dispatcher)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.resolver.Notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
	}
}
