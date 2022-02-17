package arrays

////
//   For connections = [["192.167.0.0", "192.167.0.1"], ["192.167.0.2", "192.167.0.0"], ["192.167.0.0", "192.167.0.3"]] and toggleIps = ["192.167.0.1", "192.167.0.0", "192.167.0.2", "192.167.0.0",   "0.0.0.0"], the output should be solution(connections, toggleIps) = [0, 1, 1, 2, 0]

// 	1. initially all the connections are in inactive state
// 	2. a connection is said to be active when both the device ip are toggled to active.
// 	3. from the list of toggle ips find the impacted connection for each toggle.
// 	4. for the 192.167.0.1 toggle ip the the connection the impacted connection is 0 as  "192.167.0.1" in ["192.167.0.0", "192.167.0.1"] will become active , there is no active connection
// 	5. for "192.167.0.0", toggle ips the impacted connections are 1, as "192.167.0.0" is present in ["192.167.0.0", "192.167.0.1"] which makes the active connection and the connection ["192.167.0.0", "192.167.0.1"] cannot be impacted but only the 192.167.0.0 is active.
// 	6. for 192.167.0.2 the impacted connections are 1 , as 192.167.0.2 is active now and 192.167.0.0 is present earlier makes this a impacted connection ["192.167.0.2", "192.167.0.0"]
// 	7. for the second toggle entry of "192.167.0.0" , the impacted connections are 2 as the connections ["192.167.0.0", "192.167.0.1"] and ["192.167.0.2", "192.167.0.0"] contains ["192.167.0.2", "192.167.0.0"] and makes the connections inactive from inactive.
// 	8. 0.0.0.0 is not connected to anyone as it self address. so 0 impact

// 	find the impacted connections for the array of connections pairs with the incoming toggle ips
//
type Node struct {
	Ip    string
	State bool
}
type Connection struct {
	Node1 *Node
	Node2 *Node
}

type Connections struct {
	Connections       map[string][]Connection
	StatusConnections map[Connection]bool
}

func NewConnections() Connections {
	return Connections{Connections: make(map[string][]Connection), StatusConnections: make(map[Connection]bool)}
}

func (conn *Connections) Add(connection Connection) {

	if existingConnections, ok := conn.Connections[connection.Node1.Ip]; ok {
		conn.Connections[connection.Node1.Ip] = append(existingConnections, connection)
	}
	if existingConnections, ok := conn.Connections[connection.Node2.Ip]; ok {
		conn.Connections[connection.Node2.Ip] = append(existingConnections, connection)
	}

	conn.StatusConnections[connection] = false

}

func (connections *Connections) isConnectionActive(con Connection) bool {

	return connections.StatusConnections[con]
}

func (Connections *Connections) Toggle(ipaddres string) int {
	impacted_connections := 0
	if impactableConnections, ok := Connections.Connections[ipaddres]; ok {

		for _, impactableConnection := range impactableConnections {
			initialState := Connections.isConnectionActive(impactableConnection)
			Connections.ChangeState(&impactableConnection, ipaddres)
			finalState := Connections.isConnectionActive(impactableConnection)

			if finalState != initialState {
				impacted_connections++
			}

		}

	}
	return impacted_connections
}

func (Connections *Connections) ChangeState(impactableConnection *Connection, ipaddres string) {
	if impactableConnection.Node1.Ip == ipaddres {
		if impactableConnection.Node1.State {
			impactableConnection.Node1.State = false
		} else {
			impactableConnection.Node1.State = true
		}
	}
	if impactableConnection.Node2.Ip == ipaddres {
		if impactableConnection.Node2.State {
			impactableConnection.Node2.State = false
		} else {
			impactableConnection.Node2.State = true
		}
	}
	Connections.MakeActive(impactableConnection)

}

func (Connections *Connections) MakeActive(impactableConnection *Connection) {
	if impactableConnection.Node1.State && impactableConnection.Node2.State {
		Connections.StatusConnections[*impactableConnection] = true
	} else {
		Connections.StatusConnections[*impactableConnection] = false
	}

}

//NewConnections := arrays.NewConnections()
// Input := [][]string{{"192.167.0.0", "192.167.0.1"}, {"192.167.0.2", "192.167.0.0"}, {"192.167.0.0", "192.167.0.3"}}

// for _, connection := range Input {
// 	connectionObject := arrays.Connection{
// 		Node1: &arrays.Node{Ip: connection[0], State: false},
// 		Node2: &arrays.Node{Ip: connection[1], State: false}}

// 	if existingConnections, ok := NewConnections.Connections[connection[0]]; ok {
// 		NewConnections.Connections[connection[0]] = append(existingConnections, connectionObject)
// 	} else {
// 		NewConnections.Connections[connection[0]] = append(NewConnections.Connections[connection[0]], connectionObject)
// 	}

// 	if existingConnections, ok := NewConnections.Connections[connection[1]]; ok {
// 		NewConnections.Connections[connection[1]] = append(existingConnections, connectionObject)
// 	} else {
// 		NewConnections.Connections[connection[1]] = append(NewConnections.Connections[connection[1]], connectionObject)
// 	}

// }
// toggledIps := []string{"192.167.0.1", "192.167.0.0", "192.167.0.2", "192.167.0.0", "0.0.0.0"}

// impacted_connections := make([]int, 0)

// for _, ip := range toggledIps {
// 	impacted_connections = append(impacted_connections, NewConnections.Toggle(ip))
// }
// fmt.Println(impacted_connections)
