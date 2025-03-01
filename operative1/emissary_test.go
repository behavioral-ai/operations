package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"time"
)

const (
	testDuration = time.Second
)

func officer(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher) messaging.Agent {
	return test.NewAgent("officer:" + origin.Region)
}

func ExampleEmissary() {
	ch := make(chan struct{})
	traceDispatcher := messaging.NewTraceDispatcher()
	agent := newAgent(collective.NewEphemeralResolver(), traceDispatcher)

	go func() {
		go emissaryAttend(agent, officer)
		time.Sleep(testDuration)

		agent.Shutdown()
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
