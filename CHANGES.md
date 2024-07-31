## Optimizations:

- Instead of dumping the whole candidate list, add it to the graph as soon as it is recieved.
- Implement a dedicated sessions manager
- Handle deadlocks and create locks for event routing and handlers
- Implement a ping-pong
- Implement the stop channel