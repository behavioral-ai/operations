package operative1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/content"
)

func createAssignments(agent *agentT, resolver *content.Resolution, newAgent newOfficerAgent) {
	addAssignment(agent, common.Origin{Region: common.WestRegion}, resolver, newAgent)
	addAssignment(agent, common.Origin{Region: common.CentralRegion}, resolver, newAgent)
}

func addAssignment(agent *agentT, origin common.Origin, resolver *content.Resolution, newAgent newOfficerAgent) {
	a := newAgent(origin, agent.activity, agent.notifier, agent.dispatcher)
	err := agent.caseOfficers.Register(a)
	if err != nil {
		agent.notify(messaging.NewStatusError(messaging.StatusInvalidArgument, err, agent.Uri()))
	} else {
		a.Run()
		agent.addActivity(messaging.ActivityItem{
			Agent:   agent,
			Event:   "event:add-assignment",
			Source:  messaging.Emissary,
			Content: fmt.Sprintf("added assignment: %v", origin),
		})

	}
}
