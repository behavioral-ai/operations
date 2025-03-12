package operative1

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/test"
	"github.com/behavioral-ai/core/messaging"
	"time"
)

func ExampleAgent_NotFound() {
	ch := make(chan struct{})
	agent := newAgent(content.NewEphemeralResolver(), messaging.NewTraceDispatcher())

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
	agent := newAgent(nil, nil) //content.NewEphemeralResolver(), messaging.NewTraceDispatcher())
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
