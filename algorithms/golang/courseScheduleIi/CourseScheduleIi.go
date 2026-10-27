// Source : https://leetcode.com/problems/course-schedule-ii
// Author : BradleyZhang
// Date   : 2026-10-26

/*****************************************************************************************************
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are
 * given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course
 * bi first if you want to take course ai.
 *
 * 	For example, the pair [0, 1], indicates that to take course 0 you have to first take course
 * 1.
 *
 * Return the ordering of courses you should take to finish all courses. If there are many valid
 * answers, return any of them. If it is impossible to finish all courses, return an empty array.
 *
 * Example 1:
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you should have finished
 * course 0. So the correct course order is [0,1].
 *
 * Example 2:
 *
 * Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both
 * courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
 * So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
 *
 * Example 3:
 *
 * Input: numCourses = 1, prerequisites = []
 * Output: [0]
 *
 * Constraints:
 *
 * 	1 <= numCourses <= 2000
 * 	0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * 	prerequisites[i].length == 2
 * 	0 <= ai, bi < numCourses
 * 	ai != bi
 * 	All the pairs [ai, bi] are distinct.
 ******************************************************************************************************/

package coursescheduleii

// 用数组优化
func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses) // 邻接表
	for i := range prerequisites {
		// ui->vi
		ui, vi := prerequisites[i][1], prerequisites[i][0]
		adj[ui] = append(adj[ui], vi)
	}
	indeg := make([]int, numCourses) // 入度
	for u := range numCourses {
		for _, v := range adj[u] {
			indeg[v]++
		}
	}
	ans := make([]int, 0, numCourses)
	for v := range numCourses {
		if indeg[v] == 0 {
			ans = append(ans, v)
		}
	}
	for head := 0; head < len(ans); head++ {
		for _, v := range adj[ans[head]] {
			indeg[v]--
			if indeg[v] == 0 {
				ans = append(ans, v)
			}
		}
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}

// 拓扑排序
func findOrderBackup(numCourses int, prerequisites [][]int) []int {
	graph := newGraphAdjList(prerequisites, numCourses)
	var ans []int
	for len(graph.zeroIndegree) > 0 {
		var course int
		for course = range graph.zeroIndegree {
			break
		}
		ans = append(ans, course)
		graph.trimVertex(course)
		delete(graph.zeroIndegree, course)
	}
	if len(graph.adjList) > 0 {
		return []int{}
	}
	return ans
}

type graphAdjList struct {
	vertexes     []*Vertex
	adjList      map[int][]*Vertex
	zeroIndegree map[int]struct{}
}
type Vertex struct {
	course   int
	indegree int
}

func newGraphAdjList(edges [][]int, num int) *graphAdjList {
	g := &graphAdjList{
		adjList:      make(map[int][]*Vertex, num),
		zeroIndegree: make(map[int]struct{}, num),
		vertexes:     make([]*Vertex, num),
	}
	for _, edge := range edges {
		g.addVertex(&Vertex{course: edge[0]})
		g.addVertex(&Vertex{course: edge[1]})
		g.addEdge(g.vertexes[edge[0]], g.vertexes[edge[1]])
	}
	for i := 0; i < num; i++ {
		if g.vertexes[i] == nil || g.vertexes[i].indegree == 0 {
			g.zeroIndegree[i] = struct{}{}
		}
	}
	return g
}

func (g *graphAdjList) addEdge(vet1 *Vertex, vet2 *Vertex) {
	_, ok := g.adjList[vet2.course]
	if !ok {
		panic("error")
	}
	// 添加边 vet1 <- vet2
	g.adjList[vet2.course] = append(g.adjList[vet2.course], vet1)
	vet1.indegree++
}

func (g *graphAdjList) addVertex(vet *Vertex) {
	_, ok := g.adjList[vet.course]
	if ok {
		return
	}
	g.vertexes[vet.course] = vet
	g.adjList[vet.course] = make([]*Vertex, 0)
}

func (g *graphAdjList) trimVertex(course int) {
	for _, vet := range g.adjList[course] {
		vet.indegree--
		if vet.indegree == 0 {
			g.zeroIndegree[vet.course] = struct{}{}
		}
	}
	delete(g.adjList, course)
}
