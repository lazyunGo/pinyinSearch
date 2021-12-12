package pinyinSearch

import "sync"

type SourceStore struct {
	source sync.Map // map[string][][]string // id: [[gai], [lun]]
	match  sync.Map // map[string]map[string]int // gai: {id: index}
	prefix sync.Map // map[string]map[string]int // g: {id: index}
}

func New() *SourceStore {
	return &SourceStore{
		source: sync.Map{},
		match:  sync.Map{},
		prefix: sync.Map{},
	}
}

func (s *SourceStore) Store(id string, value [][]string) {
	s.source.Store(id, value)

	var (
		ok   bool
		key  string
		info map[string]int
		old  interface{}
	)

	// key: [[v1] [v2, v2.1]]
	/* match {
		v1:   key: index1
		v2:   key: index2
		v2.1: key: index2
	}
	*/

	for index := range value {

		for i := range value[index] {
			key = value[index][i]

			old, ok = s.match.Load(key)
			if ok {
				info = old.(map[string]int)
			} else {
				info = map[string]int{}
			}

			info[id] = index

			s.match.Store(key, info)
		}

	}
}
