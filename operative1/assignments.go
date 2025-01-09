package operative1

import (
	"errors"
	"github.com/behavioral-ai/core/core"
)

func createAssignments(agent *ops, newAgent newOfficerAgent) {
	if newAgent == nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, errors.New("error: initialize newAgent is nil")))
		return
	}
	a := newAgent(westOrigin, agent, agent.dispatcher)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
	a = newAgent(centralOrigin, agent, agent.dispatcher)
	err = agent.caseOfficers.Register(a)
	if err != nil {
		agent.Notify(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		a.Run()
	}
}
