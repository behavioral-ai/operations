package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/test"
	"time"
)

func ExampleAgent_NotFound() {
	ch := make(chan struct{})
	agent := newAgent(collective.NewEphemeralResolver(), messaging.NewTraceDispatcher())

	go func() {
		agent.Run()
		time.Sleep(testDuration * 20)

		agent.Shutdown()
		time.Sleep(testDuration * 5)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}

func ExampleAgent() {
	ch := make(chan struct{})
	agent := newAgent(nil, nil) //collective.NewEphemeralResolver(), messaging.NewTraceDispatcher())
	test.Startup()

	go func() {
		agent.Run()
		time.Sleep(testDuration * 30)

		agent.Shutdown()
		time.Sleep(testDuration * 5)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
