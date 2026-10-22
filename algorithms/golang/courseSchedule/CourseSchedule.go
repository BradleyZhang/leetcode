// Source : https://leetcode.com/problems/course-schedule
// Author : BradleyZhang
// Date   : 2026-10-21

/*****************************************************************************************************
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are
 * given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course
 * bi first if you want to take course ai.
 *
 * 	For example, the pair [0, 1], indicates that to take course 0 you have to first take course
 * 1.
 *
 * Return true if you can finish all courses. Otherwise, return false.
 *
 * Example 1:
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you should also have
 * finished course 1. So it is impossible.
 *
 * Constraints:
 *
 * 	1 <= numCourses <= 2000
 * 	0 <= prerequisites.length <= 5000
 * 	prerequisites[i].length == 2
 * 	0 <= ai, bi < numCourses
 * 	All the pairs prerequisites[i] are unique.
 ******************************************************************************************************/
package courseschedule

const (
	White = iota
	Grey
	Black
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		graph[b] = append(graph[b], a)
	}

	path := make([]bool, numCourses)
	visited := make([]bool, numCourses)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		path[node] = true
		visited[node] = true

		for _, next := range graph[node] {
			if !visited[next] && dfs(next) {
				return true
			}

			if path[next] {
				return true
			}
		}

		path[node] = false
		return false
	}

	for i := 0; i < numCourses; i++ {
		if dfs(i) {
			return false
		}
	}

	return true
}

func canFinishBackup(numCourses int, prerequisites [][]int) bool {
	graph := newGraphAdjList(prerequisites)
	return !graph.findCycle()
}

type graphAdjList struct {
	adjList map[int][]int
}

func newGraphAdjList(edges [][]int) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[int][]int),
	}
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}
func (g *graphAdjList) findCycle() bool {
	visited := make(map[int]int)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		flag := visited[start]
		if flag == White {
			visited[start] = Grey
		} else if flag == Grey {
			return true
		} else {
			return false
		}

		for _, k := range g.adjList[start] {
			if dfs(k) {
				return true
			}
			// visited[k] = Black
		}
		visited[start] = Black
		return false
	}
	for start := range g.adjList {
		if dfs(start) {
			return true
		}
		visited[start] = Black
	}
	return false
}

func (g *graphAdjList) addEdge(vet1 int, vet2 int) {
	_, ok := g.adjList[vet2]
	if !ok {
		panic("error")
	}
	// 添加边 vet1 <- vet2
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

func (g *graphAdjList) addVertex(vet int) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	g.adjList[vet] = make([]int, 0)
}
