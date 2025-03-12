package operative1

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
)

func ExampleAssignments() {
	agent := newAgent(content.NewEphemeralResolver(), messaging.NewTraceDispatcher())

	createAssignments(agent, officer)

	fmt.Printf("test: createAssignments() -> [count:%v]\n", agent.caseOfficers.Count())

	agent.Shutdown()

	//Output:
	//test: createAssignments() -> [count:2]

}
