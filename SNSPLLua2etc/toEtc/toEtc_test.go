package snspllua2etc

import (
	"testing"
)

func TestReadOriginalData(t *testing.T) {
	OriginalToEtc("./snspllua-txt.txt", "etc.xlsx")
	// bytes, _ := json.Marshal(result1)
	// println("result--", string(bytes))
	// println("")
	// bytes2, _ := json.Marshal(result2)
	// println("result2--", string(bytes2))
	/*
		[
		"migu/v6-sc.miguvideo.com/CCS/ats/remap.config",
		"migu/v6-sc.miguvideo.com/CCS/nginx/v6-sc.miguvideo.com-l2.conf",
		"migu/v6-sc.miguvideo.com/SNS/ats/parent.config",
		"migu/v6-sc.miguvideo.com/SNS/ats/remap.config",
		"migu/v6-sc.miguvideo.com/SNS/nginx/v6-sc.miguvideo.com.conf"
		]
	*/
	/*
			{
		    "v6-sc.miguvideo.com":[
		        "/etc/trafficserver/parent.config",
		        "/etc/trafficserver/remap.config",
		        "/etc/nginx/sites-enabled/v6-sc.miguvideo.com.conf"
		    ],
		    "v7-sc.miguvideo.com":[
		        "/etc/trafficserver/parent.config",
		        "/etc/trafficserver/remap.config",
		        "/etc/nginx/sites-enabled/v7-sc.miguvideo.com.conf"
		    ]
		}
			{
		    "v6-sc.miguvideo.com":[
		        "/etc/trafficserver/remap.config",
		        "/etc/nginx/sites-enabled/v6-sc.miguvideo.com-l2.conf"
		    ],
		    "v7-sc.miguvideo.com":[
		        "/etc/trafficserver/remap.config",
		        "/etc/nginx/sites-enabled/v7-sc.miguvideo.com-l2.conf"
		    ]
		}

	*/
}
