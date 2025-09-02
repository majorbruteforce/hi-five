## Optimizations:

- When client closes connection, retract any existing matchmaking request
- Instead of dumping the whole candidate list, add it to the graph as soon as it is recieved.
- Handle potential deadlocks and create locks for event routing and handlers
- Implement a ping-pong
- Implement the stop channel