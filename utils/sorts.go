package utils

type person struct {
	Name string
	Rank int // 被排序
}

type persons []person

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	return p[i].Rank < p[j].Rank
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ranks []Rank

func (p ranks) Len() int {
	return len(p)
}

func (p ranks) Less(i, j int) bool {
	return p[i].index < p[j].index
}

func (p ranks) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Rank struct {
	index int
}

type SortDemo struct {
	*Rank
	Name string
}
