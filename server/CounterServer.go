package counterServer

import (
	"fmt"
	"time"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	"github.com/jfgavin/SOMAS-CW-2025/agents"
	"github.com/jfgavin/SOMAS-CW-2025/infra"
)

type CounterServer struct {
	// embed functionality from package...
	// ...and tell BaseServer we're using ICounterAgents
	*server.BaseServer[infra.ICounterAgent]
}

// RunTurn implementation - Count, and then Message
func (cs *CounterServer) RunTurn(i, j int) {
	fmt.Printf("Running iteration %d, turn %d\n\n", i, j)
	fmt.Println("COUNTING:")
	cs.RunCountTurn()
	fmt.Println("MESSAGING:")
	cs.RunMessagingTurn()
}

// just implement GameRunner
func (cs *CounterServer) RunStartOfIteration(int) {}

// just implement GameRunner
func (cs *CounterServer) RunEndOfIteration(int) {}

// make all agents count
func (cs *CounterServer) RunCountTurn() {
	for _, ag := range cs.GetAgentMap() {
		ag.DoCount()
	}
}

// make all agents message
func (cs *CounterServer) RunMessagingTurn() {
	for _, ag := range cs.GetAgentMap() {
		ag.DoMessaging()
	}
}

// override start
func (cs *CounterServer) Start() {
	// steal method from package...
	cs.BaseServer.Start()

	// ...and add some more functionality
	for _, ag := range cs.GetAgentMap() {
		fmt.Printf("%s has a final count of: %d\n", ag.GetID(), ag.GetCount())
	}
}

// constructor for CounterServer
func MakeCounterServer(iterations, turns, numAgents int) *CounterServer {
	// embed BaseServer: maxTimeout = 10ms, maxThreads = 100
	serv := &CounterServer{
		BaseServer: server.CreateBaseServer[infra.ICounterAgent](iterations, turns, 10*time.Millisecond, 100),
	}
	// set GameRunner to bind RunTurn to BaseServer
	serv.SetGameRunner(serv)

	// inject agents (50% BaseCounter)...
	for i := 0; i < numAgents/2; i++ {
		serv.AddAgent(infra.GetBaseCounterAgent(serv))
	}

	// ...and 50% UserCounter
	for i := 0; i < numAgents/2; i++ {
		serv.AddAgent(agents.GetUserCounterAgent(serv))
	}
	return serv
}
