package mutex

import "hash/fnv"

type hashRing struct {
	nodes []string
}

func newHashRing(weights map[string]int) *hashRing {
	nodes := make([]string, 0)
	for name, weight := range weights {
		if weight <= 0 {
			continue
		}
		for i := 0; i < weight; i++ {
			nodes = append(nodes, name)
		}
	}
	if len(nodes) == 0 {
		return &hashRing{nodes: []string{}}
	}
	return &hashRing{nodes: nodes}
}

func (r *hashRing) Locate(key string) string {
	if r == nil || len(r.nodes) == 0 {
		return ""
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	idx := int(h.Sum32()) % len(r.nodes)
	return r.nodes[idx]
}
