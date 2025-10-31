package infra

import (
	"fmt"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type BaseCounterAgent struct {
	// embed functionality of package
	*agent.BaseAgent[ICounterAgent]
	// add additional fields for all agents in simulator
	Count int
}

// base implementation of DoCount
func (bca *BaseCounterAgent) DoCount() {
	fmt.Printf("%s is counting...\n", bca.GetID())
	bca.Count += 1
}

// base implementation of DoMessaging (just end straight away)
func (bca *BaseCounterAgent) DoMessaging() {
	bca.SignalMessagingComplete()
}

// 'correct' implementation of GetCount - override not needed
func (bca *BaseCounterAgent) GetCount() int {
	return bca.Count
}

// base implmentation of HandleCounterMessage (just ignore)
func (bca *BaseCounterAgent) HandleCounterMessage(msg *CounterMessage) {}

// constructor for CounterMessage (QoL - callable from agent)
func (bca *BaseCounterAgent) GetCounterMessage(amt int) *CounterMessage {
	return &CounterMessage{
		BaseMessage:     bca.CreateBaseMessage(),
		amountInMessage: amt,
	}
}

// constructor for BaseCounterAgent
func GetBaseCounterAgent(funcs agent.IExposedServerFunctions[ICounterAgent]) *BaseCounterAgent {
	return &BaseCounterAgent{
		BaseAgent: agent.CreateBaseAgent(funcs),
	}
}
