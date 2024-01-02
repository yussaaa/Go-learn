//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServerInfo(servers map[string]int) {
	statuses := make(map[string]int)

	serverNum := len(servers)
	fmt.Println("Number of servers:", serverNum)

	for _, status := range servers {
		// if status == Online {
		// 	statuses["Online"]++
		// } else if status == Offline {
		// 	statuses["Offline"]++
		// } else if status == Maintenance {
		// 	statuses["Maintenance"]++
		// } else if status == Retired {
		// 	statuses["Retired"]++
		// } else {
		// 	fmt.Println("Wrong status code!")

		switch status {
		case Online:
			statuses["Online"]++
		case Offline:
			statuses["Offline"]++
		case Maintenance:
			statuses["Maintenance"]++
		case Retired:
			statuses["Retired"]++
		default:
			{
				fmt.Println("Wrong status code!")
			}
		}
	}
	fmt.Println("Number of servers for each status:\n   ", statuses)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]int)
	for _, server := range servers {
		serverMap[server] = Online
	}
	displayServerInfo(serverMap)

	serverMap["darkstar"] = Retired
	serverMap["aiur"] = Offline
	displayServerInfo(serverMap)

	for _, server := range servers {
		serverMap[server] = Offline
	}
	displayServerInfo(serverMap)

	for _, server := range servers {
		serverMap[server] = Maintenance
	}
	displayServerInfo(serverMap)
}
