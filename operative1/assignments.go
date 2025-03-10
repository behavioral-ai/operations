package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
)

func createAssignments(agent *agentT, newAgent newOfficerAgent) {
	addAssignment(agent, common.Origin{Region: common.WestRegion}, newAgent)
	addAssignment(agent, common.Origin{Region: common.CentralRegion}, newAgent)
}

func addAssignment(agent *agentT, origin common.Origin, newAgent newOfficerAgent) {
	a := newAgent(origin, agent.resolver, agent.dispatcher)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.resolver.Notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
	}
}
