package pinyinSearch

func (s *SourceStore) MatchFullSpell(pinyin string) []string {
	return s.matchSub(pinyin, 0, nil)
}

func (s *SourceStore) pinyinSplit(pinyin string) {
	//var (
	//	i    = 0
	//	key  string
	//	res  []string
	//	info map[string]int
	//)
	//
	//for i < len(pinyin) {
	//	key = pinyin[0:i]
	//
	//	v, ok := s.match.Load(key)
	//	if !ok {
	//		i++
	//		continue
	//	}
	//
	//	info = v.(map[string]int)
	//
	//	i++
	//}
}

func (s *SourceStore) matchSub(pinyin string, dept int, exists map[string]int) []string {
	var (
		ok       bool
		v        interface{}
		i        int
		before   int
		key      string
		res      []string
		info     map[string]int
		nextDict = map[string]int{}
	)

	if "" == pinyin {
		for k, _ := range exists {
			res = append(res, k)
		}

		return res
	}

	for i = 1; i <= len(pinyin); i++ {
		key = pinyin[0:i]

		//fmt.Println(key)

		v, ok = s.match.Load(key)
		if !ok {
			continue
		}

		//fmt.Println("find", key, v)

		info = v.(map[string]int)
		if 0 == dept {
			res = append(res, s.matchSub(pinyin[i:], dept+1, info)...)
			continue
		}

		for k, v := range info {
			before, ok = exists[k]
			if !ok {
				continue
			}

			if before+1 != v {
				continue
			}

			nextDict[k] = v
		}

		if 0 != len(nextDict) {
			//fmt.Println("nextDict", nextDict)
			res = append(res, s.matchSub(pinyin[i:], dept+1, nextDict)...)
			continue
		}

		//miss[key] = 1
	}

	return res
}
