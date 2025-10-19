package graph

type DisjointSet struct {
	Parent []int
	Rank   []int
	Size   []int
}

func NewDisjointSet(size int) *DisjointSet {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DisjointSet{
		Parent: parent,
		Rank:   rank,
	}
}

func (ds *DisjointSet) FindPar(u int) int {
	if ds.Parent[u]==u{
		return u
	}
	ds.Parent[u]=ds.FindPar(ds.Parent[u])
	return ds.Parent[u]
}

func (ds *DisjointSet) UnionByRank(u int, v int) {
	pu := ds.FindPar(u)
	pv := ds.FindPar(v)
	if pu == pv {
		return
	}
	if ds.Rank[pu] < ds.Rank[pv] {
		ds.Parent[pu] = pv
	} else if ds.Rank[pv] < ds.Rank[pu] {
		ds.Parent[pv] = pu
	} else {
		ds.Parent[pu] = pv
		ds.Rank[pv]++
	}
}

func (ds* DisjointSet) UnionBySize(u int, v int) {
	pu := ds.FindPar(u)
	pv := ds.FindPar(v)
	if pu == pv {
		return
	}
	if ds.Size[pu] < ds.Size[pv] {
		ds.Parent[pu] = pv
		ds.Size[pv] += ds.Size[pu]
	} else {
		ds.Parent[pv] = pu
		ds.Size[pu] += ds.Size[pv]
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	ds := NewDisjointSet(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				ds.UnionByRank(i, j)
			}
		}
	}
	provinces := 0
	for i := 0; i < n; i++ {
		if ds.FindPar(i) == i {
			provinces++
		}
	}
	return provinces
}