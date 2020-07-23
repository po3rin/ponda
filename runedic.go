package ponda

type RuneDic map[rune]int32

func buildRuneDict(data [][]rune) RuneDic {
	rd := make(RuneDic, len(data))

	index := int32(1)
	for _, rs := range data {
		for _, r := range rs {
			_, ok := rd[r]
			if ok {
				continue
			}
			rd[r] = index
			index++
		}
	}

	return rd
}
