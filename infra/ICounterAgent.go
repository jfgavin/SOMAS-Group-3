package infra

import "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"

type ICounterAgent interface {
	// embed functionality from package
	agent.IAgent[ICounterAgent]
	// perfom counting action
	DoCount()
	// getter for count value (agents are injected as interfaces,...
	// ...not structs, so Count is not visible)
	GetCount() int
	// perform messaging action
	DoMessaging()
	// get CounterMessage (QoL - convenient to call from agent)
	GetCounterMessage(int) *CounterMessage
	// handler for CounterMessage (visitor design pattern)
	HandleCounterMessage(*CounterMessage)
}
