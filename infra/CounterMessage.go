package infra

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/message"
)

type CounterMessage struct {
	// embed functionality from package
	message.BaseMessage
	// add additional fields
	amountInMessage int
}

// getter for private field
func (cm *CounterMessage) GetAmountInMessage() int {
	return cm.amountInMessage
}

// override of InvokeHandler (visitor pattern)
func (cm *CounterMessage) InvokeMessageHandler(agent ICounterAgent) {
	agent.HandleCounterMessage(cm)
}
