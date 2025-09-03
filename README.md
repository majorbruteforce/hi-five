# User Interaction

- [ ] **Accept Requests**

  - [ ] Accept matching requests based on **labels, words, and semantics**
  - [ ] Create a **vertex for the user** to send to the matching engine

- [ ] **Socket Management**

  - [ ] Track active WebSocket sessions for matched users
  - [ ] Send match notifications to connected users once a match is confirmed
  - [ ] Auto heal lost connections, retries

---

# Matching Engine Features

- [ ] **Blossom V Integration**

  - [ ] Expose Blossom V via **C++ binding (cgo)**
  - [ ] Define Go wrapper functions (`AddEdge`, `ComputeMatching`, `GetMatching`)

- [ ] **Graph Sharding**

  - [ ] Partition graph into **shards for locality**
  - [ ] Support **local fast online mode** with `limitedAugmentSearch`

- [ ] **Global Mode**

  - [ ] Trigger **full global recomputation** of matching across shards when needed

- [ ] **Data Ingestion**

  - [ ] Stream edges via **Go channels**
  - [ ] Maintain **in-memory buffer / graph state per shard**

- [ ] **Matching Parameters**

  - [ ] `MAX_DEPTH` — limit augment path exploration
  - [ ] `SIMPLE_THRESHOLD` — early-exit rule for quick matches
  - [ ] `GAIN_THRESHOLD` — filter matches based on incremental gain

---

# Future Scope

- [ ] **Edge Prioritization**

  - [ ] Weighted queues to handle high-priority edges first

- [ ] **Asynchronous Global Sync**

  - [ ] Background global recomputation without blocking local shard queries

- [ ] **Horizontal Scaling**

  - [ ] Multi-process shard workers with gRPC for inter-shard sync
  - [ ] Explore using **Raft/consensus** for distributed state if scaling across nodes

- [ ] **Dynamic Sharding**

  - [ ] Rebalance shards when edge distribution is skewed

- [ ] **Monitoring & Metrics**

  - [ ] Latency, throughput, and matching success rate dashboards
  - [ ] Prometheus metrics + Grafana visualization

- [ ] **Optimized Blossom**

  - [ ] Replace cgo wrapper with direct Go port if performance demands
  - [ ] Explore approximate/heuristic algorithms for ultra-low-latency paths

- [ ] **Persistence Layer**

  - [ ] Snapshot graph state periodically (SQLite / BadgerDB)
  - [ ] Recovery support after crash

- [ ] **Chat/Session Enhancements**

  - [ ] Match history persistence for users
  - [ ] Load balancing WebSocket servers for chat sessions
  - [ ] Chat session handoff for scaling beyond single server
