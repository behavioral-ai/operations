package operative1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"reflect"
	"time"
)

func Example_cast() {
	a := New()
	fmt.Printf("test: cast_string() -> %v\n", a)

	opsT := cast(a)
	fmt.Printf("test: cast_ops() -> %v\n", reflect.TypeOf(opsT))

	a.Shutdown()

	//Output:
	//test: cast_string() -> resiliency:agent/operations/operative
	//test: cast_ops() -> *operative1.ops

}

func ExampleAgent() {
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
