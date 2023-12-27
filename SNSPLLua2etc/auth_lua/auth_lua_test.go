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

func TestXxx(t *testing.T) {
	str := `#EXTM3U
	#EXT-X-VERSION:3
	#EXT-X-MEDIA-SEQUENCE:0
	#EXT-X-TARGETDURATION:5
	#EXT-X-PROGRAM-DATE-TIME:2023-12-17T15:24:48.000+08:00
	#EXTINF:4.291,
	1702797888.ts?expire=4070880000&start=1702797880&end=1702797910&sign=7335cd2dc2ff50225ee5320919aa57f1&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODg4LnRzP3V1aWQ9YjBkMzg0MDctYWIzMi00ZGU0LTlkNmItNWZjMTg4YjkyZTEw
	#EXTINF:4.007,
	1702797889.ts?expire=4070880000&start=1702797880&end=1702797910&sign=b3e21accb323a6e66014702f97850933&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODg5LnRzP3V1aWQ9YjBkMzg0MDctYWIzMi00ZGU0LTlkNmItNWZjMTg4YjkyZTEw
	#EXTINF:3.979,
	1702797890.ts?expire=4070880000&start=1702797880&end=1702797910&sign=a4fe2f674140702c73ef74945e0680a8&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODkwLnRzP3V1aWQ9YjBkMzg0MDctYWIzMi00ZGU0LTlkNmItNWZjMTg4YjkyZTEw
	#EXTINF:3.981,
	1702797894.ts?expire=4070880000&start=1702797880&end=1702797910&sign=2679f841b564996ac9bdb82b0f099507&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODk0LnRzP3V1aWQ9YjBkMzg0MDctYWIzMi00ZGU0LTlkNmItNWZjMTg4YjkyZTEw
	#EXTINF:1.607,
	1702797898.ts?expire=4070880000&start=1702797880&end=1702797910&sign=48f255d4f61c7acd7ce95ea429f1343c&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODk4LnRzP3V1aWQ9YjBkMzg0MDctYWIzMi00ZGU0LTlkNmItNWZjMTg4YjkyZTEw
	#EXT-X-ENDLIST`

	str2 := `#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-TARGETDURATION:5
#EXT-X-PROGRAM-DATE-TIME:2023-12-17T15:24:48.000+08:00
#EXTINF:4.291,
1702797888.ts?expire=4070880000&start=1702797880&end=1702797910&sign=7335cd2dc2ff50225ee5320919aa57f1&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODg4LnRzP3V1aWQ9ODg5NmE3MTMtZWM2MS00NGIxLTk2YTctYTM5NDA1ODFmYTcw
#EXTINF:4.007,
1702797889.ts?expire=4070880000&start=1702797880&end=1702797910&sign=b3e21accb323a6e66014702f97850933&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODg5LnRzP3V1aWQ9ODg5NmE3MTMtZWM2MS00NGIxLTk2YTctYTM5NDA1ODFmYTcw
#EXTINF:3.979,
1702797890.ts?expire=4070880000&start=1702797880&end=1702797910&sign=a4fe2f674140702c73ef74945e0680a8&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODkwLnRzP3V1aWQ9ODg5NmE3MTMtZWM2MS00NGIxLTk2YTctYTM5NDA1ODFmYTcw
#EXTINF:3.981,
1702797894.ts?expire=4070880000&start=1702797880&end=1702797910&sign=2679f841b564996ac9bdb82b0f099507&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODk0LnRzP3V1aWQ9ODg5NmE3MTMtZWM2MS00NGIxLTk2YTctYTM5NDA1ODFmYTcw
#EXTINF:1.607,
1702797898.ts?expire=4070880000&start=1702797880&end=1702797910&sign=48f255d4f61c7acd7ce95ea429f1343c&real=L21naHljZ3Rlc3QvY2xvdWRnYW1ldGVzdC8yMzk5NzQ2OTY0OTQtMTcwMjc5Nzg4MDE0Ny8xNzAyNzk3ODk4LnRzP3V1aWQ9ODg5NmE3MTMtZWM2MS00NGIxLTk2YTctYTM5NDA1ODFmYTcw
#EXT-X-ENDLIST
`
	len1 := len(str)
	fmt.Println("lll11=", len1)

	len2 := len(str2)
	fmt.Println("lll22=", len2)

}
