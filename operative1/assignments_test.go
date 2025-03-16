package operative1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/content"
)

func ExampleAssignments() {
	agent := newAgent(content.NewEphemeralResolver(), messaging.NewTraceDispatcher())

	createAssignments(agent, officer)

	fmt.Printf("test: createAssignments() -> [count:%v]\n", agent.caseOfficers.Count())

	agent.Shutdown()

	//Output:
	//test: createAssignments() -> [count:2]

}
