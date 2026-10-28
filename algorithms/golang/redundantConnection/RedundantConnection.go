// Source : https://leetcode.com/problems/redundant-connection
// Author : BradleyZhang
// Date   : 2026-10-27

/*****************************************************************************************************
 *
 * In this problem, a tree is an undirected graph that is connected and has no cycles.
 *
 * You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional
 * edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that
 * already existed. The graph is represented as an array edges of length n where edges[i] = [ai, bi]
 * indicates that there is an edge between nodes ai and bi in the graph.
 *
 * Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are
 * multiple answers, return the answer that occurs last in the input.
 *
 * Example 1:
 *
 * Input: edges = [[1,2],[1,3],[2,3]]
 * Output: [2,3]
 *
 * Example 2:
 *
 * Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
 * Output: [1,4]
 *
 * Constraints:
 *
 * 	n == edges.length
 * 	3 <= n <= 1000
 * 	edges[i].length == 2
 * 	1 <= ai < bi <= edges.length
 * 	ai != bi
 * 	There are no repeated edges.
 * 	The given graph is connected.
 ******************************************************************************************************/

package redundantconnection

// Union-Find
func findRedundantConnection(edges [][]int) []int {
	size := len(edges)
	uf := NewUnionFind(size)
	for _, edge := range edges {
		if uf.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}

type UnionFind struct {
	size   int // 连通分量
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size+1)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{
		size:   size,
		parent: parent,
	}
}

func (u *UnionFind) Union(a, b int) (connected bool) {
	ap := u.Find(a)
	bp := u.Find(b)
	if ap == bp {
		return true
	}
	u.parent[ap] = bp // representative 建立连接-> set 建立连接
	u.size--
	return false
}
func (u *UnionFind) Find(a int) int {
	ap := u.parent[a]
	if ap == a {
		return ap
	}
	path := []int{}
	for a != ap {
		path = append(path, a)
		a = ap
		ap = u.parent[ap]
	}
	for _, x := range path { // 路径压缩
		u.parent[x] = ap
	}
	return ap
}
