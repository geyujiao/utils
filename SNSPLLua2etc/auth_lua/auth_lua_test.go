package auth_lua

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestReadOriginalData(t *testing.T) {

	resultList := ReadOriginalData("所有migu_live_auth_tri.lua")
	// bytes, _ := json.Marshal(resultList)
	fmt.Println("---", len(resultList))

	ccsList := []string{}
	for _, info := range resultList {
		if strings.Contains(info, "/CCS/") {
			ccsList = append(ccsList, info)
			continue
		}
		//20220630-第二阶段/mgsp.vod.miguvideo.com/SNS/nginx/mgsp.vod.miguvideo.com.conf
		tmp := strings.Split(info, "/")
		info2 := strings.Replace(info, tmp[len(tmp)-1], "", 1)
		info2 = info2 + "lua/migu/migu_live_auth_tri.lua" ///etc/nginx/lua/migu/migu_live_auth_tri.lua
		// fmt.Println("info2", info2)

		// CopyFile("lua/migu_live_auth_tri.lua", info2)

	}
	ccslistBytes, _ := json.Marshal(ccsList)
	fmt.Println("ccslist", string(ccslistBytes))

}
func TestReadOriginalData2(t *testing.T) {

	resultList := ReadOriginalData("所有migu_live_auth.lua")
	// bytes, _ := json.Marshal(resultList)
	fmt.Println("---", len(resultList))

	ccsList := []string{}
	for _, info := range resultList {
		if strings.Contains(info, "/CCS/") {
			ccsList = append(ccsList, info)
			continue
		}
		tmp := strings.Split(info, "/")
		info2 := strings.Replace(info, tmp[len(tmp)-1], "", 1)
		info2 = info2 + "lua/migu/migu_live_auth.lua"
		// fmt.Println("info2", info2)

		// CopyFile("lua/migu_live_auth.lua", info2)

	}
	ccslistBytes, _ := json.Marshal(ccsList)
	fmt.Println("ccslist", string(ccslistBytes))

}

func TestReadOriginalData3(t *testing.T) {

	resultList := ReadOriginalData("所有migu_new_vod_auth.lua")
	// bytes, _ := json.Marshal(resultList)
	fmt.Println("---", len(resultList))

	ccsList := []string{}
	for _, info := range resultList {
		if strings.Contains(info, "/CCS/") {
			ccsList = append(ccsList, info)
			continue
		}
		tmp := strings.Split(info, "/")
		info2 := strings.Replace(info, tmp[len(tmp)-1], "", 1)
		info2 = info2 + "lua/migu/migu_new_vod_auth.lua"
		// fmt.Println("info2", info2)

		//CopyFile("lua/migu_new_vod_auth.lua", info2)

	}
	ccslistBytes, _ := json.Marshal(ccsList)
	fmt.Println("ccslist", string(ccslistBytes))

}
