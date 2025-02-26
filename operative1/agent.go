package operative1

import (
	"github.com/behavioral-ai/caseofficer/agent"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
)

const (
	Class = "operations-operative1"
	Name  = "resiliency:agent/operations/operative"
)

// TODO : need host name
type ops struct {
	running bool
	uri     string

	emissary     *messaging.Channel
	caseOfficers *messaging.Exchange
	notifier     messaging.NotifyFunc
	dispatcher   messaging.Dispatcher
}

func cast(agent any) *ops {
	o, _ := agent.(*ops)
	return o
}

// New - create a new operative
func New() messaging.Agent {
	return newAgent(nil, nil)
}

func newAgent(notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) *ops {
	r := new(ops)
	r.uri = Name

	r.caseOfficers = messaging.NewExchange()
	r.emissary = messaging.NewEmissaryChannel()
	r.notifier = notifier
	if r.notifier == nil {
		r.notifier = collective.Resolver.Notify
	}
	r.dispatcher = dispatcher
	return r
}

// String - identity
func (o *ops) String() string { return o.Uri() }

// Uri - agent identifier
func (o *ops) Uri() string { return o.uri }

// Name - agent urn
func (o *ops) Name() string { return Name }

// Message - message the agent
func (o *ops) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	o.emissary.C <- m
}

// Run - run the agent
func (o *ops) Run() {
	if o.running {
		return
	}
	go emissaryAttend(o, agent.New)
	o.running = true
}

// Shutdown - shutdown the agent
func (o *ops) Shutdown() {
	if !o.emissary.IsClosed() {
		o.emissary.C <- messaging.Shutdown
		o.caseOfficers.Shutdown()
	}
}

func (o *ops) notify(e messaging.Event) {
	o.notifier(e)
}

func (o *ops) dispatch(channel any, event string) {
	messaging.Dispatch(o, o.dispatcher, channel, event)
}

func (o *ops) finalize() {
	o.emissary.Close()
	o.caseOfficers.Shutdown()
}
