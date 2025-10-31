package main

import counterServer "github.com/jfgavin/SOMAS-CW-2025/server"

// "go run ."
func main() {
	// create server from constructor
	serv := counterServer.MakeCounterServer(5, 5, 10)
	// toggle verbose logging of messaging stats
	serv.ReportMessagingDiagnostics()
	// begin simulator
	serv.Start()
}
