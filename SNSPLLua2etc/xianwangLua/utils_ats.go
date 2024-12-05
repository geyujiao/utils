package xianwanglua

import (
	"strings"
)

func L1Ats(fileName, fileName2 string) (luaMap map[string]int, luaList []string) {
	list := ReadOriginalData(fileName)
	luaMap = make(map[string]int, 0)
	luaList = make([]string, 0)
	for _, info := range list {
		if strings.Contains(info, ".lua") {
			if strings.Contains(info, "#") {
				continue
			}
			strs := strings.Split(info, "tslua.so")
			if len(strs) >= 2 {
				strs := strings.Split(strs[1], ".lua")
				if len(strs) < 2 {
					continue
				}

				key := strs[1] + ".lua"
				key = strings.ReplaceAll(key, "@pparam=", "")
				key = strings.ReplaceAll(key, " ", "")

				if _, exist := luaMap[key]; !exist {
					luaList = append(luaList, key)
				}

				luaMap[key] = luaMap[key] + 1
			}
		}

	}
	luaNumTotal := 0
	for _, num := range luaMap {
		luaNumTotal = luaNumTotal + num
	}
	// bytes, _ := json.Marshal(luaMap)
	// writerFile("L1-nginx.json", bytes)

	println("luaNumTotal--ats----L1---=", luaNumTotal) // luaNumTotal--ats----L1---= 379
	writerFileByline(fileName2, luaList)

	return luaMap, luaList
}
