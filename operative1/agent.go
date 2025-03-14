package operative1

import (
	"github.com/behavioral-ai/caseofficer/agent"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
)

const (
	Name = "resiliency:agent/operations/operative"
)

// TODO : need host name
type agentT struct {
	running bool
	uri     string

	emissary     *messaging.Channel
	caseOfficers *messaging.Exchange
	resolver     content.Resolution
	dispatcher   messaging.Dispatcher
}

// New - create a new operative
func New() messaging.Agent {
	return newAgent(nil, nil)
}

func newAgent(resolver content.Resolution, dispatcher messaging.Dispatcher) *agentT {
	r := new(agentT)
	r.uri = Name

	r.caseOfficers = messaging.NewExchange()
	r.emissary = messaging.NewEmissaryChannel()
	if resolver == nil {
		r.resolver = content.Resolver
	} else {
		r.resolver = resolver
	}
	r.dispatcher = dispatcher
	return r
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.uri }

// Name - agent urn
func (a *agentT) Name() string { return Name }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	a.emissary.C <- m
}

// Run - run the agent
func (a *agentT) Run() {
	if a.running {
		return
	}
	go emissaryAttend(a, agent.New)
	a.running = true
}

// Shutdown - shutdown the agent
func (a *agentT) Shutdown() {
	if !a.emissary.IsClosed() {
		a.emissary.C <- messaging.Shutdown
	}
	a.running = false
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) finalize() {
	a.emissary.Close()
	a.caseOfficers.Shutdown()
}
