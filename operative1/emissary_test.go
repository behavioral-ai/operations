package operative1

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/common"
	"time"
)

const (
	testDuration = time.Second
)

func officer(origin common.Origin, resolver content.Resolution, dispatcher messaging.Dispatcher) messaging.Agent {
	return test.NewAgent("officer:" + origin.Region)
}

func ExampleEmissary() {
	ch := make(chan struct{})
	traceDispatcher := messaging.NewTraceDispatcher()
	agent := newAgent(content.NewEphemeralResolver(), traceDispatcher)

	go func() {
		go emissaryAttend(agent, officer)
		time.Sleep(testDuration * 2)

		agent.Shutdown()
		time.Sleep(testDuration * 2)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
