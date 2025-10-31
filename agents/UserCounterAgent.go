package agents

import (
	"fmt"
	"math/rand"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	"github.com/jfgavin/SOMAS-Group-3/infra"
)

// third tier of composition - embed BaseCounterAgent..
// ... and add 'user specific' fields
type UserCounterAgent struct {
	*infra.BaseCounterAgent
	amount int
}

// user implementation of DoCount ('strategic')
func (uca *UserCounterAgent) DoCount() {
	fmt.Printf("%s is counting...\n", uca.GetID())
	uca.Count += uca.amount
}

// user implementation of DoMessaging ('strategic')
func (uca *UserCounterAgent) DoMessaging() {
	uca.BroadcastMessage(uca.GetCounterMessage(uca.amount))
	uca.SignalMessagingComplete()
}

// user implementation of Handler ('strategic' - print message to console)
func (uca *UserCounterAgent) HandleCounterMessage(msg *infra.CounterMessage) {
	fmt.Printf("Receiver: %s, Sender: %s, Amount: %d\n", uca.GetID(), msg.GetSender(), msg.GetAmountInMessage())
}

// constructor for UserCounterAgent
func GetUserCounterAgent(funcs agent.IExposedServerFunctions[infra.ICounterAgent]) *UserCounterAgent {
	return &UserCounterAgent{
		BaseCounterAgent: infra.GetBaseCounterAgent(funcs),
		amount:           rand.Intn(10),
	}
}
