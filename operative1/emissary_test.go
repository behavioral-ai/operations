package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/common"
)

var (
	shutdownMsg   = messaging.NewMessage(messaging.ControlChannel, messaging.ShutdownEvent)
	dataChangeMsg = messaging.NewMessage(messaging.ControlChannel, messaging.DataChangeEvent)
	startMsg      = messaging.NewMessage(messaging.ControlChannel, messaging.StartEvent)
	stopMsg       = messaging.NewMessage(messaging.ControlChannel, messaging.StopEvent)
)

func officer(handler messaging.Agent, origin common.Origin, dispatcher messaging.Dispatcher) messaging.Agent {
	return test.NewAgent("officer:" + origin.Region)
}

func ExampleEmissary() {
	ch := make(chan struct{})
	traceDispatcher := messaging.NewTraceDispatcher()
	agent := newAgent(messaging.Notify, traceDispatcher)

	go func() {
		go emissaryAttend(agent, officer)
		//agent.Message(dataChangeMsg)
		//agent.Message(startMsg)
		//agent.Message(stopMsg)
		agent.Shutdown()
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: Trace() -> agency-operative1 : emissary event:data-change Broadcast() -> calendar data change event
	//test: dispatch(event:start-agents) -> [count>0:true]
	//test: dispatch(event:stop-agents) -> [count:0]
	//test: emissaryAttend() -> [finalized:true]

}
