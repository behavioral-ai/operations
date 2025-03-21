package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/messaging/messagingtest"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/content"
	"time"
)

const (
	testDuration = time.Second * 4
)

func officer(origin common.Origin, resolver content.Resolution, dispatcher messaging.Dispatcher) messaging.Agent {
	return messagingtest.NewAgent("officer:" + origin.Region)
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
