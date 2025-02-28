package operative1

import (
	"github.com/behavioral-ai/core/messaging"
)

func createAssignments(agent *ops, newAgent newOfficerAgent) {
	//if newAgent == nil {
	//agent.Notify(core.NewStatusError(core.StatusInvalidArgument, errors.New("error: initialize newAgent is nil")))
	//	return
	//}
	a := newAgent(agent, westOrigin, agent.dispatcher)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
	}
	a = newAgent(agent, centralOrigin, agent.dispatcher)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
	}
}
