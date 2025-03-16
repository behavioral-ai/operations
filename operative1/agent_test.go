package operative1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/content"
	"github.com/behavioral-ai/domain/test"
	"time"
)

func _ExampleAgent_NotFound() {
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
	dispatcher := messaging.NewFilteredTraceDispatcher([]string{messaging.ResumeEvent, messaging.PauseEvent}, "")
	agent := newAgent(nil, dispatcher) //content.NewEphemeralResolver(), messaging.NewTraceDispatcher())
	test.Startup()

	go func() {
		agent.Run()
		time.Sleep(testDuration * 6)
		agent.Message(messaging.Pause)
		time.Sleep(testDuration * 6)
		agent.Message(messaging.Resume)
		time.Sleep(testDuration * 6)
		agent.Shutdown()
		time.Sleep(testDuration * 4)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
