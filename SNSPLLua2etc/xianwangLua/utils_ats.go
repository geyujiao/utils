package xianwanglua

import (
	"strings"
)

func L1Ats(fileName, fileName2 string) (luaMap map[string]string, luaList []string) {
	list := ReadOriginalData(fileName)
	luaMap = make(map[string]string, 0)
	luaList = make([]string, 0)
	for _, info := range list {
		if strings.Contains(info, "#") {
			continue
		}
		if !strings.Contains(info, ".lua") {
			continue
		}

		strs := strings.Split(info, "tslua.so")
		if len(strs) < 2 {
			continue
		}
		// println("24---" + strs[1])
		strs2 := strings.Split(strs[1], ".lua")
		if len(strs2) < 2 {
			continue
		}
		// println("29---" + strs2[0])

		key := strs2[0] + ".lua"
		key = strings.ReplaceAll(key, "@pparam=", "")
		key = strings.ReplaceAll(key, " ", "")

		// println("35--" + key)
		if _, exist := luaMap[key]; !exist {
			luaList = append(luaList, key)
		}

		luaMap[key] = key

	}

	// luaNumTotal := 0
	// for _, num := range luaMap {
	// 	luaNumTotal = luaNumTotal + num
	// }
	// bytes, _ := json.Marshal(luaMap)
	// writerFile("L1-nginx.json", bytes)

	// println("luaNumTotal--ats----L1---=", luaNumTotal) // luaNumTotal--ats----L1---= 379
	// writerFileByline(fileName2, luaList)

	return luaMap, luaList
}

func L2Ats(fileName, fileName2 string) (luaMap map[string]string, luaList []string) {
	list := ReadOriginalData(fileName)
	luaMap = make(map[string]string, 0)
	luaList = make([]string, 0)
	for _, info := range list {
		if strings.Contains(info, "#") {
			continue
		}
		if !strings.Contains(info, ".lua") {
			continue
		}

		strs := strings.Split(info, "tslua.so")
		if len(strs) < 2 {
			continue
		}
		// println("24---" + strs[1])
		strs2 := strings.Split(strs[1], ".lua")
		if len(strs2) < 2 {
			continue
		}
		// println("29---" + strs2[0])

		key := strs2[0] + ".lua"
		key = strings.ReplaceAll(key, "@pparam=", "")
		key = strings.ReplaceAll(key, " ", "")

		// println("35--" + key)
		if _, exist := luaMap[key]; !exist {
			luaList = append(luaList, key)
		}

		luaMap[key] = key

	}

	// luaNumTotal := 0
	// for _, num := range luaMap {
	// 	luaNumTotal = luaNumTotal + num
	// }
	// bytes, _ := json.Marshal(luaMap)
	// writerFile("L1-nginx.json", bytes)

	// println("luaNumTotal--ats----L2---=", luaNumTotal) // luaNumTotal--ats----L2---= 367
	// writerFileByline(fileName2, luaList)

	return luaMap, luaList
}
