# max_weighted_matching.py

import networkx as nx
import json
import sys

def max_weighted_matching(edges):
    G = nx.Graph()
    G.add_weighted_edges_from(edges)
    matching = nx.max_weight_matching(G, maxcardinality=False)
    return list(matching)

if __name__ == "__main__":
    edges = json.loads(sys.argv[1])
    result = max_weighted_matching(edges)
    print(json.dumps(result))
